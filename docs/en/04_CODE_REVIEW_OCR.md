# 🔍 Alibaba Open Code Review (OCR): Delegation Mode Architecture

## 1. Why Delegation Mode?
Traditional AI code review prompts suffer from two fatal flaws:
1. **Line Hallucinations**: LLMs frequently miscount line numbers when reading long patches, causing inaccurate review comments and failed automated patches.
2. **Token & Cost Waste**: Sending entire files to remote cloud endpoints wastes millions of tokens on vendor files, lockfiles, and boilerplate code.

The **Alibaba Open Code Review (`ocr`) Delegation Mode** solves this by establishing a strict hybrid architecture:
- Deterministic engineering handles file filtering, merge-base resolution, and rule group generation.
- The Gemini host agent provides deep reasoning, line verification, severity classification, and automated repair.

---

## 2. The 5-Step Deterministic Workflow

```
[ Git Diff / Workspace ]
           │
           ▼
Step 1: ocr delegate preview --format json
           │ (Filters lockfiles & junk, identifies affected files & commit base)
           ▼
Step 2: ocr delegate rule --format json <files...>
           │ (Fetches Alibaba enterprise checklist: NPE, Thread Safety, SQLi, Resource Leaks)
           ▼
Step 3: git diff <merge_base>..<to> -- <path>
           │ (Extracts real patch without synthetic modification)
           ▼
Step 4: Gemini Deep Reasoning & Verification
           │ (Audits each rule group against unified diff, calculates exact lines & severity)
           ▼
Step 5: Standardized Report & Autonomous Auto-Fix
           │ (Patches High/Critical defects immediately if requested)
```

---

## 3. Standardized Finding Schema
Every finding reported by the agent must conform to this schema:
- **File & Range**: `path/to/file.ext#L<start>-L<end>`
- **Severity**:
  - `critical`: Security vulnerabilities (SQLi, RCE, auth bypass), severe memory leaks.
  - `high`: Logical bugs, potential null pointer exceptions (NPE), concurrency race conditions, data loss.
  - `medium`: Performance regressions, missing index queries, anti-patterns.
  - `low`: Style inconsistencies, minor maintainability suggestions.
- **Root Cause & Remediated Code**: Detailed technical explanation followed by a concrete, runnable diff.
