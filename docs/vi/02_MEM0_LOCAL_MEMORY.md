# 🧠 Kiến Trúc Bộ Nhớ Mem0 Cục Bộ & Quy Tắc Cô Lập Dự Án (Anti-Dilution)

## 1. Tại Sao Cần Bộ Nhớ Vector Cục Bộ?
Khi đóng phiên làm việc, Agent thông thường sẽ mất toàn bộ ngữ cảnh về sở thích cá nhân, các quyết định kiến trúc và bài học kinh nghiệm. Việc dùng bộ nhớ đám mây bên thứ ba có 3 nhược điểm chí mạng:
1. Rò rỉ mã nguồn bí mật và khóa API ra máy chủ bên ngoài.
2. Phát sinh chi phí duy trì hàng tháng.
3. Thiếu cơ chế phân tách dự án, dẫn tới việc trộn lẫn tai hại (ví dụ: áp dụng thói quen hook C++ game reversing vào dự án ERP/Revit C#).

## 2. Ngăn Xếp Mem0 Trong iLoveAntigravity
Hệ thống Mem0 chạy thường trực ngầm trên máy:
- **Cổng giao tiếp**: `http://127.0.0.1:8765/sse` qua Starlette & FastMCP.
- **Cơ sở dữ liệu Vector**: Embedded Qdrant tại `~/.gemini/antigravity/mem0_data/qdrant`.
- **Bộ nhúng (Embedding)**: Cục bộ Cosine 768 chiều.
- **Động cơ suy luận (LLM Engine)**: Tích hợp trực tiếp với **Google Antigravity Python SDK (`google.antigravity.Agent`)**.
  - **0% VRAM GPU**: Tuyệt đối không cần chạy các LLM cục bộ nặng nề (LM Studio, Ollama).
  - Tự động trích xuất và tự phân giải mâu thuẫn (self-updating memory).
- **Cơ chế khởi động**: Tự chạy khi mở máy qua Windows Registry `Run` (`HKCU:\Software\Microsoft\Windows\CurrentVersion\Run\Mem0_Antigravity_Service`). Điều khiển qua `scripts/mem0_ctl.ps1`.

---

## 3. Quy Tắc Bắt Buộc: Cô Lập Project (Zero Dilution)
Để triệt tiêu hiện tượng loãng bộ nhớ giữa các kho mã nguồn độc lập, Agent bắt buộc tuân thủ 2 lớp bảo vệ:

### Lớp 1: Bắt Buộc Chỉ Định Project Scope
Tuyệt đối CẤM gọi `mem0_search` hoặc `mem0_add` mà không truyền tham số `project`:
```python
# Tra cứu quy ước của repo cụ thể
mem0_search(query="Quy tắc migration database", project="DSCons")

# Tra cứu thói quen chung toàn cục
mem0_search(query="Phong cách comment và đặt tên hàm", project="Global")
```

### Lớp 2: Tiền Tố Cấu Trúc Hai Tầng
Mọi nội dung lưu trữ qua `mem0_add` phải có tiền tố nhận diện dự án:
```python
mem0_add(
    text="[BackendService: Architecture] Mọi migration database bắt buộc phải tương thích ngược (backwards-compatible).",
    project="BackendService"
)
```
Cơ chế bảo vệ 2 lớp này giúp vector trong Qdrant được phân vùng rành mạch, triệt tiêu 100% tình trạng rò rỉ ngữ cảnh chéo.

---

## 4. Quản Lý Dịch Vụ Với mem0_ctl.ps1
```powershell
# Kiểm tra trạng thái đang chạy và cổng lắng nghe
.\scripts\mem0_ctl.ps1 -Action status

# Khởi động dịch vụ
.\scripts\mem0_ctl.ps1 -Action start

# Khởi động lại dịch vụ
.\scripts\mem0_ctl.ps1 -Action restart

# Dừng dịch vụ
.\scripts\mem0_ctl.ps1 -Action stop

# Đăng ký tự khởi động cùng Windows
.\scripts\mem0_ctl.ps1 -Action install-autostart
```
