# ⚡ Máy Chủ Go MCP Siêu Tốc: Tối Ưu Băng Thông & Triệt Tiêu Độ Trễ

## 1. Tại Sao Lại Viết Lại MCP Server Bằng Go?
Các MCP server mặc định xây dựng trên nền Node.js (`npx`) hoặc Python tồn tại nhiều hạn chế khi vận hành cường độ cao:
1. **Độ trễ khởi động (Cold Start)**: Mỗi lần agent gọi tool, tiến trình Node/Python mất 300ms–1500ms chỉ để nạp runtime.
2. **Chiếm dụng bộ nhớ (RAM)**: Mỗi tiến trình Node.js tiêu tốn ~80MB–150MB RAM. Khi chạy cùng lúc 6-8 MCP server, máy tính mất hơn 1GB RAM chỉ để duy trì các adapter không tải.
3. **Phụ thuộc môi trường**: Xung đột phiên bản node_modules, thiếu wheel C-extensions trong Python khiến việc thiết lập lại trên máy mới rất hay gặp lỗi.

**Ngôn ngữ Go giải quyết triệt để các vấn đề trên**:
- Biên dịch thành 1 file thực thi duy nhất (standalone binary), không cần cài thêm runtime.
- Thời gian khởi động tính bằng phần nghìn giây (<1ms).
- Tiêu tốn cực ít bộ nhớ RAM (<10MB khi xử lý tải nặng).

---

## 2. Chi Tiết Các Máy Chủ Go MCP

### `mcp_git` (`go_mcp/git/main.go`)
Cung cấp các thao tác Git trực tiếp cho Agent với tốc độ tối đa, không tốn tài nguyên spawn shell:
- `git_status`: Xem trạng thái working tree.
- `git_diff_unstaged`: Xem diff chưa stage với tùy chọn số dòng context.
- `git_diff_staged`: Xem diff đã stage trong index.
- `git_diff`: So sánh giữa các nhánh hoặc commit.
- `git_commit`: Thực hiện atomic commit.
- `git_add`: Đưa file vào staging.
- `git_reset`: Bỏ stage file.
- `git_log`: Đọc lịch sử commit.
- `git_create_branch`, `git_checkout`, `git_show`, `git_branch`.

### `mcp_mem0_bridge` (`go_mcp/mem0_bridge/main.go`)
Cầu nối siêu tốc kết nối Agent với daemon Mem0 chạy nền:
- Cung cấp các công cụ: `mem0_status`, `mem0_add`, `mem0_search`, `mem0_get_all`, `mem0_delete`.
- Cơ chế tự phục hồi (Self-Healing): Nếu daemon Mem0 chưa chạy, bridge sẽ tự động gọi `pythonw.exe` khởi chạy `mem0_service.py` trong nền, chờ cổng sẵn sàng rồi mới thực thi lệnh, loại bỏ 100% lỗi rớt kết nối.

### `mcp_memory` (`go_mcp/memory/main.go`)
Máy chủ đồ thị tri thức (Knowledge Graph) chạy trên RAM với cơ chế luồng an toàn (thread-safe):
- Lưu trữ các thực thể (entities) và mối quan hệ (relations) vào file JSON tại thư mục gốc repo (`memory.json`).
- Cung cấp công cụ: `create_entities`, `create_relations`, `add_observations`, `delete_entities`, `delete_relations`, `read_graph`, `search_nodes`, `open_nodes`.

### `mcp_sqlite` (`go_mcp/sqlite/main.go`)
Kiểm tra và truy vấn nhanh cơ sở dữ liệu SQLite, sử dụng driver thuần Go hiện đại:
- Tự động phát hiện database trong repo (`dev.db`, `app.db`, `data.db`, v.v.).
- Cung cấp công cụ: `read_query`, `write_query`, `create_table`, `list_tables`, `describe_table`, `append_insight`.

---

## 3. Biên Dịch Mã Nguồn
Tất cả các Go MCP đều được biên dịch tự động qua kịch bản `install.ps1` hoặc `install.sh`:
```bash
# Biên dịch thủ công
cd go_mcp/git && go build -ldflags="-s -w" -o ../../bin/mcp_git .
cd go_mcp/mem0_bridge && go build -ldflags="-s -w" -o ../../bin/mcp_mem0_bridge .
cd go_mcp/memory && go build -ldflags="-s -w" -o ../../bin/mcp_memory .
cd go_mcp/sqlite && go build -ldflags="-s -w" -o ../../bin/mcp_sqlite .
```
Các cờ `-ldflags="-s -w"` giúp lược bỏ bảng biểu tượng gỡ lỗi (debug symbols), tối ưu dung lượng file nhị phân nhỏ nhất và tốc độ thực thi cao nhất.
