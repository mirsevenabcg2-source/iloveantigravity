# 🚀 iLoveAntigravity

> **Bộ Công Cụ Tối Ưu Hóa & Nâng Cấp Toàn Diện Dành Cho Google Antigravity & AI Coding Agent.**  
> *Cài đặt 1-click ăn ngay, di chuyển giữa các máy tính trong 60 giây, triệt tiêu loãng ngữ cảnh (Zero Dilution), máy chủ MCP viết bằng Go siêu tốc, bộ nhớ vector Mem0 cục bộ và quy chuẩn kỹ thuật đỉnh cao.*

---

### 🌐 Tài Liệu Đa Ngôn Ngữ
[ English ](README.md) • [ Tiếng Việt ](README.vi.md) • [ 简体中文 ](README.zh.md) • [ 日本語 ](README.ja.md)

---

## ⚡ iLoveAntigravity Là Gì?

**iLoveAntigravity** là kho lưu trữ tổng hợp toàn bộ tinh hoa kiến trúc, kinh nghiệm thực chiến, mã nguồn nâng cấp và quy chuẩn kỹ thuật được tôi đúc kết qua hàng ngàn giờ lập trình cùng [Google Antigravity](https://github.com/google/antigravity).

Mục tiêu cốt lõi của repository này là biến Agent từ một mô hình hội thoại đơn thuần thành một **Kỹ Sư Phần Mềm Cấp Cao (Senior Autonomous Engineer)**:
- **Di Chuyển Máy Tính Liền Mạch (Turnkey 1-Click Migration)**: Khi bạn đổi sang máy tính mới hoặc cài lại Win/Linux/macOS, chỉ cần chạy 1 dòng lệnh duy nhất để khôi phục toàn bộ môi trường tối ưu trong chưa đầy 1 phút.
- **Giao Thức Hermes (Hermes Protocol)**: Giới hạn System Prompt khởi tạo (<2,500 ký tự), chủ động nén ngữ cảnh (Proactive Compaction) khi đạt 50% cửa sổ token, và truy xuất cực nhanh (<2ms) qua SQLite FTS5/ripgrep để triệt tiêu tình trạng trôi dạt ngữ cảnh.
- **Hệ Thống MCP Biên Dịch Bằng Go (Native Go MCP Servers)**: Thay thế hoàn toàn các server Node.js/Python chậm chạp bằng binary Go biên dịch mã máy cho Git, Mem0 Bridge, SQLite và Memory Graph.
- **Bộ Nhớ Cá Nhân Hóa Mem0 Cục Bộ (Zero-Dilution Project Isolation)**: Vận hành ngầm qua cổng `8765`, sử dụng Qdrant embedded và trực tiếp gọi **Google Antigravity Python SDK (`google.antigravity.Agent`)** để suy luận. Hoàn toàn không tốn VRAM card đồ họa, cô lập triệt để dữ liệu giữa các project khác nhau.
- **Quy Trình Code Review Độc Quyền (Alibaba Open Code Review Delegation)**: Tận dụng cơ chế ủy thác 5 bước của Alibaba OCR để lọc file và trích xuất checklist, kết hợp suy luận dòng lệnh của Gemini để loại bỏ hoàn toàn lỗi lệch dòng và tiết kiệm chi phí token.
- **Hệ Thống Native Lifecycle Hooks & Impact Guard**: Chặn đứng sửa code gây lỗi lan truyền bằng AST Blast Radius Analyzer và tự động giải phóng RAM các Language Server bị phình to.
- **Bộ Kỹ Năng Đỉnh Cao (Elite Skills & Plugins)**: Tự động tổng hợp skill (`skill-synthesizer`), kiểm thử trình duyệt tàng hình anti-detect (`invisible-playwright`), cô lập nhánh Git Worktree và chu trình TDD khép kín.

---

## 🚀 Hướng Dẫn Cài Đặt 1-Click (Chuyển Máy Tính Sử Dụng Luôn)

Bất kỳ khi nào bạn chuyển sang máy tính mới, chỉ cần mở terminal và thực hiện:

### Trên Windows (PowerShell):
```powershell
# 1. Clone repository về máy
git clone https://github.com/mirsevenabcg2-source/iloveantigravity.git
cd iloveantigravity

# 2. Chạy kịch bản cài đặt tự động (mở PowerShell bình thường hoặc Run as Administrator)
.\install.ps1
```

### Trên Linux / macOS (Bash):
```bash
# 1. Clone repository về máy
git clone https://github.com/mirsevenabcg2-source/iloveantigravity.git
cd iloveantigravity

# 2. Cấp quyền thực thi và chạy kịch bản cài đặt
chmod +x install.sh
./install.sh
```

### Script Cài Đặt Tự Động Sẽ Làm Gì?
1. **Kiểm tra công cụ cần thiết**: Tự động phát hiện `Python 3.10+`, `Go 1.22+`, `Node.js`, `Git`.
2. **Khởi tạo thư mục chuẩn**: Tạo sẵn cấu trúc `~/.gemini/config` và `~/.gemini/antigravity`.
3. **Biên dịch mã nguồn Go**: Trực tiếp dịch 5 server MCP & hook bằng Go (`mcp_git`, `mcp_mem0_bridge`, `mcp_memory`, `mcp_sqlite`, `aider_guard`) thành file thực thi mã máy siêu tốc đặt vào `~/.gemini/config/bin`.
4. **Đồng bộ kịch bản & quy chuẩn**: Copy toàn bộ scripts, skills tùy biến, 11 plugins, cùng `AGENTS.md` và `GEMINI.md`.
5. **Sinh cấu hình chuẩn hóa**: Tự động render `mcp_config.json` và `hooks.json` theo đúng đường dẫn người dùng trên máy hiện tại.
6. **Cài đặt thư viện Python**: Tự động cài `mem0ai`, `qdrant-client`, `sentence-transformers`, `mcp`, `starlette`.
7. **Đăng ký dịch vụ Mem0 tự khởi động**: Tự động đưa Mem0 Daemon vào Windows Registry `Run` để luôn chạy nền sẵn sàng ngay khi máy khởi động.

---

## 🏛️ 7 Trụ Cột Kiến Trúc Kỹ Thuật

```
                                ┌────────────────────────┐
                                │   Google Antigravity   │
                                │    Master Host Agent   │
                                └───────────┬────────────┘
                                            │
               ┌────────────────────────────┼────────────────────────────┐
               ▼                            ▼                            ▼
    ┌──────────────────────┐     ┌──────────────────────┐     ┌──────────────────────┐
    │  Giao Thức Hermes    │     │   Native Go MCPs     │     │  Dịch Vụ Nền Mem0    │
    │  - Bounded Snapshot  │     │  - mcp_git           │     │  - Port 8765 SSE     │
    │  - Nén ngữ cảnh 50%  │     │  - mcp_mem0_bridge   │     │  - Qdrant Embedded   │
    │  - Skill Synthesis   │     │  - mcp_memory        │     │  - AGY Python SDK    │
    │  - Chống AI Slop     │     │  - mcp_sqlite        │     │  - Zero Dilution     │
    └──────────────────────┘     └──────────────────────┘     └──────────────────────┘
               │                            │                            │
               ▼                            ▼                            ▼
    ┌──────────────────────┐     ┌──────────────────────┐     ┌──────────────────────┐
    │  Lifecycle Hooks     │     │  Hybrid Code Review  │     │ Bộ Kỹ Năng Đỉnh Cao  │
    │  - Impact Guard      │     │  - Alibaba OCR       │     │  - Invisible Browser │
    │  - AST Blast Radius  │     │  - Delegation Mode   │     │  - Skill Synthesis   │
    │  - Aider Mem Trim    │     │  - 5 Bước Chuẩn Hóa  │     │  - Worktree & TDD    │
    └──────────────────────┘     └──────────────────────┘     └──────────────────────┘
```

### 1. Giao Thức Hermes (Hermes Protocol) & Nén Ngữ Cảnh
- **Bounded Snapshot (<2,500 ký tự)**: Giữ System Prompt luôn ngắn gọn, súc tích. Dữ liệu ngữ cảnh sâu được tra cứu theo nhu cầu qua `mem0_search` hoặc SQLite FTS5.
- **Chủ Động Nén Ngữ Cảnh @ 50% (Proactive Context Compaction)**: Khi cửa sổ token đạt mức 50%, Agent tự động tóm tắt và thu gọn các kết quả tool khổng lồ (log build, diff lớn) thành 2-3 dòng cốt lõi, bảo toàn tối đa độ nhạy bén của não bộ AI.
- **Tự Động Đúc Kết Kỹ Năng (`skill-synthesizer`)**: Sau khi xử lý xong các bài toán kỹ thuật phức tạp, Agent tự động phản tư và sinh ra file `SKILL.md` theo chuẩn `agentskills.io` để tái sử dụng vĩnh viễn.
- **Local Fast Recall**: Dùng SQLite FTS5 hoặc ripgrep cục bộ để tra cứu log/mã lỗi cũ (<2ms, 0 API token, 0% VRAM).

### 2. Dịch Vụ Mem0 Cục Bộ & Quy Tắc Cô Lập Project (Anti-Dilution)
- **Kiến Trúc Tối Ưu**: Dịch vụ nền chạy trên `http://127.0.0.1:8765/sse` kết hợp Qdrant embedded và bộ nhúng cosine 768 chiều.
- **Động Cơ Suy Luận**: Sử dụng trực tiếp **Google Antigravity Python SDK (`google.antigravity.Agent`)**, hoàn toàn không phụ thuộc các local LLM nặng nề (LM Studio, Ollama).
- **Quy Tắc Cô Lập Project Tuyệt Đối (Zero Dilution)**:
  - **Zero-Unscoped Rule**: Tuyệt đối CẤM gọi `mem0_search` hoặc `mem0_add` mà không chỉ định rõ ràng `project="<ProjectName>"` (hoặc `project="Global"`).
  - **Bảo Vệ Hai Lớp**: Mọi bản ghi bộ nhớ bắt buộc có tiền tố `[<ProjectName>: <Phân loại>] Nội dung chi tiết...` (ví dụ: `[FreeExile: Architecture] ...`, `[DSCons: Backend] ...`). Cơ chế này triệt tiêu 100% tình trạng rò rỉ hoặc trôi dạt ngữ cảnh giữa các dự án.

### 3. Máy Chủ MCP Bằng Ngôn Ngữ Go (Native Go-Compiled MCPs)
Các MCP server viết bằng Node.js hoặc Python thường tiêu tốn nhiều RAM và có độ trễ khởi động cao. Toàn bộ MCP cốt lõi trong hệ thống đã được viết lại bằng Go:
- **`mcp_git`**: Thao tác git (status, diff staged/unstaged, branch, commit) với tốc độ tức thì, tối ưu bộ nhớ gần như bằng 0.
- **`mcp_mem0_bridge`**: Cầu nối SSE tới Mem0 daemon có khả năng tự phục hồi (auto-healing) và tự động kích hoạt daemon nếu chưa chạy.
- **`mcp_memory`**: Lưu trữ và truy vấn đồ thị thực thể - quan hệ (knowledge graph) trên RAM với tốc độ cực đại.
- **`mcp_sqlite`**: Truy vấn nhanh, giải mã cấu trúc schema bảng dữ liệu SQLite mà không cần cài thêm driver cồng kềnh.

### 4. Quy Trình Review Mã Nguồn Lai (Alibaba Open Code Review Delegation)
Thay vì để Agent tự mò mẫm qua prompt tự nhiên gây ảo giác sai dòng, hệ thống áp dụng cơ chế ủy thác của Alibaba OCR qua 5 bước:
1. `ocr delegate preview`: Lọc sạch file rác/lockfile, xác định commit/merge-base chính xác.
2. `ocr delegate rule`: Kéo bộ quy chuẩn kiểm tra cấp doanh nghiệp (NPE, thread-safety, SQLi, tài nguyên rò rỉ).
3. `git diff`: Đọc nội dung thay đổi thực tế từ git.
4. **Đối Soát Chuyên Sâu**: Agent đối chiếu checklist, định vị chính xác khoảng dòng và phân loại mức độ (`critical`, `high`, `medium`, `low`).
5. **Tự Động Sửa Code (Auto-fix)**: Khắc phục trực tiếp các lỗi nghiêm trọng nếu có yêu cầu.

### 5. Native Impact Guard & Kiểm Soát Bán Kính Ảnh Hưởng (Blast Radius)
- **PreToolUse Hook (`native_impact_guard.py`)**: Can thiệp ngay trước khi Agent ghi hoặc sửa file (`replace_file_content`, `write_to_file`). Dùng bộ phân tích cú pháp AST quét toàn bộ workspace; nếu file sửa đổi có $\ge 6$ module phụ thuộc, hook sẽ lập tức phong tỏa lệnh sửa và yêu cầu mở khóa bằng cờ `--ack` nhằm bảo vệ mã nguồn dự án khỏi bị vỡ kiến trúc.
- **PostToolUse Hook (`aider_guard`)**: Binary Go thực hiện dọn dẹp RAM của các tiến trình Language Server chạy ngầm (nếu vượt 1.5 GB), đồng thời kiểm tra cú pháp JSON (chặn lỗi UTF-8 BOM) và Python ngay lập tức.

### 6. Cách Ly Nhánh Git Worktree & Vòng Lặp TDD Khép Kín
- **Ủy Thác Git Worktree**: Các đợt refactor quy mô lớn hoặc cập nhật thư viện rủi ro cao bắt buộc chạy trong workspace nhánh độc lập (`invoke_subagent` với `Workspace: 'branch'`). Agent chính chỉ merge khi subagent vượt qua 100% bài kiểm thử.
- **Chu Trình TDD Tự Chữa Lành**: Viết test trước -> Xác nhận FAIL -> Viết code tối thiểu -> Tự đọc traceback và sửa lặp lại (tối đa 5 lần không làm phiền người dùng) -> Atomic git commit khi PASS.

### 7. Tiêu Chuẩn Kỹ Thuật Chống AI Rác (Anti-AI Slop)
- **Modular & SRP**: Mỗi file/hàm chỉ đảm nhận một trách nhiệm duy nhất.
- **Tái Sử Dụng Triệt Để**: Tách logic chung vào thư mục `common`, `utils`, `helper`.
- **Độ Dài File Cứng (500–800 Dòng)**: Tuyệt đối không để file code vượt quá 500–800 dòng.
- **Zero Placeholder**: Nghiêm cấm `// TODO`, `/* giữ nguyên */`; toàn bộ code sinh ra phải chạy được ngay và xử lý tận gốc nguyên nhân lỗi.

---

## 📂 Cấu Trúc Thư Mục Kho Lưu Trữ

```
iloveantigravity/
├── install.ps1                 # Script cài đặt tự động cho Windows
├── install.sh                  # Script cài đặt tự động cho Linux/macOS
├── .env.example                # Mẫu biến môi trường và API key
├── configs/
│   ├── AGENTS.md               # Chỉ thị kỹ thuật toàn cục Antigravity
│   ├── GEMINI.md               # Quy tắc vận hành và nhân cách mô hình
│   ├── hooks.template.json     # Mẫu cấu hình Pre/PostToolUse hooks
│   ├── mcp_config.template.json# Mẫu cấu hình máy chủ MCP chuẩn hóa
│   └── workspace_template.json # Mẫu cấu hình workspace dự án mới
├── go_mcp/                     # Mã nguồn máy chủ Go MCP (Full source)
│   ├── git/                    # Go Git MCP server
│   ├── mem0_bridge/            # Go Mem0 SSE Client bridge
│   ├── memory/                 # Go In-Memory Knowledge Graph MCP
│   └── sqlite/                 # Go SQLite MCP server
├── go_hooks/                   # Mã nguồn Go Hook
│   └── aider_guard/            # Hook kiểm tra cú pháp & giải phóng RAM
├── scripts/                    # Các kịch bản lõi Python & PowerShell
│   ├── mem0_service.py         # Mem0 FastAPI daemon (Qdrant + AGY SDK)
│   ├── mem0_ctl.ps1            # Script điều khiển service (start/stop/status)
│   ├── mem0_stdio_bridge.py    # Bridge stdio dự phòng
│   ├── native_impact_guard.py  # Hook phân tích AST Blast Radius
│   ├── markitdown_mcp.py       # Trích xuất tài liệu sang Markdown
│   └── cleanup_antigravity.ps1 # Dọn rác cache, log tạm thời
├── skills/                     # Bộ kỹ năng tùy biến đặc biệt (chuẩn agentskills.io)
│   ├── invisible-playwright/   # Tự động hóa web tàng hình Firefox C++
│   └── skill-synthesizer/      # Tự động đúc kết kinh nghiệm thành SKILL.md
├── plugins/                    # 11 Plugin cao cấp đi kèm
│   ├── agentic-engineering/    # Worktree delegation, TDD, ast-grep
│   ├── developer-power-skills/ # Biểu đồ, Word/PDF/Excel, Vercel, tư duy giải quyết vấn đề
│   ├── elite-workflows-and-taste/# Thiết kế giao diện cao cấp, code review
│   ├── open-code-review/       # Alibaba Open Code Review integration
│   ├── openspec-plugin/        # Spec-driven development
│   └── science/                # Sinh học phân tử, AlphaFold, NCBI, ChEMBL
└── docs/                       # Tài liệu hướng dẫn chuyên sâu (VI, EN, ZH, JA)
```

---

## ⚙️ Biến Môi Trường Cấu Hình

Toàn bộ hệ thống vận hành hoàn hảo ngay lập tức (zero-config) mà không yêu cầu bắt buộc bất kỳ API key bên ngoài nào. Để tùy chỉnh cổng cho các dịch vụ nền, tạo file `.env` hoặc cấu hình biến môi trường:

| Biến Môi Trường | Mô Tả | Mặc Định |
| :--- | :--- | :--- |
| `MEM0_PORT` | Cổng dịch vụ SSE Mem0 cục bộ | `8765` |
| `MEM0_HOST` | Địa chỉ máy chủ Mem0 | `127.0.0.1` |

---

## 📚 Tài Liệu Hướng Dẫn Kỹ Thuật

- 📘 [Giao Thức Hermes & Khống Chế Bộ Nhớ](docs/vi/01_HERMES_PROTOCOL.md)
- 🧠 [Kiến Trúc Mem0 Cục Bộ & Quy Tắc Chống Loãng Dự Án](docs/vi/02_MEM0_LOCAL_MEMORY.md)
- ⚡ [Chi Tiết Máy Chủ Go MCP Siêu Tốc](docs/vi/03_GO_MCP_SERVERS.md)
- 🔍 [Quy Trình Review Code Alibaba OCR](docs/vi/04_CODE_REVIEW_OCR.md)
- 🛡️ [Hệ Thống Lifecycle Hooks & AST Blast Radius](docs/vi/05_LIFECYCLE_HOOKS.md)
- 🔄 [Hướng Dẫn Di Chuyển Máy Tính Nhanh](docs/vi/06_MIGRATION_GUIDE.md)

---

## 🙏 Ghi Nhận Tác Giả & Bản Quyền (Credits & Acknowledgements)

**iLoveAntigravity** được xây dựng và kế thừa tinh hoa từ các dự án mã nguồn mở, nghiên cứu học thuật và công cụ kỹ thuật xuất sắc của cộng đồng lập trình viên thế giới. Chúng tôi xin chân thành tri ân và ghi nhận đóng góp của:
- **Tập đoàn Alibaba (Alibaba Group)** ([open-code-review](https://github.com/alibaba/open-code-review)) với kiến trúc ủy thác review code tất định và bộ quy tắc kiểm tra sâu.
- **Herrington Darkholme** ([ast-grep](https://github.com/ast-grep/ast-grep)) với công cụ refactor và tìm kiếm AST theo cấu trúc cú pháp.
- **TheUpsider** ([@theupsider/lsp-mcp](https://github.com/theupsider/lsp-mcp)) với giải pháp kết nối Language Server Protocol qua MCP.
- **Tập đoàn Microsoft (Microsoft Corporation)** ([MarkItDown](https://github.com/microsoft/markitdown), [Playwright](https://github.com/microsoft/playwright)) với công cụ trích xuất tài liệu sang Markdown và tự động hóa trình duyệt web.
- **Đội ngũ Mem0.ai & Qdrant** ([mem0](https://github.com/mem0ai/mem0), [qdrant](https://github.com/qdrant/qdrant)) với mô hình bộ nhớ vector thích ứng và CSDL vector nhúng cục bộ.
- **Google DeepMind** với bộ công cụ Antigravity SDK, mô hình nền tảng Gemini, và hệ thống AlphaFold/AlphaGenome.
- **Paul Gauthier** ([Aider](https://github.com/paul-gauthier/aider)) với các chuẩn mực kiểm tra cú pháp AST và tối ưu luồng lập trình viên.
- **Nous Research** với triết lý khống chế bộ nhớ và tự phản tư đúc kết kỹ năng của giao thức Hermes Protocol.
- **Vercel, Cloudflare, OpenSpec, cùng cộng đồng gstack / Claude Code** với các bộ kỹ năng thực chiến, chuẩn thiết kế cao cấp và quy trình kỹ thuật chuẩn hóa.

Xem chi tiết danh sách đầy đủ tác giả, kho mã nguồn gốc, trích dẫn học thuật và giấy phép tại [CREDITS.md](CREDITS.md).

---

## 📄 Giấy Phép Sử Dụng (License)

Dự án phát hành theo giấy phép mã nguồn mở [MIT License](LICENSE). Hoàn toàn tự do cho mục đích cá nhân, học tập, nghiên cứu và thương mại trên toàn cầu.
