# 🔄 Cross-Machine Migration & Turnkey Setup Guide

## 1. Overview
Whenever you switch to a new laptop, desktop PC, cloud VM, or reinstall your operating system, **iLoveAntigravity** allows you to replicate your entire tailored agent development environment in under 60 seconds.

---

## 2. Step-by-Step Migration

### Step 1: Install Base Prerequisites
Ensure your target machine has:
1. **Git**: [https://git-scm.com/](https://git-scm.com/)
2. **Go Compiler (1.22+)**: [https://go.dev/dl/](https://go.dev/dl/)
3. **Python (3.10+)**: [https://python.org/](https://python.org/) (ensure "Add python.exe to PATH" is checked)
4. **Node.js (18+)**: [https://nodejs.org/](https://nodejs.org/)

### Step 2: Clone the Repository
```bash
git clone https://github.com/mirsevenabcg2-source/iloveantigravity.git
cd iloveantigravity
```

### Step 3: Run the Turnkey Installer

#### On Windows:
```powershell
# Open PowerShell (Administrator recommended for registry auto-start)
.\install.ps1
```

#### On Linux or macOS:
```bash
chmod +x install.sh
./install.sh
```

### Step 4: Configure Environment (Optional)
The entire suite runs completely offline/locally without requiring any external API keys. If you want to customize local ports, you can create a `.env` file in `~/.gemini/antigravity/mem0_data/.env`:
```bash
MEM0_PORT=8765
MEM0_HOST=127.0.0.1
```

### Step 5: Verification & Launch
1. Open or restart **Google Antigravity IDE** or CLI.
2. In the chat interface, run `/help` or check the loaded MCP servers in Settings. You will see:
   - `git` (Native Go)
   - `mem0` (Go bridge to local vector engine)
   - `memory` (Native Go Knowledge Graph)
   - `sqlite` (Native Go DB inspector)
   - `ast-grep` & `lsp-mcp`
   - `invisible-playwright`
3. Verify Mem0 background daemon health:
   ```powershell
   powershell -File ~/.gemini/config/scripts/mem0_ctl.ps1 -Action status
   ```
4. You are ready to build at maximum velocity with zero context loss!
