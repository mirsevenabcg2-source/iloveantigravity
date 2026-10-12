<#
.SYNOPSIS
    iloveantigravity - 1-Click Turnkey Setup Script for Windows
.DESCRIPTION
    Installs, compiles, and configures the complete Antigravity Supercharged Agent Ecosystem:
    - High-Performance Go MCP Servers (git, mem0_bridge, memory, sqlite)
    - Native Go Lifecycle Hooks (aider_guard)
    - PreToolUse Impact Guard & Blast Radius engine
    - Local Mem0 Background Memory Service (SSE port 8765, Qdrant + Google Antigravity SDK)
    - Global Engineering Directives (AGENTS.md & GEMINI.md)
    - Battle-Tested Plugins & Custom Skills
.EXAMPLE
    .\install.ps1
    .\install.ps1 -EnableAutoStart -SkipPip
#>

[CmdletBinding()]
param (
    [switch]$EnableAutoStart = $true,
    [switch]$SkipPip = $false,
    [switch]$ForceRebuild = $false
)

$ErrorActionPreference = "Stop"

function Write-Header {
    param([string]$text)
    Write-Host "`n=================================================================" -ForegroundColor Cyan
    Write-Host "   $text" -ForegroundColor Cyan
    Write-Host "=================================================================" -ForegroundColor Cyan
}

function Write-Step {
    param([string]$title, [string]$status = "OK")
    Write-Host "  [*] $title... " -NoNewline -ForegroundColor White
    Write-Host "[$status]" -ForegroundColor Green
}

function Write-WarnMsg {
    param([string]$text)
    Write-Host "  [!] WARNING: $text" -ForegroundColor Yellow
}

function Write-ErrMsg {
    param([string]$text)
    Write-Host "  [X] ERROR: $text" -ForegroundColor Red
}

Write-Header "iLoveAntigravity - Turnkey Agent Supercharger (Windows)"

$RepoRoot = Split-Path -Parent $MyInvocation.MyCommand.Path
$UserHome = [System.Environment]::GetFolderPath('UserProfile')
$GeminiDir = Join-Path $UserHome ".gemini"
$ConfigDir = Join-Path $GeminiDir "config"
$AntigravityDir = Join-Path $GeminiDir "antigravity"
$BinDir = Join-Path $ConfigDir "bin"
$ScriptsDir = Join-Path $ConfigDir "scripts"
$SkillsDir = Join-Path $ConfigDir "skills"
$PluginsDir = Join-Path $ConfigDir "plugins"
$Mem0DataDir = Join-Path $AntigravityDir "mem0_data"

# 1. Check Prerequisites
Write-Header "Step 1: Checking System Prerequisites"

# Python Detection
$PythonCmd = Get-Command "python" -ErrorAction SilentlyContinue
if (-not $PythonCmd) {
    $PythonCmd = Get-Command "python3" -ErrorAction SilentlyContinue
}
if (-not $PythonCmd) {
    Write-ErrMsg "Python is not installed or not in PATH! Please install Python 3.10+ from python.org"
    exit 1
}
$PythonExe = $PythonCmd.Source
$PythonWExe = $PythonExe -replace "python\.exe$", "pythonw.exe"
if (-not (Test-Path $PythonWExe)) {
    $PythonWExe = $PythonExe
}
Write-Step "Python detected ($PythonExe)"

# Go Detection
$GoCmd = Get-Command "go" -ErrorAction SilentlyContinue
if (-not $GoCmd) {
    $standardGo = "C:\Program Files\Go\bin\go.exe"
    if (Test-Path $standardGo) {
        $GoExe = $standardGo
    } else {
        Write-ErrMsg "Go compiler not found! Please install Go 1.22+ from https://go.dev"
        exit 1
    }
} else {
    $GoExe = $GoCmd.Source
}
Write-Step "Go compiler detected ($GoExe)"

# Node & NPX Detection
$NodeCmd = Get-Command "node" -ErrorAction SilentlyContinue
$NpxCmd = Get-Command "npx" -ErrorAction SilentlyContinue
if ($NodeCmd) {
    Write-Step "Node.js detected ($($NodeCmd.Source))"
} else {
    Write-WarnMsg "Node.js not detected in PATH. Node MCPs will require Node.js 18+"
}

