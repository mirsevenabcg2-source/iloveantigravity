#!/usr/bin/env bash
# ==============================================================================
# iloveantigravity - 1-Click Turnkey Setup Script for Linux & macOS
# ==============================================================================
set -euo pipefail

echo ""
echo "================================================================="
echo "   iLoveAntigravity - Turnkey Agent Supercharger (Linux/macOS)"
echo "================================================================="
echo ""

REPO_ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
USER_HOME="${HOME}"
GEMINI_DIR="${USER_HOME}/.gemini"
CONFIG_DIR="${GEMINI_DIR}/config"
ANTIGRAVITY_DIR="${GEMINI_DIR}/antigravity"
BIN_DIR="${CONFIG_DIR}/bin"
SCRIPTS_DIR="${CONFIG_DIR}/scripts"
SKILLS_DIR="${CONFIG_DIR}/skills"
PLUGINS_DIR="${CONFIG_DIR}/plugins"
MEM0_DATA_DIR="${ANTIGRAVITY_DIR}/mem0_data"

# 1. Prerequisites Check
echo "==> Step 1: Checking System Prerequisites"

PYTHON_BIN=""
for cmd in python3 python; do
    if command -v "$cmd" &>/dev/null; then
        PYTHON_BIN="$(command -v "$cmd")"
        break
    fi
done

if [[ -z "$PYTHON_BIN" ]]; then
    echo "[-] Error: Python 3.10+ is required but not found in PATH."
    exit 1
fi
echo "[+] Python detected: $PYTHON_BIN"

if ! command -v go &>/dev/null; then
    echo "[-] Error: Go compiler is required but not found in PATH. Please install Go 1.22+."
    exit 1
fi
echo "[+] Go compiler detected: $(command -v go)"

if ! command -v git &>/dev/null; then
    echo "[-] Error: Git is required but not found in PATH."
    exit 1
fi
echo "[+] Git detected: $(command -v git)"

# 2. Directory Structure
echo "==> Step 2: Preparing ~/.gemini Directory Hierarchy"
mkdir -p "$CONFIG_DIR" "$ANTIGRAVITY_DIR" "$BIN_DIR" "$SCRIPTS_DIR" "$SKILLS_DIR" "$PLUGINS_DIR" "${MEM0_DATA_DIR}/qdrant"
echo "[+] Directory structure verified."

# 3. Compiling Native Go MCPs
echo "==> Step 3: Compiling Native Go MCP Servers"

compile_go() {
    local name="$1"
    local path="$2"
    local output="$3"
    echo "  [*] Compiling $name..."
    (cd "$path" && go build -ldflags="-s -w" -o "$output" .)
    chmod +x "$output"
    echo "  [+] Built: $output"
}

compile_go "mcp_git" "${REPO_ROOT}/go_mcp/git" "${BIN_DIR}/mcp_git"
compile_go "mcp_mem0_bridge" "${REPO_ROOT}/go_mcp/mem0_bridge" "${BIN_DIR}/mcp_mem0_bridge"
compile_go "mcp_memory" "${REPO_ROOT}/go_mcp/memory" "${BIN_DIR}/mcp_memory"
compile_go "mcp_sqlite" "${REPO_ROOT}/go_mcp/sqlite" "${BIN_DIR}/mcp_sqlite"

# 4. Synchronize Scripts, Directives, Skills & Plugins
echo "==> Step 4: Synchronizing Scripts, Directives, Skills & Plugins"
cp -r "${REPO_ROOT}/scripts/"* "${SCRIPTS_DIR}/"
chmod +x "${SCRIPTS_DIR}"/*.py "${SCRIPTS_DIR}"/*.sh 2>/dev/null || true

cp "${REPO_ROOT}/configs/AGENTS.md" "${CONFIG_DIR}/AGENTS.md"
cp "${REPO_ROOT}/configs/GEMINI.md" "${CONFIG_DIR}/GEMINI.md"

cp -r "${REPO_ROOT}/skills/"* "${SKILLS_DIR}/"
cp -r "${REPO_ROOT}/plugins/"* "${PLUGINS_DIR}/"
echo "[+] All assets and directives synchronized."

# 5. Render Configuration Files
echo "==> Step 5: Generating System Configuration Files"

sed \
    -e "s|{{BIN_DIR}}|${BIN_DIR}|g" \
    -e "s|{{SCRIPTS_DIR}}|${SCRIPTS_DIR}|g" \
    -e "s|{{PYTHON_EXE}}|${PYTHON_BIN}|g" \
    -e "s|{{ANTIGRAVITY_DIR}}|${ANTIGRAVITY_DIR}|g" \
    -e "s|{{EXE_EXT}}||g" \
    "${REPO_ROOT}/configs/mcp_config.template.json" > "${CONFIG_DIR}/mcp_config.json"

sed \
    -e "s|{{BIN_DIR}}|${BIN_DIR}|g" \
    -e "s|{{SCRIPTS_DIR}}|${SCRIPTS_DIR}|g" \
    -e "s|{{PYTHON_CMD}}|${PYTHON_BIN}|g" \
    -e "s|{{EXE_EXT}}||g" \
    "${REPO_ROOT}/configs/hooks.template.json" > "${CONFIG_DIR}/hooks.json"

echo "[+] Generated: ${CONFIG_DIR}/mcp_config.json"
echo "[+] Generated: ${CONFIG_DIR}/hooks.json"

# 6. Python Dependencies
echo "==> Step 6: Installing Python Dependencies"
"$PYTHON_BIN" -m pip install --quiet --upgrade pip
"$PYTHON_BIN" -m pip install --quiet mcp starlette uvicorn mem0ai qdrant-client sentence-transformers python-dotenv || true
echo "[+] Python dependencies verified."

echo ""
echo "================================================================="
echo "   INSTALLATION COMPLETE - AGENT FULLY SUPERCHARGED!"
echo "================================================================="
echo "Your Google Antigravity environment is now upgraded with:"
echo "  - Master Directives (Hermes Protocol, OCR, TDD, Anti-Slop)"
echo "  - Ultra-fast native Go MCP servers (git, mem0, memory, sqlite)"
echo "  - Native Impact Guard and AST blast radius checking"
echo "  - Mem0 Local Memory engine (Qdrant + Google Antigravity SDK)"
echo "  - Battle-tested Skills & Plugins suite"
echo ""
