package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"syscall"
	"unsafe"
)

type HookPayload struct {
	ToolCall struct {
		Name string                 `json:"name"`
		Args map[string]interface{} `json:"args"`
	} `json:"toolCall"`
	WorkspacePaths []string `json:"workspacePaths"`
}

type PreToolResponse struct {
	Decision string `json:"decision"`
	Reason   string `json:"reason,omitempty"`
}

const (
	TH32CS_SNAPPROCESS        = 0x00000002
	PROCESS_QUERY_INFORMATION = 0x0400
	PROCESS_SET_QUOTA         = 0x0100
)

type PROCESSENTRY32W struct {
	Size              uint32
	CntUsage          uint32
	ProcessID         uint32
	DefaultHeapID     uintptr
	ModuleID          uint32
	CntThreads        uint32
	ParentProcessID   uint32
	PcPriClassBase    int32
	Flags             uint32
	ExeFile           [260]uint16
}

type PROCESS_MEMORY_COUNTERS struct {
	Cb                         uint32
	PageFaultCount             uint32
	PeakWorkingSetSize         uintptr
	WorkingSetSize             uintptr
	QuotaPeakPagedPoolUsage    uintptr
	QuotaPagedPoolUsage        uintptr
	QuotaPeakNonPagedPoolUsage uintptr
	QuotaNonPagedPoolUsage     uintptr
	PagefileUsage              uintptr
	PeakPagefileUsage          uintptr
}

var (
	kernel32                     = syscall.NewLazyDLL("kernel32.dll")
	procCreateToolhelp32Snapshot = kernel32.NewProc("CreateToolhelp32Snapshot")
	procProcess32FirstW          = kernel32.NewProc("Process32FirstW")
	procProcess32NextW           = kernel32.NewProc("Process32NextW")
	procOpenProcess              = kernel32.NewProc("OpenProcess")
	procCloseHandle              = kernel32.NewProc("CloseHandle")

	psapi                    = syscall.NewLazyDLL("psapi.dll")
	procGetProcessMemoryInfo = psapi.NewProc("GetProcessMemoryInfo")
	procEmptyWorkingSet      = psapi.NewProc("EmptyWorkingSet")
)

func trimLanguageServers() {
	snap, _, _ := procCreateToolhelp32Snapshot.Call(TH32CS_SNAPPROCESS, 0)
	if snap == uintptr(syscall.InvalidHandle) {
		return
	}
	defer procCloseHandle.Call(snap)

	var entry PROCESSENTRY32W
	entry.Size = uint32(unsafe.Sizeof(entry))

	ret, _, _ := procProcess32FirstW.Call(snap, uintptr(unsafe.Pointer(&entry)))
	for ret != 0 {
		exeName := syscall.UTF16ToString(entry.ExeFile[:])
		if strings.EqualFold(exeName, "language_server.exe") {
			hProcess, _, _ := procOpenProcess.Call(
				uintptr(PROCESS_QUERY_INFORMATION|PROCESS_SET_QUOTA),
				0,
				uintptr(entry.ProcessID),
			)
			if hProcess != 0 {
				var memCounters PROCESS_MEMORY_COUNTERS
				memCounters.Cb = uint32(unsafe.Sizeof(memCounters))
				ok, _, _ := procGetProcessMemoryInfo.Call(
					hProcess,
					uintptr(unsafe.Pointer(&memCounters)),
					uintptr(memCounters.Cb),
				)
				if ok != 0 {
					wsMB := memCounters.WorkingSetSize / (1024 * 1024)
					if wsMB > 1500 {
						procEmptyWorkingSet.Call(hProcess)
					}
				}
				procCloseHandle.Call(hProcess)
			}
		}
		ret, _, _ = procProcess32NextW.Call(snap, uintptr(unsafe.Pointer(&entry)))
	}
}

func validateSyntax(filePath string) (bool, string) {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return true, "File does not exist yet"
	}

	ext := strings.ToLower(filepath.Ext(filePath))
	switch ext {
	case ".json":
		data, err := os.ReadFile(filePath)
		if err != nil {
			return false, fmt.Sprintf("I/O error reading JSON: %v", err)
		}
		if len(data) >= 3 && data[0] == 0xEF && data[1] == 0xBB && data[2] == 0xBF {
			return false, "UTF-8 BOM detected in JSON. Please save without BOM."
		}
		var js interface{}
		if err := json.Unmarshal(data, &js); err != nil {
			return false, fmt.Sprintf("JSON syntax error: %v", err)
		}
	case ".py":
		cmd := exec.Command("python", "-m", "py_compile", filePath)
		var stderr bytes.Buffer
		cmd.Stderr = &stderr
		if err := cmd.Run(); err != nil {
			return false, fmt.Sprintf("Python syntax error: %s", stderr.String())
		}
	}
	return true, "OK"
}

func main() {
	mode := flag.String("mode", "post", "hook mode: pre or post")
	flag.Parse()

	// 1. Always trim memory in background
	go trimLanguageServers()

	// 2. Read STDIN
	inputBytes, _ := io.ReadAll(os.Stdin)

	if *mode == "pre" {
		var payload HookPayload
		_ = json.Unmarshal(inputBytes, &payload)
		args := payload.ToolCall.Args
		var targetFile string
		for _, key := range []string{"TargetFile", "target_file", "targetFile", "AbsolutePath", "path"} {
			if v, ok := args[key]; ok && v != nil {
				if s, ok := v.(string); ok && s != "" {
					targetFile = s
					break
				}
			}
		}
		if targetFile != "" {
			valid, reason := validateSyntax(targetFile)
			if !valid {
				out, _ := json.Marshal(PreToolResponse{
					Decision: "deny",
					Reason:   fmt.Sprintf("[AIDER GUARD BLOCKED] %s", reason),
				})
				fmt.Println(string(out))
				return
			}
		}
		out, _ := json.Marshal(PreToolResponse{Decision: "allow"})
		fmt.Println(string(out))
		return
	}

	// 3. PostToolUse contract strictly expects `{}` on stdout
	fmt.Println("{}")
}
