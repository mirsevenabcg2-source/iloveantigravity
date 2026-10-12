param (
    [ValidateSet("start", "stop", "restart", "status", "test", "install-autostart", "uninstall-autostart")]
    [string]$Action = "status"
)

$PythonExe = "C:\Users\Admin\AppData\Local\Programs\Python\Python311\python.exe"
$PythonWExe = "C:\Users\Admin\AppData\Local\Programs\Python\Python311\pythonw.exe"
$ServiceScript = "C:\Users\Admin\.gemini\config\scripts\mem0_service.py"
$LogFile = "C:\Users\Admin\.gemini\antigravity\mem0_data\mem0_service.log"
$Port = 8765
$RegPath = "HKCU:\Software\Microsoft\Windows\CurrentVersion\Run"
$RegName = "Mem0_Antigravity_Service"

function Get-Mem0Process {
    $conns = Get-NetTCPConnection -LocalPort $Port -State Listen -ErrorAction SilentlyContinue
    if ($conns) {
        $pids = $conns | Select-Object -ExpandProperty OwningProcess -Unique
        return Get-Process -Id $pids -ErrorAction SilentlyContinue
    }
    return $null
}

function Show-Status {
    $proc = Get-Mem0Process
    Write-Host "==========================================" -ForegroundColor Cyan
    Write-Host "   MEM0 ANTIGRAVITY SERVICE STATUS" -ForegroundColor Cyan
    Write-Host "==========================================" -ForegroundColor Cyan
    
    if ($proc) {
        Write-Host "Service State : RUNNING (PID: $($proc.Id -join ', '))" -ForegroundColor Green
        Write-Host "Listening Port: $Port (SSE Endpoint: http://127.0.0.1:$Port/sse)" -ForegroundColor Green
    } else {
        Write-Host "Service State : STOPPED" -ForegroundColor Yellow
    }

    # Autostart registry status
    $regVal = (Get-ItemProperty -Path $RegPath -Name $RegName -ErrorAction SilentlyContinue).$RegName
    if ($regVal) {
        Write-Host "Auto-Start    : ENABLED (Starts automatically on Windows Boot/Logon)" -ForegroundColor Green
    } else {
        Write-Host "Auto-Start    : DISABLED" -ForegroundColor DarkGray
    }

    if (Test-Path $LogFile) {
        Write-Host "`nRecent Logs ($LogFile):" -ForegroundColor Gray
        Get-Content $LogFile -Tail 5 | ForEach-Object { Write-Host "  $_" -ForegroundColor DarkGray }
    }
    Write-Host "==========================================" -ForegroundColor Cyan
}

function Start-Mem0Service {
    $proc = Get-Mem0Process
    if ($proc) {
        Write-Host "Mem0 Service is already running on port $Port (PID: $($proc.Id -join ', '))." -ForegroundColor Yellow
        return
    }

    Write-Host "Starting Mem0 Service in background..." -ForegroundColor Cyan
    $workDir = "C:\Users\Admin\.gemini\antigravity\mem0_data"
    $cmd = "`"$PythonWExe`" `"$ServiceScript`""
    
    # Use WMI Win32_Process.Create so that the spawned process is decoupled from the current PowerShell console/Job Object
    $wmiRes = Invoke-CimMethod -ClassName Win32_Process -MethodName Create -Arguments @{
        CommandLine = $cmd
        CurrentDirectory = $workDir
    }
    
    if ($wmiRes.ReturnValue -ne 0) {
        # Fallback to Start-Process
        Start-Process -FilePath $PythonWExe -ArgumentList "`"$ServiceScript`"" -WorkingDirectory $workDir -WindowStyle Hidden
    }
    
    $timeout = 15
    $elapsed = 0
    $runningProc = $null
    while ($elapsed -lt $timeout) {
        Start-Sleep -Seconds 1
        $elapsed++
        $runningProc = Get-Mem0Process
        if ($runningProc) { break }
    }

    if ($runningProc) {
        Write-Host "Mem0 Service started successfully! (PID: $($runningProc.Id -join ', '))" -ForegroundColor Green
        Write-Host "SSE Endpoint: http://127.0.0.1:$Port/sse" -ForegroundColor Green
    } else {
        Write-Host "Service started in background. Check 'status' or logs if port isn't open yet." -ForegroundColor Yellow
    }
}

function Stop-Mem0Service {
    $proc = Get-Mem0Process
    if ($proc) {
        Write-Host "Stopping Mem0 Service (PID: $($proc.Id -join ', '))..." -ForegroundColor Yellow
        $proc | Stop-Process -Force -ErrorAction SilentlyContinue
        Start-Sleep -Seconds 1
        Write-Host "Mem0 Service stopped." -ForegroundColor Green
    } else {
        Write-Host "Mem0 Service is not running." -ForegroundColor Gray
    }
}

function Test-Mem0Endpoint {
    try {
        $req = [System.Net.WebRequest]::Create("http://127.0.0.1:$Port/sse")
        $req.Timeout = 3000
        $resp = $req.GetResponse()
        Write-Host "Mem0 SSE Endpoint is healthy! HTTP Status: $([int]$resp.StatusCode)" -ForegroundColor Green
        $resp.Close()
    } catch {
        Write-Host "Failed to connect to Mem0 SSE endpoint: $_" -ForegroundColor Red
    }
}

function Install-Autostart {
    $cmd = "`"$PythonWExe`" `"$ServiceScript`""
    Set-ItemProperty -Path $RegPath -Name $RegName -Value $cmd -Force
    Write-Host "Mem0 Auto-Start successfully installed in Windows Startup ($RegPath)." -ForegroundColor Green
    Write-Host "Service will automatically launch every time Windows boots/logs in." -ForegroundColor Green
}

function Uninstall-Autostart {
    Remove-ItemProperty -Path $RegPath -Name $RegName -ErrorAction SilentlyContinue
    Write-Host "Mem0 Auto-Start removed from Windows Startup." -ForegroundColor Yellow
}

switch ($Action) {
    "start" { Start-Mem0Service }
    "stop" { Stop-Mem0Service }
    "restart" { Stop-Mem0Service; Start-Mem0Service }
    "status" { Show-Status }
    "test" { Test-Mem0Endpoint }
    "install-autostart" { Install-Autostart }
    "uninstall-autostart" { Uninstall-Autostart }
}
