---
name: skill-synthesizer
description: "Tự động phân tích, chắt lọc và tạo mới hoặc cập nhật các kỹ năng (skills) dưới dạng SKILL.md theo chuẩn agentskills.io từ quy trình giải quyết vấn đề thành công trong phiên làm việc. Kế thừa triết lý Closed Learning Loop từ Hermes Agent. Sử dụng khi: 'tạo skill mới', 'đúc kết kỹ năng', 'synthesize skill', 'lưu quy trình thành skill', 'skill manage', 'cập nhật skill', hoặc tự động kích hoạt sau khi giải quyết xong một bài toán kỹ thuật phức tạp."
---

# Skill Synthesizer (Autonomous Skill Synthesis & Progressive Disclosure Engine)

Quy chuẩn kỹ thuật và chu trình tự hành khép kín (**Autonomous Skill Synthesis**) lấy cảm hứng từ kiến trúc học tập của **Hermes Agent (Nous Research)**. Kỹ năng này cho phép Agent biến kinh nghiệm thực chiến thành tài sản tri thức tái sử dụng lâu dài, ngăn ngừa việc phải "dò lại đường từ đầu" trong các phiên tương lai.

---

## 1. NGUYÊN TẮC CỐT LÕI (CORE PRINCIPLES)

1. **Progressive Disclosure (Mở khóa Ngữ cảnh Tiệm tiến)**:
   - Kỹ năng được lưu dưới dạng file `SKILL.md`.
   - Phần frontmatter YAML (`name`, `description`) chỉ chứa thông tin nhận diện và các từ khóa trigger.
   - Toàn bộ nội dung hướng dẫn chi tiết chỉ được nạp vào context khi Agent thực sự cần thực thi tác vụ đó. Giữ cho System Prompt luôn tinh gọn (< 2,500 ký tự).
2. **Empirical Grounding (Dựa trên Thực chứng)**:
   - Tuyệt đối KHÔNG sáng tác kỹ năng từ lý thuyết giả định.
   - Mọi kỹ năng sinh ra phải dựa trên chuỗi lệnh, mã nguồn hoặc thao tác debug đã chạy **THÀNH CÔNG** trên hệ thống thực tế.
3. **Action-Driven, Zero AI Slop**:
   - Cấm các câu sáo rỗng ("Certainly", "This guide will help you").
   - Kỹ năng phải chứa lệnh thực thi chính xác (CLI, API, Code Snippets) và cách giải quyết triệt để các bẫy lỗi (pitfalls).

---

## 2. KHI NÀO CẦN KÍCH HOẠT ĐÚC KẾT KỸ NĂNG (TRIGGER CONDITIONS)

Agent chủ động đề xuất hoặc tự động thực hiện đúc kết kỹ năng khi gặp các tình huống:
- **Tình huống 1 (Novel Bug Resolution)**: Vừa giải quyết xong một lỗi hệ thống hiếm gặp, lỗi xung đột thư viện hoặc lỗi runtime phức tạp mất nhiều bước điều tra.
- **Tình huống 2 (New Integration / Setup)**: Vừa xây dựng thành công một pipeline build, tích hợp MCP Server mới, cấu hình công cụ dev hoặc reverse-engineering một target mới.
- **Tình huống 3 (User Explicit Request)**: Người dùng ra lệnh "lưu cái này thành skill", "đúc kết quy trình", "nhớ cách làm này cho lần sau".

---

## 3. CẤU TRÚC CHUẨN CỦA MỘT TỆP `SKILL.md`

Mỗi kỹ năng được đóng gói trong thư mục riêng: `<skill-directory>/<skill-name>/SKILL.md`.

```markdown
---
name: <ten-skill-kebab-case>
description: "<Mô tả ngắn gọn chức năng trong 1-2 câu. BẮT BUỘC liệt kê các từ khóa kích hoạt: 'từ khóa 1', 'từ khóa 2', 'intent-action'>."
---

# <Tên Kỹ Năng Đầy Đủ>

Mô tả ngắn gọn về giá trị và mục tiêu của kỹ năng này.

## 1. Điều Kiện Tiên Quyết (Prerequisites)
- Các công cụ CLI, biến môi trường, hoặc quyền hạn cần thiết.

## 2. Quy Trình Thực Thi Chuẩn (Step-by-Step Workflow)
1. **Bước 1: Khảo sát / Chuẩn bị**:
   - Lệnh thực thi cụ thể.
2. **Bước 2: Triển khai / Thực thi**:
   - Lệnh hoặc mã nguồn chuẩn.
3. **Bước 3: Kiểm chứng (Verification)**:
   - Lệnh test hoặc tiêu chí nghiệm thu PASS/FAIL rõ ràng.

## 3. Các Lỗi Thường Gặp & Cách Khắc Phục (Pitfalls & Troubleshooting)
- **Lỗi A**: Nguyên nhân và cách xử lý nhanh.
- **Lỗi B**: Cách phòng tránh.
```

---

## 4. VỊ TRÍ LƯU TRỮ VÀ QUẢN LÝ PHẠM VI (SCOPE MANAGEMENT)

1. **Kỹ năng Toàn cục (Global Skills)**:
   - Dành cho các công cụ hệ thống, quy chuẩn code toàn cục, hoặc các framework dùng chung xuyên suốt mọi dự án.
   - Đường dẫn: `C:\Users\Admin\.gemini\config\skills\<skill-name>\SKILL.md`.
2. **Kỹ năng Riêng của Dự án (Project-Scoped Skills)**:
   - Dành cho các quy tắc domain nghiệp vụ, cấu trúc schema riêng biệt, hoặc workflow chỉ áp dụng trong repo hiện tại.
   - Đường dẫn: `<ProjectRoot>/.agents/skills/<skill-name>/SKILL.md`.

---

## 5. QUY TRÌNH TỰ ĐỘNG THỰC THI (SYNTHESIS WORKFLOW)

Khi phát hiện cơ hội đúc kết kỹ năng:
1. **Trích xuất chuỗi hành động (Trajectory Extraction)**:
   - Rà soát các lệnh CLI đã chạy và tìm ra đường đi ngắn nhất dẫn tới kết quả thành công (loại bỏ các lệnh thử-sai vô nghĩa).
2. **Soạn thảo tài liệu**:
   - Viết tệp `SKILL.md` theo đúng cấu trúc mục 3.
   - Đảm bảo frontmatter YAML có đủ tên và bộ từ khóa trigger tự nhiên tiếng Việt & tiếng Anh.
3. **Lưu trữ & Xác nhận**:
   - Ghi file xuống ổ đĩa bằng `write_to_file`.
   - Báo cáo ngắn gọn cho người dùng: tên kỹ năng mới, vị trí lưu, và các từ khóa trigger kích hoạt trong tương lai.
