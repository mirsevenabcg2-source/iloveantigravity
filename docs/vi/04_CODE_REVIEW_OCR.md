# 🔍 Quy Trình Review Mã Nguồn Lai (Alibaba Open Code Review Delegation)

## 1. Tại Sao Lại Dùng Chế Độ Ủy Thác (Delegation Mode)?
Các phương pháp review code bằng AI thông thường hay gặp phải hai vấn đề nghiêm trọng:
1. **Ảo giác lệch dòng (Line Hallucination)**: Mô hình ngôn ngữ thường đếm sai số dòng khi duyệt các file diff dài, dẫn tới việc chỉ định sai vị trí lỗi và sinh mã sửa đổi bị lệch ngữ cảnh.
2. **Lãng phí token & chi phí**: Gửi toàn bộ file lên các API bên ngoài khiến hàng triệu token bị tiêu tốn vô ích vào các file rác, file lock (`package-lock.json`), hay thư mục thư viện bên ngoài.

Mô hình **Delegation Mode** với công cụ **Open Code Review (Alibaba Group)** giải quyết bài toán này bằng kiến trúc phân tách rõ ràng:
- Tầng kỹ thuật xác định (Deterministic Engineering) đảm nhận lọc file, tìm merge-base và kéo bộ quy chuẩn doanh nghiệp.
- Tầng suy luận sâu của Gemini Host Agent đảm nhận phân tích logic, định vị chính xác dòng, phân loại mức độ nghiêm trọng và sửa lỗi tự động.

---

## 2. Quy Trình Thực Thi 5 Bước Chuẩn Hóa

```
[ Git Diff / Workspace ]
           │
           ▼
Bước 1: ocr delegate preview --format json
           │ (Lọc bỏ file rác/lockfile, xác định danh sách file và merge-base)
           ▼
Bước 2: ocr delegate rule --format json <files...>
           │ (Kéo checklist quy chuẩn doanh nghiệp: NPE, Luồng an toàn, SQLi, Resource Leaks)
           ▼
Bước 3: git diff <merge_base>..<to> -- <path>
           │ (Lấy nội dung thay đổi thực tế từ Git)
           ▼
Bước 4: Gemini Suy Luận Sâu & Đối Soát
           │ (Duyệt theo từng nhóm quy tắc, định vị chính xác dòng, phân loại severity)
           ▼
Bước 5: Báo Cáo Chuẩn & Tự Động Sửa (Auto-fix)
           │ (Sửa trực tiếp lỗi High/Critical nếu người dùng yêu cầu)
```

---

## 3. Cấu Trúc Báo Cáo Lỗi Tiêu Chuẩn
Mọi lỗi phát hiện trong quá trình review đều phải chuẩn hóa theo cấu trúc:
- **File & Vị Trí Dòng**: `path/to/file.ext#L<start>-L<end>`
- **Phân Loại Mức Độ (Severity)**:
  - `critical`: Lỗ hổng bảo mật nghiêm trọng (SQL Injection, RCE, vượt quyền), rò rỉ tài nguyên hệ thống lớn.
  - `high`: Lỗi logic nghiệp vụ, nguy cơ Null Pointer Exception (NPE), Race Condition đa luồng, mất mát dữ liệu.
  - `medium`: Suy giảm hiệu năng, truy vấn thiếu index, vi phạm mẫu thiết kế (anti-pattern).
  - `low`: Quy chuẩn đặt tên, phong cách code, gợi ý cải thiện khả năng bảo trì.
- **Nguyên Nhân Gốc Rễ & Đoạn Code Sửa Đổi**: Phân tích kỹ thuật chuyên sâu kèm đoạn diff chạy được ngay.
