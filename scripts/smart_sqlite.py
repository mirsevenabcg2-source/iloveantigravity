import sys
import os
import asyncio
from pathlib import Path
from mcp_server_sqlite import server

def find_project_root(start_dir: str | Path) -> Path | None:
    current = Path(start_dir).resolve()
    for parent in [current, *current.parents]:
        for marker in ['.git', '.agents', '.agent', 'package.json', 'pyproject.toml', 'Cargo.toml', 'go.mod']:
            if (parent / marker).exists():
                return parent
    return None

def resolve_db_path() -> str:
    # 1. Command-line argument: --db-path <path> or --db-path=<path>
    for i, arg in enumerate(sys.argv):
        if arg.startswith("--db-path="):
            return arg.split("=", 1)[1]
        if arg == "--db-path" and i + 1 < len(sys.argv):
            return sys.argv[i + 1]

    # 2. Environment variable
    env_path = os.environ.get("SQLITE_DB_PATH")
    if env_path:
        return os.path.abspath(env_path)

    # 3. Project root detection
    cwd = Path.cwd()
    project_root = find_project_root(cwd) or cwd

    # Check common db names in project root
    common_names = ["dev.db", "app.db", "db.sqlite3", "database.sqlite", "data.db", "sqlite.db"]
    for name in common_names:
        candidate = project_root / name
        if candidate.exists():
            return str(candidate)

    # Check for any .sqlite or .db file in project root
    db_files = list(project_root.glob("*.sqlite")) + list(project_root.glob("*.db")) + list(project_root.glob("*.sqlite3"))
    if db_files:
        return str(db_files[0])

    # Fallback to local project-specific db file
    return str(project_root / ".sqlite_mcp.db")

if __name__ == "__main__":
    db_path = resolve_db_path()
    asyncio.run(server.main(db_path))
