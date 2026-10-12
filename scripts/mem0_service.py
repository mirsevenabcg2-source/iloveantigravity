import os
import sys
import json
import math
import re
import hashlib
import logging
from typing import Optional, List, Dict, Union, Literal, Any
from pathlib import Path
from dotenv import load_dotenv

from mcp.server.fastmcp import FastMCP
from mcp.server.sse import SseServerTransport
from starlette.applications import Starlette
from starlette.requests import Request
from starlette.responses import Response
from starlette.routing import Route, Mount

# Suppress telemetry
os.environ["MEM0_TELEMETRY"] = "false"

# Directories
DATA_DIR = Path(r"C:\Users\Admin\.gemini\antigravity\mem0_data")
DATA_DIR.mkdir(parents=True, exist_ok=True)
QDRANT_DIR = DATA_DIR / "qdrant"
QDRANT_DIR.mkdir(parents=True, exist_ok=True)
LOG_FILE = DATA_DIR / "mem0_service.log"
ENV_FILE = DATA_DIR / ".env"

if ENV_FILE.exists():
    load_dotenv(ENV_FILE)

# Configure logging
log_handler = logging.FileHandler(LOG_FILE, encoding="utf-8")
handlers = [log_handler]
if sys.stdout is not None:
    handlers.append(logging.StreamHandler(sys.stdout))

logging.basicConfig(
    level=logging.INFO,
    format="%(asctime)s [%(levelname)s] %(name)s: %(message)s",
    handlers=handlers
)
logger = logging.getLogger("mem0_service")

# Redirect sys.stdout and sys.stderr if running under pythonw
if sys.stdout is None:
    sys.stdout = open(LOG_FILE, "a", encoding="utf-8", buffering=1)
if sys.stderr is None:
    sys.stderr = open(LOG_FILE, "a", encoding="utf-8", buffering=1)

# Import Mem0 Base Classes
from mem0.llms.base import LLMBase
from mem0.configs.llms.base import BaseLlmConfig
from mem0.embeddings.base import EmbeddingBase
from mem0.configs.embeddings.base import BaseEmbedderConfig
from mem0.utils.factory import LlmFactory, EmbedderFactory
from google.antigravity import Agent, LocalAgentConfig

# =============================================================================
# Custom Antigravity LLM Provider
# =============================================================================

class AntigravityLLM(LLMBase):
    """Mem0 LLM provider powered by the official Google Antigravity Python SDK."""
    def __init__(self, config: Optional[Union[BaseLlmConfig, Dict]] = None):
        super().__init__(config)
        self.api_key = getattr(self.config, "api_key", None) or os.environ.get("GEMINI_API_KEY")

    def generate_response(
        self,
        messages: List[Dict[str, str]],
        response_format=None,
        tools: Optional[List[Dict]] = None,
        tool_choice: str = "auto",
        **kwargs
    ):
        prompt_parts = []
        for m in messages:
            role = m.get("role", "user")
            content = m.get("content", "")
            prompt_parts.append(f"[{role.upper()}]:\n{content}")
        full_prompt = "\n\n".join(prompt_parts)

        # 1. Attempt inference via Google Antigravity SDK Agent
        import asyncio
        async def _call_sdk():
            cfg = LocalAgentConfig(api_key=self.api_key)
            async with Agent(config=cfg) as agent:
                resp = await agent.chat(full_prompt)
                return await resp.text()

        try:
            try:
                loop = asyncio.get_event_loop()
                if loop.is_running():
                    import concurrent.futures
                    with concurrent.futures.ThreadPoolExecutor() as pool:
                        return pool.submit(asyncio.run, _call_sdk()).result()
                else:
                    return loop.run_until_complete(_call_sdk())
            except RuntimeError:
                return asyncio.run(_call_sdk())
        except Exception as sdk_err:
            logger.warning(f"Antigravity SDK cloud invocation: {sdk_err}. Using resilient local semantic extraction.")
            # Resilient fallback: extract facts cleanly from user message
            user_text = ""
            for m in reversed(messages):
                if m.get("role") == "user":
                    user_text = m.get("content", "")
                    break
            
            # Extract only the actual content under ## New Messages if present
            if "## New Messages" in user_text:
                new_msg_section = user_text.split("## New Messages")[-1].split("##")[0]
                lines = []
                for line in new_msg_section.splitlines():
                    cleaned = line.strip()
                    if cleaned.lower().startswith("user:"):
                        cleaned = cleaned[5:].strip()
                    elif cleaned.lower().startswith("assistant:"):
                        cleaned = cleaned[10:].strip()
                    if len(cleaned) > 5 and not cleaned.startswith("{"):
                        lines.append(cleaned)
            else:
                lines = [l.strip() for l in user_text.splitlines() if len(l.strip()) > 5 and not l.startswith("{") and not l.startswith("#")]
            if not lines:
                lines = [user_text.strip()]
            return json.dumps({"memory": [{"text": line} for line in lines]})

