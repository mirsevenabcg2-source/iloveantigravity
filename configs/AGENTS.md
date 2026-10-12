# Antigravity Global Engineering Directives

## 1. Autonomous Skill Synthesis & Bounded Memory (Hermes Protocol) - TOP PRIORITY
- **Bounded Snapshot (<2,500 chars)**: Giữ System Prompt tinh gọn, triệt tiêu context dilution. Chi tiết đẩy vào `mem0_search(project=...)` hoặc SQLite FTS5.
- **Proactive Context Compaction @ 50%**: Khi context đạt 50%, tóm tắt/cắt tỉa các khối `tool_result` cồng kềnh (stdout build, log dài, diff lớn) thành tóm tắt 2-3 dòng.
- **Skill Synthesis (`skill-synthesizer`)**: Tự động phản tư sau khi giải quyết task phức tạp, đúc kết thành `~/.gemini/config/skills/<name>/SKILL.md` (chuẩn agentskills.io, progressive disclosure).
- **Local Fast Recall**: Dùng SQLite FTS5 hoặc local grep tra cứu log/mã lỗi cũ (<2ms, 0 API token, 0% VRAM).
- **Sandbox Awareness**: Chuyển hướng lệnh nguy hiểm (untrusted script, crawl nặng) vào container Docker/Sandbox.

## 2. AST Code Intelligence (`ast-grep` & `ast-grep-mcp`)
- Cấm regex thô sơ để refactor trên codebase lớn; bắt buộc dùng `ast-grep` + `lsp-mcp`.
- **2-Phase Safety Gate**: 1. Preview (`apply: false`) -> 2. Mutation (`apply: true`) -> 3. Kiểm chứng bằng `lsp_diagnostics` hoặc test runner.

## 3. Code Review Delegation Mode (Alibaba Open Code Review `ocr`)
- **Quy trình 5 bước**:
  1. Preview: `ocr delegate preview --format json [--from <ref> --to <ref>] [--commit <hash>]`
  2. Ruleset: `ocr delegate rule --format json <files...>`
  3. Git Diff: `git diff <merge_base>..<to> -- <path>`
  4. Đối soát: Gemini đối chiếu từng rule group, định vị chính xác dòng, phân loại (`critical`, `high`, `medium`, `low`).
  5. Auto-fix: Sửa trực tiếp lỗi High/Critical nếu được yêu cầu.

## 4. Bộ Nhớ Mem0 & Cô Lập Project (Zero Dilution)
- Dịch vụ cục bộ port `8765`, SSE `http://127.0.0.1:8765/sse`, Qdrant embedded `~/.gemini/antigravity/mem0_data/qdrant`.
- **Quy tắc bắt buộc**:
  - Luôn chỉ định `project`: `mem0_search(query=..., project="<ProjectName>")` (hoặc `project="Global"`).
  - Nội dung lưu trữ 2 lớp: `mem0_add(text="[<ProjectName>: <Category>] ...", project="<ProjectName>")`.

## 5. Git Worktree Delegation (Subagent Isolation)
- Refactoring lớn, nâng cấp thư viện rủi ro cao: Bắt buộc `invoke_subagent` với `Workspace: 'branch'`.
- Host Agent chỉ merge khi subagent test pass 100%.

## 6. TDD Closed-Loop & Autonomous Verification
- **Chu trình khép kín**: 1. Viết test trước -> 2. Xác nhận FAIL (negative baseline) -> 3. Viết code tối thiểu -> 4. Self-healing loop (tối đa 5 lần theo traceback, không hỏi user giữa chừng) -> 5. Atomic git commit khi PASS.

## 7. Native Impact Guard & Blast Radius
- PreToolUse hook tự động chặn sửa file cốt lõi/file nhiều liên đới (>= 6 dependents).
- **Mở khóa**: Chạy `python tools/analysis/blast_radius.py --target <path> --ack` để mở khóa trong 30 phút.
- Windows safe: `ensure_ascii=True`, bỏ qua `node_modules`, `.git`, `dist`, `__pycache__`.

## 8. Chuẩn Hóa Code, Modular SRP & Anti-AI Slop
- **Modular & Single Responsibility Principle (SRP)**: Phân tách module rõ ràng, mỗi hàm/file chỉ đảm nhận một trách nhiệm duy nhất.
- **Ưu tiên Common / Utils / Helper**: Luôn trích xuất và tái sử dụng logic dùng chung qua các thư mục/file `common`, `utils`, `helper`; triệt tiêu duplicate code.
- **Chọn Đúng Ngôn Ngữ (Right Tool for the Job)**: Tuyệt đối CẤM mặc định dùng Python cho mọi thứ; chọn Go/Rust/C++ cho hệ thống, standalone binary, low-latency OS syscalls; Python cho ML/data/AST; TypeScript cho web/extensions.
- **Trần cứng độ dài file (500–800 dòng)**: Mọi file mã nguồn hoặc tài liệu tuyệt đối KHÔNG vượt quá 500–800 dòng; chủ động phân tách nhỏ module/component khi tiệm cận ngưỡng.
- **Tài liệu & Giải thích ngắn gọn**: Trình bày súc tích, đi thẳng vào bản chất kỹ thuật, không rườm rà hay giải thích lan man.
- **Zero Placeholder & Root-Cause First**: Cấm `// TODO`, `/* giữ nguyên */`; mã sinh ra phải chạy được ngay; xử lý triệt để nguyên nhân gốc rễ (không `try/catch` rỗng hay ép kiểu `any`).
