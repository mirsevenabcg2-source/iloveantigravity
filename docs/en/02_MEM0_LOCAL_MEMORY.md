# 🧠 Mem0 Local Memory Architecture & Anti-Dilution Isolation

## 1. Motivation: Why Local Vector Memory?
Standard agent sessions lose conversational memory once closed. External cloud memory solutions often:
1. Leak proprietary code and sensitive keys to third parties.
2. Incur recurrent API subscription fees.
3. Lack strict project boundary separation, leading to catastrophic context contamination (e.g., mixing game-reversing conventions with enterprise ERP database logic).

## 2. The iLoveAntigravity Local Mem0 Stack
Our Mem0 architecture runs as a permanent, high-speed Windows/Linux daemon:
- **Port & Transport**: `http://127.0.0.1:8765/sse` via Starlette and FastMCP.
- **Vector Database**: Embedded Qdrant located at `~/.gemini/antigravity/mem0_data/qdrant`.
- **Embedding Engine**: Local 768-dimensional cosine vector embeddings.
- **Inference Engine**: Direct integration with the **Google Antigravity Python SDK (`google.antigravity.Agent`)**.
  - **Zero Local VRAM**: Does NOT run heavy local LLMs like Ollama or LM Studio.
  - Lightning-fast extraction and self-updating conflict resolution.
- **Daemon Lifecycle**: Automatically starts on Windows login via Registry `Run` (`HKCU:\Software\Microsoft\Windows\CurrentVersion\Run\Mem0_Antigravity_Service`). Controlled via `scripts/mem0_ctl.ps1`.

---

## 3. The Anti-Dilution Protocol (Zero-Unscoped Rule)
To prevent cross-repository memory contamination, every agent MUST obey the two-layer protection protocol:

### Layer 1: Mandatory Project Scope Parameter
Calling `mem0_search` or `mem0_add` without specifying the `project` argument is strictly prohibited:
```python
# Tra cứu thói quen của dự án cụ thể
mem0_search(query="Database migration conventions", project="DSCons")

# Tra cứu sở thích dev toàn cục
mem0_search(query="Preferred indent and comment style", project="Global")
```

### Layer 2: Double-Layer Structured Prefix
When saving memory with `mem0_add`, content must always feature a structured category header:
```python
mem0_add(
    text="[BackendService: Architecture] Database migrations must be backwards-compatible and tested on staging.",
    project="BackendService"
)
```
This dual-layer boundary physically segregates vectors in Qdrant, preventing cross-project memory bleed.

---

## 4. Managing the Service
Using the included `scripts/mem0_ctl.ps1`:
```powershell
# Check running status and port binding
.\scripts\mem0_ctl.ps1 -Action status

# Start daemon
.\scripts\mem0_ctl.ps1 -Action start

# Restart daemon
.\scripts\mem0_ctl.ps1 -Action restart

# Stop daemon
.\scripts\mem0_ctl.ps1 -Action stop

# Enable automatic start on Windows boot
.\scripts\mem0_ctl.ps1 -Action install-autostart
```