# Git Detection
$GitCmd = Get-Command "git" -ErrorAction SilentlyContinue
if ($GitCmd) {
    Write-Step "Git detected ($($GitCmd.Source))"
} else {
    Write-ErrMsg "Git is required but not found in PATH."
    exit 1
}

# 2. Prepare Directories
Write-Header "Step 2: Preparing ~/.gemini Directory Hierarchy"
$Dirs = @($ConfigDir, $AntigravityDir, $BinDir, $ScriptsDir, $SkillsDir, $PluginsDir, $Mem0DataDir, (Join-Path $Mem0DataDir "qdrant"))
foreach ($d in $Dirs) {
    if (-not (Test-Path $d)) {
        New-Item -ItemType Directory -Path $d -Force | Out-Null
    }
}
Write-Step "Directories initialized successfully"

# 3. Compile Native Go MCP Servers & Hooks
Write-Header "Step 3: Compiling High-Performance Native Go MCPs & Hooks"

$GoProjects = @(
    @{ Name = "mcp_git"; Path = Join-Path $RepoRoot "go_mcp\git"; Output = Join-Path $BinDir "mcp_git.exe" },
    @{ Name = "mcp_mem0_bridge"; Path = Join-Path $RepoRoot "go_mcp\mem0_bridge"; Output = Join-Path $BinDir "mcp_mem0_bridge.exe" },
    @{ Name = "mcp_memory"; Path = Join-Path $RepoRoot "go_mcp\memory"; Output = Join-Path $BinDir "mcp_memory.exe" },
    @{ Name = "mcp_sqlite"; Path = Join-Path $RepoRoot "go_mcp\sqlite"; Output = Join-Path $BinDir "mcp_sqlite.exe" },
    @{ Name = "aider_guard"; Path = Join-Path $RepoRoot "go_hooks\aider_guard"; Output = Join-Path $BinDir "aider_guard.exe" }
)

foreach ($gp in $GoProjects) {
    if ((Test-Path $gp.Output) -and (-not $ForceRebuild)) {
        Write-Step "$($gp.Name) already compiled (skipping, use -ForceRebuild to recompile)" "CACHED"
        continue
    }
    Write-Host "  [*] Compiling $($gp.Name)..." -ForegroundColor White
    Push-Location $gp.Path
    try {
        & $GoExe build -ldflags="-s -w" -o $gp.Output .
        Write-Step "$($gp.Name) compiled to $($gp.Output)"
    } catch {
        Write-ErrMsg "Failed to compile $($gp.Name): $_"
        exit 1
    } finally {
        Pop-Location
    }
}

# 4. Copy Scripts, Directives, Skills & Plugins
Write-Header "Step 4: Synchronizing Scripts, Directives, Skills & Plugins"

# Copy Scripts
Get-ChildItem -Path (Join-Path $RepoRoot "scripts") -File | ForEach-Object {
    Copy-Item $_.FullName -Destination $ScriptsDir -Force
}
Write-Step "Python & PowerShell engine scripts installed"

# Copy Global Directives
Copy-Item (Join-Path $RepoRoot "configs\AGENTS.md") -Destination (Join-Path $ConfigDir "AGENTS.md") -Force
Copy-Item (Join-Path $RepoRoot "configs\GEMINI.md") -Destination (Join-Path $ConfigDir "GEMINI.md") -Force
Write-Step "Directives AGENTS.md & GEMINI.md installed"

# Copy Skills
Copy-Item -Path (Join-Path $RepoRoot "skills\*") -Destination $SkillsDir -Recurse -Force
Write-Step "Custom skills installed (invisible-playwright, skill-synthesizer, etc.)"

# Copy Plugins
Copy-Item -Path (Join-Path $RepoRoot "plugins\*") -Destination $PluginsDir -Recurse -Force
Write-Step "Plugins installed (agentic-engineering, elite-workflows, open-code-review, etc.)"

