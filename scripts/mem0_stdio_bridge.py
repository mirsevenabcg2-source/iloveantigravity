import sys
import json
import socket
import asyncio
import subprocess
from pathlib import Path
from typing import Optional, Union, Dict, Any
from mcp.server.fastmcp import FastMCP
from mcp.client.sse import sse_client
from mcp.client.session import ClientSession

# Suppress console reconfigure issues
if hasattr(sys.stdout, "reconfigure"):
    sys.stdout.reconfigure(encoding="utf-8")
if hasattr(sys.stderr, "reconfigure"):
    sys.stderr.reconfigure(encoding="utf-8")

mcp = FastMCP("mem0")

HOST = "127.0.0.1"
PORT = 8765
SSE_URL = f"http://{HOST}:{PORT}/sse"
PYTHONW_EXE = r"C:\Users\Admin\AppData\Local\Programs\Python\Python311\pythonw.exe"
SERVICE_SCRIPT = r"C:\Users\Admin\.gemini\config\scripts\mem0_service.py"
DATA_DIR = r"C:\Users\Admin\.gemini\antigravity\mem0_data"

def is_service_port_open(timeout: float = 0.5) -> bool:
    """Check if the Mem0 service port is actively accepting connections."""
    try:
        with socket.create_connection((HOST, PORT), timeout=timeout):
            return True
    except (socket.timeout, ConnectionRefusedError, OSError):
        return False

def spawn_background_service() -> bool:
    """Spawns the Mem0 background service decoupled from any parent terminal or job object."""
    # Attempt 1: Windows WMI Process Creation (completely decoupled from job objects)
    cmd = f'"{PYTHONW_EXE}" "{SERVICE_SCRIPT}"'
    wmi_cmd = f'Invoke-CimMethod -ClassName Win32_Process -MethodName Create -Arguments @{{CommandLine=\'{cmd}\'; CurrentDirectory=\'{DATA_DIR}\'}}'
    try:
        res = subprocess.run(
            ["powershell", "-NoProfile", "-NonInteractive", "-ExecutionPolicy", "Bypass", "-Command", wmi_cmd],
            capture_output=True,
            timeout=5
        )
        if res.returncode == 0:
            return True
    except Exception:
        pass

    # Attempt 2: Direct detached process
    try:
        # DETACHED_PROCESS = 0x00000008, CREATE_NEW_PROCESS_GROUP = 0x00000200, CREATE_BREAKAWAY_FROM_JOB = 0x01000000
        flags = 0x00000008 | 0x00000200 | 0x01000000
        subprocess.Popen(
            [PYTHONW_EXE, SERVICE_SCRIPT],
            cwd=DATA_DIR,
            creationflags=flags,
            close_fds=True
        )
        return True
    except Exception:
        pass

    return False

async def ensure_service_ready(max_retries: int = 15, delay: float = 0.5) -> bool:
    """Ensures the Mem0 background service is running and ready. Auto-heals if stopped."""
    if is_service_port_open(timeout=0.3):
        return True

    # Port is not open, spawn the service
    spawn_background_service()

    # Poll until ready
    for _ in range(max_retries):
        await asyncio.sleep(delay)
        if is_service_port_open(timeout=0.3):
            return True

    return False

async def _call_remote_tool(name: str, args: dict) -> str:
    # 1. Proactive Self-Healing: check and auto-start if service is down
    service_ready = await ensure_service_ready()
    if not service_ready:
        return json.dumps({
            "status": "error",
            "message": f"Failed to start Mem0 background service at {SSE_URL}. Service did not bind to port {PORT} in time.",
            "hint": "Check C:\\Users\\Admin\\.gemini\\antigravity\\mem0_data\\mem0_service.log for startup errors."
        }, indent=2)

    # 2. Execute call with 1 auto-retry in case of transient disconnect
    for attempt in range(2):
        try:
            async with sse_client(SSE_URL) as (read, write):
                async with ClientSession(read, write) as session:
                    await session.initialize()
                    result = await session.call_tool(name, args)
                    parts = []
                    for content in result.content:
                        if hasattr(content, "text"):
                            parts.append(content.text)
                        else:
                            parts.append(str(content))
                    return "\n".join(parts)
        except Exception as e:
            if attempt == 0:
                # Give service a moment to recover or restart
                await asyncio.sleep(1.0)
                await ensure_service_ready()
                continue
            return json.dumps({
                "status": "error",
                "message": f"Failed to connect to Mem0 background service at {SSE_URL}: {str(e)}",
                "hint": "Ensure mem0 service is running via mem0_ctl.ps1 status"
            }, indent=2)

@mcp.tool()
async def mem0_status() -> str:
    """Check the operational health and statistics of the Mem0 memory engine."""
    return await _call_remote_tool("mem0_status", {})

@mcp.tool()
async def mem0_add(
    text: str,
    user_id: str = "developer",
    metadata: Optional[Union[Dict[str, Any], str]] = None,
    project: Optional[str] = None
) -> str:
    """Store a new fact, user preference, architectural decision, or convention into persistent memory.
    
    Args:
        text: The fact or decision to remember.
        user_id: The user or entity identifier (default: 'developer').
        metadata: Optional JSON string or dictionary containing additional context metadata.
        project: Target project name (e.g. 'FreeExile', 'DSCons', 'JX', 'Global'). MANDATORY: Always specify project to prevent memory dilution.
    """
    args = {"text": text, "user_id": user_id}
    if project is not None:
        args["project"] = project
    if metadata is not None:
        if isinstance(metadata, dict):
            args["metadata"] = json.dumps(metadata, ensure_ascii=False)
        else:
            args["metadata"] = str(metadata)
    return await _call_remote_tool("mem0_add", args)

@mcp.tool()
async def mem0_search(
    query: str,
    user_id: str = "developer",
    limit: int = 5,
    project: Optional[str] = None
) -> str:
    """Search long-term memories for relevant facts, preferences, or technical conventions matching the query.
    
    Args:
        query: Search query (e.g. 'coding conventions', 'architecture choices', 'user preferences').
        user_id: User identifier to filter memories (default: 'developer').
        limit: Maximum number of memories to return (default: 5).
        project: Target project filter (e.g. 'FreeExile', 'DSCons', 'Global'). MANDATORY: Always filter by project to prevent cross-project memory contamination.
    """
    args = {"query": query, "user_id": user_id, "limit": limit}
    if project is not None:
        args["project"] = project
    return await _call_remote_tool("mem0_search", args)

@mcp.tool()
async def mem0_get_all(
    user_id: str = "developer"
) -> str:
    """Retrieve all stored memories and facts for a given user.
    
    Args:
        user_id: The user ID to retrieve memories for (default: 'developer').
    """
    return await _call_remote_tool("mem0_get_all", {"user_id": user_id})

@mcp.tool()
async def mem0_delete(
    memory_id: str
) -> str:
    """Delete a specific memory by its unique ID.
    
    Args:
        memory_id: The unique ID of the memory to delete.
    """
    return await _call_remote_tool("mem0_delete", {"memory_id": memory_id})

if __name__ == "__main__":
    mcp.run(transport="stdio")
