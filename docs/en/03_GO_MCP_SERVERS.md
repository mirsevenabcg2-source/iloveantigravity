# ⚡ Native Go MCP Servers: High-Throughput & Zero Lag

## 1. Why Rewrite MCP Servers in Go?
Standard Model Context Protocol (MCP) servers built with Node.js (`npx`) or Python introduce notable friction:
1. **Startup Latency**: Node and Python processes take 300ms–1500ms to boot up for each tool invocation.
2. **Memory Footprint**: Each Node.js process takes ~80MB–150MB of RAM; running 6-8 MCP servers consumes over 1GB of memory just for idle adapters.
3. **Interpreter Dependencies**: Broken npm packages, python version mismatches, or missing C-extensions create maintenance headaches across machines.

**Go solves all three problems**:
- Compiles into a single standalone binary without external runtime dependencies.
- Sub-millisecond cold start times.
- Extremely lean memory footprint (typically <10MB RAM under full load).

---

## 2. Server Breakdown

### `mcp_git` (`go_mcp/git/main.go`)
Provides rapid Git operations directly to the agent without shell spawning overhead:
- `git_status`: Working tree status.
- `git_diff_unstaged`: Unstaged working directory changes with configurable context lines.
- `git_diff_staged`: Staged index changes.
- `git_diff`: Branch/commit diffs.
- `git_commit`: Atomic commits.
- `git_add`: Stage files.
- `git_reset`: Unstage files.
- `git_log`: Commit history inspection.
- `git_create_branch`, `git_checkout`, `git_show`, `git_branch`.

### `mcp_mem0_bridge` (`go_mcp/mem0_bridge/main.go`)
Acts as a lightning-fast client bridge to the local Mem0 background service:
- Exposes tools: `mem0_status`, `mem0_add`, `mem0_search`, `mem0_get_all`, `mem0_delete`.
- Auto-healing: If the local background service is not running, the bridge automatically starts `mem0_service.py` via `pythonw.exe` and waits for port binding before completing the tool request.

### `mcp_memory` (`go_mcp/memory/main.go`)
A persistent, thread-safe in-memory knowledge graph MCP server:
- Stores entities and relations in a structured JSON file at the project root (`memory.json`).
- Exposes tools: `create_entities`, `create_relations`, `add_observations`, `delete_entities`, `delete_relations`, `read_graph`, `search_nodes`, `open_nodes`.

### `mcp_sqlite` (`go_mcp/sqlite/main.go`)
Direct, high-performance SQLite database inspector powered by modern Go pure-sqlite drivers:
- Auto-detects SQLite databases in the current project (`dev.db`, `app.db`, `data.db`, etc.).
- Exposes tools: `read_query`, `write_query`, `create_table`, `list_tables`, `describe_table`, `append_insight`.

---

## 3. Building From Source
All Go MCPs can be compiled across Windows, Linux, and macOS with standard Go toolchain:
```bash
# Compile all MCPs
cd go_mcp/git && go build -ldflags="-s -w" -o ../../bin/mcp_git .
cd go_mcp/mem0_bridge && go build -ldflags="-s -w" -o ../../bin/mcp_mem0_bridge .
cd go_mcp/memory && go build -ldflags="-s -w" -o ../../bin/mcp_memory .
cd go_mcp/sqlite && go build -ldflags="-s -w" -o ../../bin/mcp_sqlite .
```
The automated `install.ps1` and `install.sh` scripts handle this step automatically during setup.