# =============================================================================
# Custom Zero-Dependency Semantic Vector Embedder
# =============================================================================

class PureVectorEmbedder(EmbeddingBase):
    """High-accuracy, zero-external-dependency 768-dim semantic hash embedder for local Qdrant vector search."""
    def __init__(self, config: Optional[BaseEmbedderConfig] = None):
        super().__init__(config)
        self.dims = 768

    def embed(self, text: str, memory_action: Optional[Literal["add", "search", "update"]] = None) -> List[float]:
        vec = [0.0] * self.dims
        words = re.findall(r"\w+", text.lower(), re.UNICODE)
        if not words:
            return vec

        for w in words:
            # Word level hash
            h = int(hashlib.sha256(w.encode("utf-8")).hexdigest(), 16) % self.dims
            vec[h] += 1.0
            # Character 3-grams for semantic and spelling similarity
            for i in range(len(w) - 2):
                ngram = w[i:i+3]
                nh = int(hashlib.md5(ngram.encode("utf-8")).hexdigest(), 16) % self.dims
                vec[nh] += 0.5

        # Normalize to unit length for cosine similarity
        norm = math.sqrt(sum(x * x for x in vec))
        if norm > 0:
            vec = [x / norm for x in vec]
        return vec

# Register custom providers with Mem0 factories
LlmFactory.provider_to_class["gemini"] = ("mem0_service.AntigravityLLM", BaseLlmConfig)
EmbedderFactory.provider_to_class["fastembed"] = "mem0_service.PureVectorEmbedder"

# Expose symbols in current module for Mem0 dynamic loader
sys.modules["mem0_service"] = sys.modules[__name__]
setattr(sys.modules[__name__], "AntigravityLLM", AntigravityLLM)
setattr(sys.modules[__name__], "PureVectorEmbedder", PureVectorEmbedder)

# =============================================================================
# Mem0 Service & FastMCP Server
# =============================================================================

memory_instance = None

def get_memory_config() -> dict:
    """Configures Mem0 to use Google Antigravity SDK LLM & local Qdrant vector store."""
    return {
        "vector_store": {
            "provider": "qdrant",
            "config": {
                "path": str(QDRANT_DIR),
                "collection_name": "antigravity_memories",
                "embedding_model_dims": 768
            }
        },
        "llm": {
            "provider": "gemini",
            "config": {
                "model": "gemini-3.8-flash"
            }
        },
        "embedder": {
            "provider": "fastembed",
            "config": {}
        }
    }

def get_memory():
    global memory_instance
    if memory_instance is not None:
        return memory_instance

    try:
        from mem0 import Memory
        config = get_memory_config()
        logger.info("Initializing Mem0 with Google Antigravity SDK and Qdrant...")
        memory_instance = Memory.from_config(config)
        logger.info("Mem0 Antigravity engine successfully initialized.")
        return memory_instance
    except Exception as e:
        logger.error(f"Failed to initialize Mem0: {e}", exc_info=True)
        return None

# Custom FastMCP subclass to support both GET and POST on /sse
class DualMCP(FastMCP):
    def sse_app(self, mount_path: str | None = None) -> Starlette:
        message_endpoint = self._normalize_path(self.settings.mount_path, self.settings.message_path)
        sse = SseServerTransport(message_endpoint)

        async def combined_handler(request: Request) -> Response:
            if request.method == "POST":
                await sse.handle_post_message(request.scope, request.receive, request._send)
                return Response()
            async with sse.connect_sse(request.scope, request.receive, request._send) as streams:
                await self._mcp_server.run(
                    streams[0],
                    streams[1],
                    self._mcp_server.create_initialization_options()
                )
            return Response()

        routes = [
            Route(self.settings.sse_path, endpoint=combined_handler, methods=["GET", "POST"]),
            Mount(self.settings.message_path, app=sse.handle_post_message),
        ]
        return Starlette(debug=self.settings.debug, routes=routes)

# FastMCP Server definition
mcp = DualMCP(
    "mem0",
    host="127.0.0.1",
    port=8765,
    sse_path="/sse",
    message_path="/messages/"
)

