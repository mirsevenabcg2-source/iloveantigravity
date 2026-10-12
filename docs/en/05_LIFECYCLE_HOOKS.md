# 🛡️ Native Lifecycle Hooks & Blast Radius Protection

## 1. Overview
Google Antigravity provides lifecycle hooks (`PreToolUse`, `PostToolUse`) that allow intercepting tool calls before and after execution. **iLoveAntigravity** harnesses these hooks to enforce system safety, prevent architectural destruction, and control memory leaks.

---

## 2. PreToolUse: Native Impact Guard (`scripts/native_impact_guard.py`)
Intercepts file modification tools:
- `write_to_file`
- `replace_file_content`

### How It Works:
1. **Target File Identification**: Detects the file path intended for modification.
2. **Safe Filter**: Markdown documentation, text files, test fixtures, and scratch scripts are bypassed automatically.
3. **AST Blast Radius Scan**:
   - Walks through workspace source files.
   - Parses Abstract Syntax Trees (AST) to discover modules importing or dependent on the target file.
4. **Safety Gate Threshold**:
   - If the target file affects **$\ge 6$ dependent modules**, the tool call is denied immediately.
   - **Unlock Protocol**: To bypass the safety gate for legitimate large-scale refactors, execute:
     ```bash
     python tools/analysis/blast_radius.py --target <path> --ack
     ```
     This grants a 30-minute modification lease for that specific target.

---

## 3. PostToolUse: Aider Native Guard (`go_hooks/aider_guard`)
A compiled native Go hook executed after file modification tools complete:

### Functions:
1. **Instant Syntax Gate**:
   - Validates JSON files (blocks invisible UTF-8 Byte Order Marks `\xEF\xBB\xBF` which break Node/Webpack build chains).
   - Validates Python files using `python -m py_compile` to catch syntax errors before compilation.
2. **Language Server Memory Trimmer**:
   - In background goroutines, queries Windows process snapshots (`CreateToolhelp32Snapshot`).
   - If any `language_server.exe` process consumes more than 1,500 MB of Working Set memory, triggers `EmptyWorkingSet` via Win32 API to reclaim physical RAM.
