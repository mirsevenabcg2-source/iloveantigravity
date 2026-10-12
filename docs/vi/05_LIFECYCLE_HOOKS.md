# 🛡️ Hệ Thống Lifecycle Hooks & Kiểm Soát Bán Kính Ảnh Hưởng (Blast Radius)

## 1. Tổng Quan
Google Antigravity cung cấp cơ chế Hook vòng đời (`PreToolUse`, `PostToolUse`) cho phép can thiệp trước và sau khi Agent thực thi công cụ. **iLoveAntigravity** tận dụng hai điểm can thiệp này để bảo vệ toàn vẹn kiến trúc mã nguồn và kiểm soát tài nguyên hệ thống.

---

## 2. PreToolUse: Native Impact Guard (`scripts/native_impact_guard.py`)
Can thiệp trước các công cụ sửa đổi file:
- `write_to_file`
- `replace_file_content`

### Cơ Chế Hoạt Động:
1. **Nhận Diện File Đích**: Trích xuất đường dẫn file mà Agent chuẩn bị chỉnh sửa.
2. **Bộ Lọc An Toàn**: Tự động bỏ qua các file tài liệu Markdown, text, file cấu hình, log, và file trong thư mục test/scratch.
3. **Quét Cây Cú Pháp AST (AST Blast Radius Scan)**:
   - Quét toàn bộ mã nguồn trong workspace dự án.
   - Phân tích cú pháp AST để đếm chính xác số lượng module khác đang `import` hoặc phụ thuộc vào file đích này.
4. **Ngưỡng Phong Tỏa (Hard Gate Threshold)**:
   - Nếu file ảnh hưởng đến **$\ge 6$ module khác**, hook lập tức từ chối (`deny`) quyền ghi đè của Agent.
   - **Quy Trình Mở Khóa**: Để cho phép sửa đổi trong các đợt refactor lớn có chủ đích, chạy lệnh:
     ```bash
     python tools/analysis/blast_radius.py --target <đường_dẫn_file> --ack
     ```
     Lệnh này sẽ cấp quyền sửa đổi an toàn trong vòng 30 phút cho file đó.

---

## 3. PostToolUse: Aider Native Guard (`go_hooks/aider_guard`)
Hook biên dịch bằng Go chạy ngay sau khi file vừa được ghi xong:

### Chức Năng Chính:
1. **Kiểm Soát Cú Pháp Ngay Lập Tức**:
   - Quét file JSON: Chặn đứng ký tự ẩn UTF-8 BOM (`\xEF\xBB\xBF`) vốn thường làm hỏng các luồng build Webpack/Node.
   - Quét file Python: Chạy `py_compile` để phát hiện lỗi cú pháp trước khi chương trình thực thi.
2. **Tự Động Thu Hồi RAM Language Server**:
   - Sử dụng Win32 API (`CreateToolhelp32Snapshot`) quét các tiến trình `language_server.exe` chạy ngầm.
   - Nếu tiến trình nào chiếm dụng quá 1,500 MB RAM, hook tự động gọi `EmptyWorkingSet` để thu hồi bộ nhớ vật lý, chống tràn RAM máy tính.
