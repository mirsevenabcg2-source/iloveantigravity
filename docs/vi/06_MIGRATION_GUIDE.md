# 🔄 Hướng Dẫn Di Chuyển Môi Trường Giữa Các Máy Tính (Migration Guide)

## 1. Tổng Quan
Bất kỳ khi nào bạn sắm máy tính mới, chuyển đổi giữa laptop và máy bàn ở công ty/nhà, hoặc cài lại Windows/Linux/macOS, **iLoveAntigravity** cho phép bạn tái lập 100% môi trường làm việc tối ưu của Agent trong chưa đầy 60 giây mà không cần cấu hình thủ công từng bước.

---

## 2. Các Bước Thực Hiện Cụ Thể

### Bước 1: Cài Đặt Công Cụ Nền Tảng
Đảm bảo máy tính mới đã cài sẵn:
1. **Git**: [https://git-scm.com/](https://git-scm.com/)
2. **Trình Biên Dịch Go (1.22+)**: [https://go.dev/dl/](https://go.dev/dl/)
3. **Python (3.10+)**: [https://python.org/](https://python.org/) (nhớ tích chọn *Add python.exe to PATH* khi cài đặt)
4. **Node.js (18+)**: [https://nodejs.org/](https://nodejs.org/)

### Bước 2: Clone Kho Mã Nguồn Về Máy
```bash
git clone https://github.com/mirsevenabcg2-source/iloveantigravity.git
cd iloveantigravity
```

### Bước 3: Chạy Script Cài Đặt Tự Động (1-Click)

#### Trên Windows (PowerShell):
```powershell
# Chạy script cài đặt (khuyến khích mở PowerShell bình thường hoặc Run as Administrator)
.\install.ps1
```

#### Trên Linux hoặc macOS (Bash):
```bash
chmod +x install.sh
./install.sh
```

### Bước 4: Tùy Chỉnh Cấu Hình Môi Trường (Tùy Chọn)
Toàn bộ hệ thống hoạt động ngay lập tức (zero-config) hoàn toàn cục bộ mà không bắt buộc bất kỳ API key bên ngoài nào. Nếu cần tùy chỉnh cổng dịch vụ Mem0, khai báo trong file `~/.gemini/antigravity/mem0_data/.env`:
```bash
MEM0_PORT=8765
MEM0_HOST=127.0.0.1
```

### Bước 5: Kiểm Tra & Bắt Đầu Sử Dụng
1. Khởi động lại **Google Antigravity IDE** hoặc CLI.
2. Kiểm tra danh sách MCP server đã nạp thành công:
   - `git` (Go Native siêu tốc)
   - `mem0` (Go bridge tới vector memory cục bộ)
   - `memory` (Go đồ thị tri thức trên RAM)
   - `sqlite` (Go thanh tra CSDL)
   - `ast-grep` & `lsp-mcp`
   - `invisible-playwright`
3. Kiểm tra trạng thái dịch vụ nền Mem0:
   ```powershell
   powershell -File ~/.gemini/config/scripts/mem0_ctl.ps1 -Action status
   ```
4. Toàn bộ kho kinh nghiệm, thói quen code, và năng lực của Agent đã sẵn sàng phục vụ bạn trên máy tính mới!
