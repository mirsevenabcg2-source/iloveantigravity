# Cleanup script for Antigravity maintenance
$cutoff = (Get-Date).AddDays(-3)

# 1. Clean task logs older than 3 days
$brainPath = "$env:USERPROFILE\.gemini\antigravity\brain"
$deletedLogs = 0
if (Test-Path $brainPath) {
    $oldLogs = Get-ChildItem -Path $brainPath -Recurse -Filter "task-*.log" -ErrorAction SilentlyContinue |
        Where-Object { $_.LastWriteTime -lt $cutoff }
    $deletedLogs = ($oldLogs | Measure-Object).Count
    $oldLogs | Remove-Item -Force -ErrorAction SilentlyContinue
}

# 2. Clean 0-byte crash logs
$crashesPath = "$env:USERPROFILE\.gemini\antigravity\crashes"
$deletedCrashes = 0
if (Test-Path $crashesPath) {
    $emptyCrashes = Get-ChildItem -Path $crashesPath -Filter "*.log" -ErrorAction SilentlyContinue |
        Where-Object { $_.Length -eq 0 }
    $deletedCrashes = ($emptyCrashes | Measure-Object).Count
    $emptyCrashes | Remove-Item -Force -ErrorAction SilentlyContinue
}

# 3. Clean temporary cache
$cachePath = "$env:USERPROFILE\.gemini\antigravity\cache"
if (Test-Path $cachePath) {
    Get-ChildItem -Path $cachePath -Recurse -ErrorAction SilentlyContinue |
        Remove-Item -Recurse -Force -ErrorAction SilentlyContinue
}

[PSCustomObject]@{
    "Deleted Task Logs (>3d)" = $deletedLogs
    "Deleted Empty Crashes"   = $deletedCrashes
    "Status"                  = "Completed"
} | Format-Table
