# 🚀 iLoveAntigravity

> **The Ultimate Battle-Tested Supercharger for Google Antigravity & Autonomous AI Coding Agents.**  
> *Turnkey setup, instant cross-machine migration, zero context dilution, blazing-fast native Go MCP servers, local vector memory, and deterministic engineering protocols.*

---

### 🌐 Multi-Language Documentation
[ English ](README.md) • [ Tiếng Việt ](README.vi.md) • [ 简体中文 ](README.zh.md) • [ 日本語 ](README.ja.md)

---

## ⚡ What is iLoveAntigravity?

**iLoveAntigravity** is a unified, turnkey optimization suite designed to transform [Google Antigravity](https://github.com/google/antigravity) from a standard agent into an ultra-high-performance autonomous software engineering partner. 

Created from real-world, high-intensity production experience, this repository packs:
- **Instant Cross-Machine Migration**: Clone and run 1 command on any computer (Windows, Linux, macOS) to restore your entire customized environment in under 60 seconds.
- **The Hermes Protocol**: Bounded memory snapshots (<2,500 chars), proactive context compaction at 50%, and sub-millisecond local recall to eliminate token bloat and context drift.
- **Native Go-Compiled MCP Servers**: High-throughput, zero-latency Model Context Protocol servers for Git, Mem0, SQLite, and Knowledge Graphs written in Go.
- **Mem0 Local Memory (Zero Dilution)**: An embedded Qdrant vector memory service driven directly by the **Google Antigravity Python SDK**, running locally with 0% VRAM consumption and strict multi-project boundary isolation.
- **Alibaba Open Code Review (OCR) Delegation Mode**: 5-step deterministic code review workflow that eradicates line-offset hallucinations and API costs.
- **Native Lifecycle Hooks & Blast Radius Protection**: PreToolUse AST blast-radius analyzer and PostToolUse syntax verification & language server memory trimmer.
- **Elite Skills & Plugins**: Production-grade tools for stealth anti-detect browser automation (`invisible-playwright`), autonomous skill synthesis (`skill-synthesizer`), closed-loop TDD, and Git worktree isolation.

---

## 🚀 Quick Start (1-Click Migration)

Whenever you switch to a new laptop, PC, or server, you can restore your complete optimized environment with a single command:

### On Windows (PowerShell):
```powershell
# 1. Clone repository
git clone https://github.com/mirsevenabcg2-source/iloveantigravity.git
cd iloveantigravity

# 2. Run Turnkey Installer
.\install.ps1
```

### On Linux / macOS (Bash):
```bash
# 1. Clone repository
git clone https://github.com/mirsevenabcg2-source/iloveantigravity.git
cd iloveantigravity

# 2. Make executable and run
chmod +x install.sh
./install.sh
```

The installer will:
1. Verify system prerequisites (`Python 3.10+`, `Go 1.22+`, `Node.js`, `Git`).
2. Initialize directory hierarchies in `~/.gemini/config` and `~/.gemini/antigravity`.
3. Compile all native Go MCP servers (`mcp_git`, `mcp_mem0_bridge`, `mcp_memory`, `mcp_sqlite`, `aider_guard`) into native machine code.
4. Render parameterized configurations (`mcp_config.json`, `hooks.json`) resolving exact paths for your OS.
5. Install Python dependencies (`mem0ai`, `qdrant-client`, `sentence-transformers`, `mcp`, `starlette`).
6. Register the local Mem0 background service to auto-start on system boot/logon.

---

## 🏛️ Architectural Pillars

```
                                ┌────────────────────────┐
                                │   Google Antigravity   │
                                │    Master Host Agent   │
                                └───────────┬────────────┘
                                            │
               ┌────────────────────────────┼────────────────────────────┐
               ▼                            ▼                            ▼
    ┌──────────────────────┐     ┌──────────────────────┐     ┌──────────────────────┐
    │   Hermes Protocol    │     │   Native Go MCPs     │     │ Local Mem0 Service   │
    │  - Bounded Snapshot  │     │  - mcp_git           │     │  - Port 8765 SSE     │
    │  - 50% Compaction    │     │  - mcp_mem0_bridge   │     │  - Qdrant Embedded   │
    │  - Skill Synthesis   │     │  - mcp_memory        │     │  - AGY Python SDK    │
    │  - Anti-AI Slop      │     │  - mcp_sqlite        │     │  - Anti-Dilution     │
    └──────────────────────┘     └──────────────────────┘     └──────────────────────┘
               │                            │                            │
               ▼                            ▼                            ▼
    ┌──────────────────────┐     ┌──────────────────────┐     ┌──────────────────────┐
    │  Lifecycle Hooks     │     │  Hybrid Code Review  │     │ Battle-Tested Skills │
    │  - Impact Guard      │     │  - Alibaba OCR       │     │  - Stealth Browser   │
    │  - Blast Radius AST  │     │  - Delegation Mode   │     │  - Skill Synthesis   │
    │  - Aider Mem Trim    │     │  - 5-Step Checklist  │     │  - Worktree & TDD    │
    └──────────────────────┘     └──────────────────────┘     └──────────────────────┘
```

### 1. The Hermes Protocol (Bounded Memory & Context Compaction)
- **Bounded Snapshot (<2,500 chars)**: Keeps global System Prompt lean, preventing initial context bloat. Deep historical context is retrieved on-demand via Mem0 or SQLite FTS5.
- **Proactive Context Compaction @ 50%**: Compresses oversized tool output blocks (build logs, large diffs) into clean 2-3 line summaries at 50% window utilization.
- **Autonomous Skill Synthesis (`skill-synthesizer`)**: Self-reflects after solving difficult technical tasks and synthesizes reusable skills conforming to the `agentskills.io` standard.
- **Local Fast Recall**: Sub-2ms recall via SQLite FTS5 and ripgrep, with 0 token consumption and 0% VRAM overhead.

### 2. Mem0 Local Background Service & Zero Dilution Rule
- **Self-Contained Local Daemon**: Operates on `http://127.0.0.1:8765/sse` via Starlette/FastMCP, auto-starting on Windows boot/logon via Registry `Run`.
- **Zero Local LLM Bloat**: Uses the official **Google Antigravity Python SDK (`google.antigravity.Agent`)** directly for memory extraction, requiring NO heavy local LLM runtimes (LM Studio, Ollama).
- **Strict Anti-Dilution (Multi-Project Isolation)**:
  - **Zero-Unscoped Rule**: `mem0_search` and `mem0_add` MUST explicitly declare `project="<ProjectName>"` or `project="Global"`.
  - **Double-Layer Guard**: Every stored memory is prefixed with `[<ProjectName>: <Category>] Content...` to permanently eradicate cross-project context contamination.

### 3. High-Performance Native Go MCP Servers
Standard Node.js or Python MCP servers can introduce significant latency and memory overhead. This repository replaces them with compiled native Go executables:
- **`mcp_git`**: Zero-allocation Git operations (status, staged/unstaged diffs, branch management, commit inspection).
- **`mcp_mem0_bridge`**: Persistent, auto-healing SSE bridge to the local Mem0 service.
- **`mcp_memory`**: Blazing-fast in-memory entity-relation graph store.
- **`mcp_sqlite`**: Direct SQLite database inspection, query execution, and schema analysis.

### 4. Hybrid Code Review with Alibaba Open Code Review (OCR)
- Adopts Alibaba's deterministic delegation mode (`ocr delegate`).
- **5-Step Workflow**:
  1. `ocr delegate preview`: Mechanically extracts modified files, computes merge-base, filters locks/vendor files.
  2. `ocr delegate rule`: Pulls enterprise rule groups (NPE, thread-safety, SQLi, resource leaks).
  3. `git diff`: Fetches real unified diff against merge-base.
  4. **Host Agent Deep Audit**: Evaluates rules line-by-line, determining exact line numbers and severity (`critical`, `high`, `medium`, `low`).
  5. **Autonomous Auto-Fix**: Automatically patches High and Critical issues upon request.

### 5. Native Lifecycle Hooks & Blast Radius Protection
- **PreToolUse Hook (`native_impact_guard.py`)**: Intercepts file modification tools (`replace_file_content`, `write_to_file`). Uses AST parsing to analyze dependent modules. If a file impacts $\ge 6$ dependent modules, the edit is safely blocked until an explicit blast radius acknowledgment is provided.
- **PostToolUse Hook (`aider_guard`)**: Compiled Go hook that validates syntax for JSON (anti-BOM check) and Python, and automatically trims memory consumption of background language server processes.

---

## 📂 Repository Structure

```
iloveantigravity/
├── install.ps1                 # Turnkey Windows 1-Click Installer
├── install.sh                  # Turnkey Linux/macOS 1-Click Installer
├── .env.example                # Environment variables template
├── configs/
│   ├── AGENTS.md               # Master engineering directives
│   ├── GEMINI.md               # Model behavioral & persona guidelines
│   ├── hooks.template.json     # Lifecycle hooks configuration template
│   ├── mcp_config.template.json# Universal MCP server configuration template
│   └── workspace_template.json # Template for project-level workspace settings
├── go_mcp/                     # Native Go MCP Servers (Full Source Code)
│   ├── git/                    # High-throughput Git MCP
│   ├── mem0_bridge/            # Auto-healing Mem0 SSE client bridge
│   ├── memory/                 # In-Memory Knowledge Graph MCP
│   └── sqlite/                 # SQLite database MCP
├── go_hooks/                   # Native Go Lifecycle Hooks
│   └── aider_guard/            # Language server memory trimmer & syntax validator
├── scripts/                    # Core Python & PowerShell Daemons
│   ├── mem0_service.py         # Mem0 FastAPI/SSE server (Qdrant + AGY SDK)
│   ├── mem0_ctl.ps1            # Windows service controller (start/stop/status)
│   ├── mem0_stdio_bridge.py    # Stdio fallback bridge
│   ├── native_impact_guard.py  # AST Blast Radius analyzer
│   ├── markitdown_mcp.py       # Document to Markdown MCP
│   └── cleanup_antigravity.ps1 # Maintenance & cache cleaner
├── skills/                     # Specialized Skills (agentskills.io format)
│   ├── invisible-playwright/   # Stealth anti-detect Firefox C++ web automation
│   └── skill-synthesizer/      # Autonomous reflection & skill synthesizer
├── plugins/                    # 11 Curated Production Plugins
│   ├── agentic-engineering/    # Worktree delegation, TDD closed-loop, ast-grep
│   ├── developer-power-skills/ # Charting, Office docs, Vercel, problem-solving
│   ├── elite-workflows-and-taste/# High-end UI design, code review, brutalist UI
│   ├── open-code-review/       # Alibaba Open Code Review integration
│   ├── openspec-plugin/        # Spec-driven development workflows
│   └── science/                # AlphaFold, NCBI, ChEMBL, UniProt bioinformatics
└── docs/                       # In-Depth Guides in EN, VI, ZH, JA
```

---

## 🛠️ Configuration & Environment

The entire suite runs out-of-the-box with zero mandatory external API keys. To customize local daemon ports, copy `.env.example` to `.env` or set environment variables:

| Variable | Description | Default |
| :--- | :--- | :--- |
| `MEM0_PORT` | Local Mem0 SSE service port | `8765` |
| `MEM0_HOST` | Local Mem0 SSE host address | `127.0.0.1` |

---

## 📖 In-Depth Guides

- 📘 [Hermes Protocol & Bounded Memory](docs/en/01_HERMES_PROTOCOL.md)
- 🧠 [Local Mem0 Memory Architecture & Anti-Dilution](docs/en/02_MEM0_LOCAL_MEMORY.md)
- ⚡ [High-Speed Go MCP Servers Guide](docs/en/03_GO_MCP_SERVERS.md)
- 🔍 [Alibaba Open Code Review Delegation](docs/en/04_CODE_REVIEW_OCR.md)
- 🛡️ [Native Lifecycle Hooks & Blast Radius](docs/en/05_LIFECYCLE_HOOKS.md)
- 🔄 [Cross-Machine Migration Guide](docs/en/06_MIGRATION_GUIDE.md)

---

## 🙏 Credits & Acknowledgements

**iLoveAntigravity** is built upon and inspired by exceptional open-source projects, tools, and research from around the world. We gratefully acknowledge:
- **Alibaba Group** ([open-code-review](https://github.com/alibaba/open-code-review)) for deterministic code review delegation and rule engines.
- **Herrington Darkholme** ([ast-grep](https://github.com/ast-grep/ast-grep)) for structural AST syntax refactoring.
- **TheUpsider** ([@theupsider/lsp-mcp](https://github.com/theupsider/lsp-mcp)) for Language Server Protocol MCP integration.
- **Microsoft Corporation** ([MarkItDown](https://github.com/microsoft/markitdown), [Playwright](https://github.com/microsoft/playwright)) for document parsing and web automation.
- **Mem0.ai & Qdrant Team** ([mem0](https://github.com/mem0ai/mem0), [qdrant](https://github.com/qdrant/qdrant)) for vector memory paradigms and local vector search.
- **Google DeepMind** for the Antigravity SDK, Gemini foundation models, and AlphaFold/AlphaGenome systems.
- **Paul Gauthier** ([Aider](https://github.com/paul-gauthier/aider)) for syntax validation and developer flow inspiration.
- **Nous Research** for the Hermes Protocol bounded memory & self-reflection philosophy.
- **Vercel, Cloudflare, OpenSpec, and the gstack / Claude Code community** for curated skills, design patterns, and engineering workflows.

For the full list of authors, upstream repositories, citations, and licenses, see [CREDITS.md](CREDITS.md).

---

## 📄 License

This project is licensed under the [MIT License](LICENSE). Free for personal, commercial, and research use worldwide.