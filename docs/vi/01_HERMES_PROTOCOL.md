# 🏛️ Giao Thức Hermes: Tối Ưu Hóa Ngữ Cảnh & Khống Chế Bộ Nhớ

## 1. Vấn Nạn Trôi Dạt Ngữ Cảnh (Context Dilution)
Trong quá trình lập trình thực tế, các AI Agent thường nhanh chóng bị suy giảm chất lượng suy luận do:
- System prompt bị nhồi nhét quá nhiều quy tắc, ví dụ và schema phức tạp.
- Kết quả trả về từ công cụ (`tool_result`) như log build, danh sách thư viện, git diff hàng ngàn dòng chiếm trọn cửa sổ ngữ cảnh.
- Mô hình gặp hiện tượng "loãng chú ý", quên mất các quyết định kiến trúc đã thống nhất trước đó và bắt đầu sinh mã lỗi (hallucination).

## 2. Giải Pháp Giao Thức Hermes
**Giao thức Hermes** thiết lập 4 chốt chặn kỹ thuật bắt buộc:

### Chốt Chặn A: Bounded Memory Snapshot (<2,500 Ký Tự)
- Giữ System Prompt khởi tạo (`AGENTS.md`) luôn dưới 2,500 ký tự.
- Toàn bộ tài liệu chi tiết, thói quen code, cấu trúc CSDL được chuyển sang lưu trữ tại **Mem0** (vector embedding) và SQLite FTS5.
- Tri thức được nạp động theo cơ chế *Just-In-Time* (khi cần mới tra cứu) thay vì gánh nặng thường trực.

### Chốt Chặn B: Chủ Động Nén Ngữ Cảnh @ 50% (Proactive Context Compaction)
- Khi dung lượng phiên trò chuyện đạt 50% ngưỡng cửa sổ context:
  - Agent chủ động tóm tắt các khối dữ liệu tool cồng kềnh (stdout build, diff lớn) thành 2-3 dòng súc tích.
  - Bảo toàn tuyệt đối các quyết định kỹ thuật cốt lõi và traceback quan trọng, loại bỏ rác văn bản.
- Não bộ AI luôn duy trì được độ tập trung và nhạy bén tối đa.

### Chốt Chặn C: Tự Động Đúc Kết Kỹ Năng (`skill-synthesizer`)
- Khi Agent giải quyết thành công một tác vụ kỹ thuật phức tạp (cấu hình toolchain lạ, xử lý lỗi hệ điều hành chuyên sâu, tích hợp pipeline mới), Agent tự động phản tư.
- Quy trình chuẩn được đúc kết thành file `SKILL.md` theo chuẩn `agentskills.io` đặt trong `~/.gemini/config/skills/<name>/`.
- Các phiên làm việc sau sẽ tự động kế thừa tri thức này mà không cần giải thích lại từ đầu.

### Chốt Chặn D: Tiêu Chuẩn Kỹ Thuật Chống AI Slop (Anti-AI Slop)
- **Trần Cứng 500–800 Dòng**: Không một file mã nguồn hoặc tài liệu nào được phép vượt quá 800 dòng. Khi tiệm cận ngưỡng này, bắt buộc tách module theo nguyên lý Đơn Trách Nhiệm (SRP).
- **Tuyệt Đối Không Code Giả / Placeholder**: Nghiêm cấm `// TODO`, `/* giữ nguyên */`, `catch` rỗng. Mã nguồn sinh ra phải chạy được ngay trong thực tế.
- **Chọn Đúng Ngôn Ngữ**: Go/Rust/C++ cho tầng hệ thống, công cụ tốc độ cao; Python cho AI/Data/AST; TypeScript cho giao diện Web.
