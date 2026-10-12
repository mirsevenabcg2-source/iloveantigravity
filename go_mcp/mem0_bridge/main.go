package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/exec"
	"time"

	"github.com/mark3labs/mcp-go/client"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

const (
	sseURL        = "http://127.0.0.1:8765/sse"
	hostPort      = "127.0.0.1:8765"
	pythonwExe    = `C:\Users\Admin\AppData\Local\Programs\Python\Python311\pythonw.exe`
	serviceScript = `C:\Users\Admin\.gemini\config\scripts\mem0_service.py`
	dataDir       = `C:\Users\Admin\.gemini\antigravity\mem0_data`
)

func isServiceListening() bool {
	conn, err := net.DialTimeout("tcp", hostPort, 300*time.Millisecond)
	if err != nil {
		return false
	}
	_ = conn.Close()
	return true
}

func ensureServiceRunning() error {
	if isServiceListening() {
		return nil
	}

	cmd := exec.Command(pythonwExe, serviceScript)
	cmd.Dir = dataDir
	if err := cmd.Start(); err != nil {
		return fmt.Errorf("failed to start mem0 background service: %w", err)
	}

	for i := 0; i < 20; i++ {
		time.Sleep(300 * time.Millisecond)
		if isServiceListening() {
			return nil
		}
	}
	return fmt.Errorf("timed out waiting for mem0 service to bind to %s", hostPort)
}

func callRemoteMem0(ctx context.Context, toolName string, args map[string]interface{}) (*mcp.CallToolResult, error) {
	if err := ensureServiceRunning(); err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	c, err := client.NewSSEMCPClient(sseURL)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Failed to create SSE client: %v", err)), nil
	}
	defer c.Close()

	if err := c.Start(ctx); err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Failed to start SSE transport: %v", err)), nil
	}

	initReq := mcp.InitializeRequest{}
	initReq.Params.ProtocolVersion = mcp.LATEST_PROTOCOL_VERSION
	initReq.Params.ClientInfo = mcp.Implementation{
		Name:    "mem0_go_bridge",
		Version: "1.0.0",
	}

	if _, err := c.Initialize(ctx, initReq); err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Failed to initialize MCP session: %v", err)), nil
	}

	req := mcp.CallToolRequest{}
	req.Params.Name = toolName
	req.Params.Arguments = args

	res, err := c.CallTool(ctx, req)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Tool call error: %v", err)), nil
	}
	return res, nil
}

func main() {
	s := server.NewMCPServer("mem0", "1.0.0")

	// 1. mem0_status
	s.AddTool(mcp.NewTool("mem0_status",
		mcp.WithDescription("Check the operational health and statistics of the Mem0 memory engine"),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		return callRemoteMem0(ctx, "mem0_status", map[string]interface{}{})
	})

	// 2. mem0_add
	s.AddTool(mcp.NewTool("mem0_add",
		mcp.WithDescription("Store a new fact, user preference, architectural decision, or convention into persistent memory"),
		mcp.WithString("text", mcp.Required(), mcp.Description("The fact or decision to remember")),
		mcp.WithString("user_id", mcp.Description("The user or entity identifier (default: 'developer')")),
		mcp.WithString("metadata", mcp.Description("Optional JSON string containing additional context metadata")),
		mcp.WithString("project", mcp.Description("Target project name (e.g. 'FreeExile', 'DSCons', 'JX', 'Global')")),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		if _, ok := args["user_id"]; !ok {
			args["user_id"] = "developer"
		}
		return callRemoteMem0(ctx, "mem0_add", args)
	})

	// 3. mem0_search
	s.AddTool(mcp.NewTool("mem0_search",
		mcp.WithDescription("Search long-term memories for relevant facts, preferences, or technical conventions"),
		mcp.WithString("query", mcp.Required(), mcp.Description("Search query")),
		mcp.WithString("user_id", mcp.Description("User identifier to filter memories")),
		mcp.WithNumber("limit", mcp.Description("Maximum number of memories to return (default: 5)")),
		mcp.WithString("project", mcp.Description("Target project filter (e.g. 'FreeExile', 'DSCons', 'Global')")),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		if _, ok := args["user_id"]; !ok {
			args["user_id"] = "developer"
		}
		if _, ok := args["limit"]; !ok {
			args["limit"] = 5
		}
		return callRemoteMem0(ctx, "mem0_search", args)
	})

	// 4. mem0_get_all
	s.AddTool(mcp.NewTool("mem0_get_all",
		mcp.WithDescription("Retrieve all stored memories and facts for a given user"),
		mcp.WithString("user_id", mcp.Description("The user ID to retrieve memories for")),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		if _, ok := args["user_id"]; !ok {
			args["user_id"] = "developer"
		}
		return callRemoteMem0(ctx, "mem0_get_all", args)
	})

	// 5. mem0_delete
	s.AddTool(mcp.NewTool("mem0_delete",
		mcp.WithDescription("Delete a specific memory by its unique ID"),
		mcp.WithString("memory_id", mcp.Required(), mcp.Description("The unique ID of the memory to delete")),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		return callRemoteMem0(ctx, "mem0_delete", request.GetArguments())
	})

	if err := server.ServeStdio(s); err != nil {
		fmt.Fprintf(os.Stderr, "Server error: %v\n", err)
	}
}
