# Gemini System & Protocol Directives

## 1. Autonomous Skill Synthesis & Bounded Memory (Hermes Protocol) - TOP PRIORITY
- **Bounded Memory (<2,500 chars)**: Khống chế snapshot khởi tạo, chống token bloat. Tra cứu sâu qua `mem0_search(project=...)` hoặc SQLite FTS5.
- **Proactive Context Compaction @ 50%**: Khi đạt 50% context, chủ động nén/cắt tỉa tool outputs lớn thành tóm tắt 2-3 dòng, bảo toàn flow trao đổi.
- **Autonomous Skill Synthesis**: Tự động phản tư và sinh file `SKILL.md` qua `skill-synthesizer` sau khi xử lý xong bài toán kỹ thuật phức tạp.
- **Execution Sandbox Awareness**: Ưu tiên container/sandbox cô lập cho các lệnh rủi ro cao.

## 2. AST Code Intelligence Protocol (ast-grep & ast-grep-mcp)
- Cấm regex thô sơ để refactor trên codebase lớn. Sử dụng `ast-grep` kết hợp `lsp-mcp`.
- 2-Phase Safety Gate: Preview (`apply: false`) -> Mutation (`apply: true`) -> Compiler/Diagnostics check.

## 3. Code Review Protocol (Alibaba Open Code Review Delegation)
- Thực thi qua CLI `ocr delegate`: Preview (`preview --format json`) -> Ruleset (`rule --format json`) -> Git Diff -> Gemini đối soát theo rule checklist -> Auto-fix. Zero external API cost.

## 4. Mem0 Local Memory & Project Isolation (Zero Dilution)
- Dịch vụ nền port 8765, Qdrant embedded.
- Bắt buộc tham số `project`: `project="<ProjectName>"` hoặc `project="Global"`.
- Tiền tố lưu trữ chuẩn: `mem0_add(text="[<ProjectName>: <Category>] ...", project="<ProjectName>")`.

## 5. Autonomous Git Worktree Isolation
- Refactor lớn / rủi ro: Bắt buộc dùng `invoke_subagent(Workspace='branch')` để chạy trong Git Worktree cô lập.

## 6. Autonomous TDD Closed-Loop
- Test-First -> Verify FAIL -> Code tối thiểu -> Tự đọc traceback và sửa lặp lại (tối đa 5 lần) -> Atomic git commit.

## 7. Native Lifecycle Hook & Blast Radius
- Hard gate chặn sửa file liên đới cao. Mở khóa qua: `python tools/analysis/blast_radius.py --target <path> --ack` (30 phút).
- Windows safe: `ensure_ascii=True`, bỏ qua `node_modules`, `.git`, `dist`, `__pycache__`.

## 8. Chuẩn Hóa Code, Modular SRP & Anti-AI Slop
- **Modular & SRP**: Phân tách Single Responsibility Principle; ưu tiên dùng `common`, `utils`, `helper`.
- **Chọn đúng ngôn ngữ (Right Tool for Job)**: Cấm mặc định dùng Python; Go/Rust/C++ cho hệ thống/standalone binary/low-latency; Python cho AI/data/AST; TypeScript cho web/runtime.
- **Trần file 500-800 dòng**: Tuyệt đối không để file code/tài liệu vượt quá 500-800 dòng; chủ động tách nhỏ module.
- **Code & Doc tinh gọn**: Giải thích ngắn gọn, đi thẳng vào bản chất kỹ thuật; full runnable code, cấm placeholder `// TODO`.