# 5. Generate Configuration Files (mcp_config.json & hooks.json)
Write-Header "Step 5: Generating System Configuration Files"

# Template Replacements
$mcpTemplate = Get-Content (Join-Path $RepoRoot "configs\mcp_config.template.json") -Raw
$hooksTemplate = Get-Content (Join-Path $RepoRoot "configs\hooks.template.json") -Raw

$escapedBin = $BinDir.Replace("\", "\\")
$escapedScripts = $ScriptsDir.Replace("\", "\\")
$escapedPython = $PythonExe.Replace("\", "\\")
$escapedAntigravity = $AntigravityDir.Replace("\", "\\")

$mcpRendered = $mcpTemplate `
    -replace "\{\{BIN_DIR\}\}", $escapedBin `
    -replace "\{\{SCRIPTS_DIR\}\}", $escapedScripts `
    -replace "\{\{PYTHON_EXE\}\}", $escapedPython `
    -replace "\{\{ANTIGRAVITY_DIR\}\}", $escapedAntigravity `
    -replace "\{\{EXE_EXT\}\}", ".exe"

$hooksRendered = $hooksTemplate `
    -replace "\{\{BIN_DIR\}\}", $BinDir.Replace("\", "/") `
    -replace "\{\{SCRIPTS_DIR\}\}", $ScriptsDir.Replace("\", "/") `
    -replace "\{\{PYTHON_CMD\}\}", "python" `
    -replace "\{\{EXE_EXT\}\}", ".exe"

$ActiveMcpConfig = Join-Path $ConfigDir "mcp_config.json"
$ActiveHooksConfig = Join-Path $ConfigDir "hooks.json"

Set-Content -Path $ActiveMcpConfig -Value $mcpRendered -Encoding UTF8
Set-Content -Path $ActiveHooksConfig -Value $hooksRendered -Encoding UTF8
Write-Step "Generated $ActiveMcpConfig"
Write-Step "Generated $ActiveHooksConfig"

# 6. Python Dependencies for Mem0 & AST Guards
if (-not $SkipPip) {
    Write-Header "Step 6: Installing Python Dependencies"
    Write-Host "  [*] Verifying and installing required packages (mcp, starlette, mem0ai, qdrant-client)..." -ForegroundColor White
    & $PythonExe -m pip install --quiet --upgrade pip
    & $PythonExe -m pip install --quiet mcp starlette uvicorn mem0ai qdrant-client sentence-transformers python-dotenv
    Write-Step "Python dependencies successfully verified"
}

# 7. Mem0 Background Service & Autostart
Write-Header "Step 7: Mem0 Local Memory Background Service"
if ($EnableAutoStart) {
    $Mem0Ctl = Join-Path $ScriptsDir "mem0_ctl.ps1"
    if (Test-Path $Mem0Ctl) {
        Write-Host "  [*] Registering Mem0 autostart on Windows boot/logon..." -ForegroundColor White
        & powershell.exe -ExecutionPolicy Bypass -File $Mem0Ctl -Action install-autostart
        Write-Host "  [*] Starting Mem0 service daemon..." -ForegroundColor White
        & powershell.exe -ExecutionPolicy Bypass -File $Mem0Ctl -Action restart
        Write-Step "Mem0 service registered and started on port 8765"
    }
}

Write-Header "INSTALLATION COMPLETE - YOUR AGENT IS SUPERCHARGED!"
Write-Host "
Summary of installed enhancements:
  * Master Directives : $ConfigDir\AGENTS.md & GEMINI.md
  * High-Speed MCPs   : Git, Mem0 Bridge, Memory, SQLite in compiled Go ($BinDir)
  * Lifecycle Hooks   : Native Impact Guard (AST Blast Radius) & Aider Guard
  * Local Memory      : Mem0 background SSE daemon on port 8765 (Anti-Dilution project scope)
  * Elite Skills      : 2 specialized skills & 11 production plugins
  * Config Path       : $ConfigDir\mcp_config.json

Next steps:
  1. Restart your Google Antigravity IDE / CLI to load new MCPs and hooks.
  2. Enjoy unmatched agent speed, memory, and architectural precision!
" -ForegroundColor Green