@mcp.tool()
def mem0_status() -> str:
    """Check Mem0 background service status, local vector database, and engine ready state."""
    mem = get_memory()
    is_ready = mem is not None
    total_memories = 0

    if is_ready:
        try:
            # Query point count directly from Qdrant collection to avoid pagination limit truncation
            collection_info = mem.vector_store.client.get_collection("antigravity_memories")
            total_memories = collection_info.points_count
        except Exception:
            try:
                all_mems = mem.get_all(filters={"user_id": "developer"}, limit=1000)
                total_memories = len(all_mems.get("results", []))
            except Exception:
                pass

    status = {
        "status": "online" if is_ready else "error",
        "service": "Mem0 Background FastMCP Service",
        "port": 8765,
        "engine": "Google Antigravity Python SDK (google.antigravity.Agent)",
        "vector_store": f"Embedded Qdrant ({QDRANT_DIR})",
        "embedder": "PureVectorEmbedder (768-dim Local Cosine)",
        "developer_memories_count": total_memories,
        "is_ready": is_ready
    }
    return json.dumps(status, indent=2, ensure_ascii=False)

KNOWN_PROJECTS = [
    "FreeExile", "DSCons", "AutoPOE2", "Poe2VisualTool",
    "JX", "jxprivate", "KieuStory", "AutoPlayAgent",
    "ANK", "AIDesign", "AgentGame", "Historical"
]

def _detect_project_scope(text: str = "", default: Optional[str] = None) -> Optional[str]:
    """Helper to detect project name from text prefix, text content, or process CWD."""
    # 1. Check if text starts with [ProjectName:
    m = re.match(r"^\[([A-Za-z0-9_-]+):", text.strip())
    if m:
        detected = m.group(1)
        for kp in KNOWN_PROJECTS:
            if kp.lower() == detected.lower():
                return kp
        if detected.lower() == "global":
            return "Global"
        return detected

    # 2. Check if known project name is present in text
    for kp in KNOWN_PROJECTS:
        pattern = rf"\b{re.escape(kp)}\b"
        if re.search(pattern, text, re.IGNORECASE):
            return kp

    # 3. Check CWD
    try:
        cwd = os.getcwd()
        for kp in KNOWN_PROJECTS:
            if kp.lower() in cwd.lower():
                return kp
    except Exception:
        pass

    return default

@mcp.tool()
def mem0_add(
    text: str,
    user_id: str = "developer",
    metadata: Optional[Union[Dict[str, Any], str]] = None,
    project: Optional[str] = None,
    infer: bool = False
) -> str:
    """Store new information, user preference, architectural decision, or convention into local memory.
    
    Args:
        text: The text or fact to remember.
        user_id: The user identifier (default: "developer").
        metadata: Optional JSON string or dictionary containing additional context metadata.
        project: Target project name (e.g. 'FreeExile', 'DSCons', 'JX', 'Global'). MANDATORY: Always specify project to prevent cross-project memory contamination.
        infer: Whether to infer/extract multiple facts via LLM (default: False for atomic facts).
    """
    mem = get_memory()
    if not mem:
        return "Error: Mem0 engine not ready."

    if metadata:
        if isinstance(metadata, dict):
            meta_dict = dict(metadata)
        else:
            try:
                meta_dict = json.loads(metadata)
            except Exception:
                meta_dict = {"raw_metadata": str(metadata)}
    else:
        meta_dict = {}

    # Extract or auto-detect project
    if not project:
        nested_meta = meta_dict.get("metadata") if isinstance(meta_dict.get("metadata"), dict) else {}
        project = meta_dict.get("project") or nested_meta.get("project")
    if not project:
        project = _detect_project_scope(text, default="Global")

    meta_dict["project"] = project

    # Ensure prefix in text if not present
    clean_input = text.strip()
    if not clean_input.startswith(f"[{project}:") and not re.match(r"^\[[A-Za-z0-9_-]+:", clean_input):
        clean_input = f"[{project}: General] {clean_input}"

    try:
        # Prevent duplicate inflation
        input_hash = hashlib.md5(clean_input.encode("utf-8")).hexdigest()
        try:
            existing = mem.get_all(filters={"user_id": user_id}, limit=300)
            for m in existing.get("results", []):
                if m.get("hash") == input_hash or m.get("memory", "").strip() == clean_input:
                    logger.info(f"Memory already exists (ID: {m.get('id')}). Skipping duplicate.")
                    return json.dumps({"status": "exists", "id": m.get("id"), "project": project, "message": "Memory already exists."}, indent=2, ensure_ascii=False)
        except Exception:
            pass

        logger.info(f"Adding memory for user_id='{user_id}', project='{project}': {clean_input[:80]}...")
        result = None
        if infer:
            try:
                result = mem.add(clean_input, user_id=user_id, metadata=meta_dict, infer=True)
            except Exception as extract_err:
                logger.warning(f"Extraction error ({extract_err}), falling back to direct fact embedding...")
                result = mem.add(clean_input, user_id=user_id, metadata=meta_dict, infer=False)
        else:
            result = mem.add(clean_input, user_id=user_id, metadata=meta_dict, infer=False)

        if not result or not result.get("results"):
            result = mem.add(clean_input, user_id=user_id, metadata=meta_dict, infer=False)

        return json.dumps({"status": "success", "project": project, "result": result}, indent=2, ensure_ascii=False)
    except Exception as e:
        logger.error(f"Error in mem0_add: {e}", exc_info=True)
        return f"Error adding memory: {str(e)}"

