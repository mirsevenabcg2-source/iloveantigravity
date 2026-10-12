# 🏛️ The Hermes Protocol: Context Optimization & Bounded Memory

## 1. The Context Dilution Problem
Traditional AI coding agents suffer from rapid degradation as conversations lengthen:
- System prompts balloon with rules, examples, and tool definitions.
- Massive tool outputs (build logs, verbose compiler output, multi-thousand line diffs) eat up the context window.
- The model experiences "attention drift" or "context dilution," hallucinating previously verified details and forgetting fundamental architectural rules.

## 2. The Hermes Solution
The **Hermes Protocol** introduces three non-negotiable architectural gates:

### Gate A: Bounded Memory Snapshot (<2,500 Chars)
- The global System Prompt (`AGENTS.md`) is kept strictly under 2,500 characters.
- Long-tail specifications, schemas, project habits, and domain contracts are offloaded to **Mem0** vector storage and local SQLite full-text search (FTS5).
- Deep knowledge is retrieved *just-in-time* on demand rather than preemptively cluttering the context window.

### Gate B: Proactive Context Compaction @ 50% Window
- When total conversation tokens reach 50% of the active context ceiling:
  - Giant tool responses (`tool_result`) such as build logs, package installation lists, and compiler tracebacks are pruned into concise 2-3 line structured summaries.
  - Core decisions, errors discovered, and actionable items are retained; redundant ASCII logs are discarded.
- This maintains maximum attention resolution on the primary problem.

### Gate C: Autonomous Skill Synthesis (`skill-synthesizer`)
- When the agent successfully completes a non-trivial engineering workflow (e.g., configuring an unfamiliar build tool, resolving complex OS-level dependencies, reverse engineering a binary protocol), it self-reflects.
- The workflow is synthesized into a standardized `SKILL.md` conforming to the `agentskills.io` specification.
- Future sessions discover and load this skill instantly without re-learning from scratch.

### Gate D: Anti-AI Slop & Structural Hygiene
- **500–800 Line Hard Ceiling**: No file should exceed 800 lines. When approaching this limit, logic must be split into modular Single-Responsibility Principle (SRP) sub-components.
- **Zero Placeholders**: Banned patterns include `// TODO`, `/* keep existing code */`, or empty catch blocks. Generated code must be deployable immediately.
- **Language Alignment**: Go/Rust/C++ for low-latency systems and tooling; Python for ML, AST, and data pipelines; TypeScript for web interfaces.
