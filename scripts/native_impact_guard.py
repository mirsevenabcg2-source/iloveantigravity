#!/usr/bin/env python3
"""
Antigravity Global Native Impact Guard.
PreToolUse lifecycle hook intercepting file modifications across all workspaces.
Dispatches to local blast radius calculators or falls back to generic AST scanning.
"""
from __future__ import annotations

import ast
import json
import os
import subprocess
import sys
from typing import Any, Dict, List, Optional

if hasattr(sys.stdout, "reconfigure"):
    sys.stdout.reconfigure(encoding="utf-8")
if hasattr(sys.stderr, "reconfigure"):
    sys.stderr.reconfigure(encoding="utf-8")


def _is_safe_file(target_file: str) -> bool:
    norm = target_file.replace("\\", "/").lower()
    # Documentation and wiki specifications are domain contracts; not safe files
    if (norm.startswith("docs/") or "/docs/" in norm or norm.startswith("wiki/") or "/wiki/" in norm) and norm.endswith(".md"):
        return False
    # Skip non-doc text, configs, logs, artifacts, scratch, and test files
    safe_extensions = (".md", ".txt", ".json", ".lock", ".log", ".csv", ".svg", ".png")
    if any(norm.endswith(ext) for ext in safe_extensions):
        return True
    safe_dirs = [
        "/.gemini/",
        "/.git/",
        "/tests/",
        "/test/",
        "/scratch/",
        "/artifacts/",
        "/node_modules/",
        "/__pycache__/",
    ]
    return any(sd in norm for sd in safe_dirs)


def _find_workspace_root(target_file: str, workspace_paths: List[str]) -> Optional[str]:
    if workspace_paths and len(workspace_paths) > 0 and os.path.exists(workspace_paths[0]):
        return os.path.abspath(workspace_paths[0])
    curr = os.path.abspath(os.path.dirname(target_file))
    while curr:
        if os.path.exists(os.path.join(curr, ".git")):
            return curr
        parent = os.path.dirname(curr)
        if parent == curr:
            break
        curr = parent
    return None


def _find_local_calculator(workspace_root: str) -> Optional[str]:
    candidates = [
        os.path.join(workspace_root, "tools", "analysis", "blast_radius.py"),
        os.path.join(workspace_root, ".agents", "blast_radius.py"),
        os.path.join(workspace_root, "tools", "lint", "blast_radius.py"),
    ]
    for c in candidates:
        if os.path.exists(c):
            return c
    return None


def _run_generic_ast_scan(workspace_root: str, target_file: str) -> int:
    """Fallback scanner for repos without local domain calculator."""
    if not target_file.endswith(".py"):
        return 0
    stem = os.path.splitext(os.path.basename(target_file))[0]
    dependent_count = 0

    for root, dirs, files in os.walk(workspace_root):
        dirs[:] = [d for d in dirs if d not in [".git", ".gemini", "node_modules", "venv", ".venv", "__pycache__", "dist", "build", "downloads", "chrome_profile"]]
        for f in files:

            if not f.endswith(".py"):
                continue
            full_p = os.path.join(root, f)
            if os.path.abspath(full_p) == os.path.abspath(target_file):
                continue
            try:
                with open(full_p, "r", encoding="utf-8", errors="ignore") as src:
                    content = src.read()
                if stem in content:
                    tree = ast.parse(content, filename=full_p)
                    for node in ast.walk(tree):
                        if isinstance(node, (ast.Import, ast.ImportFrom)):
                            dependent_count += 1
                            break
            except Exception:
                continue
    return dependent_count


def process_hook() -> Dict[str, Any]:
    try:
        raw_input = sys.stdin.read()
        if not raw_input.strip():
            return {"decision": "allow", "reason": "Empty payload received."}
        payload = json.loads(raw_input)
    except Exception as e:
        return {"decision": "allow", "reason": f"Payload parse bypass: {e}"}

    tool_call = payload.get("toolCall", {})
    tool_name = tool_call.get("name", "")
    args = tool_call.get("args", {})

    target_file = (
        args.get("TargetFile")
        or args.get("target_file")
        or args.get("targetFile")
        or args.get("AbsolutePath")
        or args.get("path")
    )

    if not target_file:
        return {"decision": "allow", "reason": "No target file identified in tool arguments."}

    if _is_safe_file(str(target_file)):
        return {"decision": "allow", "reason": "Safe / non-production file permitted."}

    workspace_paths = payload.get("workspacePaths", [])
    workspace_root = _find_workspace_root(str(target_file), workspace_paths)

    if not workspace_root:
        return {"decision": "allow", "reason": "Workspace root undetermined; permitted."}

    local_calculator = _find_local_calculator(workspace_root)

    # 1. Local Domain Calculator
    if local_calculator:
        try:
            cmd = [
                sys.executable,
                local_calculator,
                "--target",
                str(target_file),
                "--format",
                "json",
            ]
            res = subprocess.run(
                cmd,
                capture_output=True,
                text=True,
                cwd=workspace_root,
                encoding="utf-8",
                timeout=12,
            )
            if res.returncode == 0:
                data = json.loads(res.stdout)
                if data.get("blocked", False):
                    return {
                        "decision": "deny",
                        "reason": data.get("reason", "Blocked by local project Blast Radius Calculator."),
                    }
                return {
                    "decision": "allow",
                    "reason": data.get("reason", "Cleared by local project Blast Radius Calculator."),
                }
        except Exception as e:
            # On timeout or error, do not brick agent; log and permit
            return {"decision": "allow", "reason": f"Local calculator exception bypass: {e}"}

    # 2. Generic AST Fallback
    dependents = _run_generic_ast_scan(workspace_root, str(target_file))
    if dependents >= 6:
        return {
            "decision": "deny",
            "reason": (
                f"[GLOBAL NATIVE IMPACT GUARD] File '{os.path.basename(str(target_file))}' có {dependents} "
                f"module khác phụ thuộc. Đề xuất Agent đọc kỹ cấu trúc trước khi sửa đổi!"
            ),
        }

    return {"decision": "allow", "reason": f"Generic scan passed ({dependents} dependents)."}


def main() -> int:
    try:
        decision = process_hook()
    except Exception as e:
        decision = {"decision": "allow", "reason": f"Global hook unexpected error: {e}"}

    print(json.dumps(decision, ensure_ascii=True))
    return 0


if __name__ == "__main__":
    sys.exit(main())