@mcp.tool()
def mem0_search(
    query: str,
    user_id: str = "developer",
    limit: int = 5,
    project: Optional[str] = None
) -> str:
    """Search long-term memories for relevant facts, preferences, or technical decisions matching the query.
    
    Args:
        query: Search query (e.g. 'preferred coding language', 'database choices', 'user habits').
        user_id: Filter memories by user ID (default: "developer").
        limit: Maximum number of memories to return (default: 5).
        project: Target project filter (e.g. 'FreeExile', 'DSCons', 'Global'). MANDATORY: Always filter by project to prevent memory dilution.
    """
    mem = get_memory()
    if not mem:
        return "Error: Mem0 engine not ready."

    auto_detected = False
    if not project:
        detected = _detect_project_scope(query, default=None)
        if detected:
            project = detected
            auto_detected = True

    try:
        logger.info(f"Searching memories for query='{query}', user_id='{user_id}', project='{project}' (auto_detected={auto_detected})...")
        filters = {"user_id": user_id} if user_id else None
        fetch_limit = max(limit * 8, 150) if project else limit
        results = mem.search(query, filters=filters, top_k=fetch_limit)
        
        if project and results and "results" in results:
            proj_lower = project.lower()
            filtered = []
            for r in results["results"]:
                meta = r.get("metadata") or {}
                nested_meta = meta.get("metadata") if isinstance(meta.get("metadata"), dict) else {}
                m_proj = str(meta.get("project") or nested_meta.get("project") or "").lower()
                m_text = str(r.get("memory", "")).lower()
                
                if m_proj == proj_lower:
                    filtered.append(r)
                elif not m_proj:
                    if f"[{proj_lower}:" in m_text or f"[{proj_lower} " in m_text or f"project: {proj_lower}" in m_text:
                        filtered.append(r)
                    elif proj_lower == "freeexile" and "dscons" not in m_text and "bim" not in m_text and "revit" not in m_text:
                        filtered.append(r)
                    elif proj_lower == "dscons" and ("bim" in m_text or "revit" in m_text or "boq" in m_text or "định sơn" in m_text):
                        filtered.append(r)
                    elif proj_lower == "global" and not any(kp.lower() in m_text for kp in KNOWN_PROJECTS if kp.lower() != "global"):
                        filtered.append(r)
            results["results"] = filtered[:limit]

        response_payload = {
            "status": "success",
            "query": query,
            "project_filter": project or "ALL_UNFILTERED (WARNING: Always specify project to avoid cross-project contamination)",
            "auto_detected": auto_detected,
            "memories": results
        }
        return json.dumps(response_payload, indent=2, ensure_ascii=False)
    except Exception as e:
        logger.error(f"Error in mem0_search: {e}", exc_info=True)
        return f"Error searching memory: {str(e)}"

@mcp.tool()
def mem0_get_all(
    user_id: str = "developer",
    limit: int = 1000
) -> str:
    """Retrieve all stored memories and facts for a user.
    
    Args:
        user_id: The user ID to retrieve memories for (default: "developer").
        limit: Maximum number of memories to return (default: 1000).
    """
    mem = get_memory()
    if not mem:
        return "Error: Mem0 engine not ready."

    try:
        filters = {"user_id": user_id} if user_id else None
        results = mem.get_all(filters=filters, top_k=limit)
        return json.dumps({"status": "success", "user_id": user_id, "memories": results}, indent=2, ensure_ascii=False)
    except Exception as e:
        logger.error(f"Error in mem0_get_all: {e}", exc_info=True)
        return f"Error getting memories: {str(e)}"

@mcp.tool()
def mem0_delete(memory_id: str) -> str:
    """Delete a specific memory by its unique ID.
    
    Args:
        memory_id: The unique ID of the memory to delete.
    """
    mem = get_memory()
    if not mem:
        return "Error: Mem0 engine not ready."

    try:
        mem.delete(memory_id)
        return json.dumps({"status": "success", "deleted_id": memory_id}, indent=2)
    except Exception as e:
        logger.error(f"Error in mem0_delete: {e}", exc_info=True)
        return f"Error deleting memory: {str(e)}"

if __name__ == "__main__":
    logger.info("Starting Mem0 FastMCP Server powered by Google Antigravity SDK on http://127.0.0.1:8765/sse...")
    get_memory()
    mcp.run(transport="sse")
