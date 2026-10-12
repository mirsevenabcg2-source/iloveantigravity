---
name: tdd-self-healing
description: >-
  Autonomous closed-loop Test-Driven Development (TDD) and self-healing verification stack.
  Use when implementing features, fixing bugs, or refactoring code to ensure zero human hand-holding
  through automated test-first execution and iterative self-repair.
---

# TDD Closed-Loop & Autonomous Self-Healing Protocol

## Overview
Quy chuẩn kiểm thử tự hành khép kín (Autonomous Verification Stack). Agent không bao giờ tuyên bố "đã làm xong" chỉ dựa trên suy đoán chủ quan. Mọi thay đổi mã nguồn phải trải qua chu trình kiểm chứng thực tế bằng lệnh thực thi test runner trong terminal.

---

## 5 Bước Vòng Lặp Tự Hành (Closed-Loop Workflow)

```
[User Request / Bug Report]
            │
            ▼
Step 1: Test-First Formulation
        (Viết test case / repro script trước khi sửa code)
            │
            ▼
Step 2: Negative Baseline Verification
        (Chạy test qua CLI -> BẮT BUỘC kết quả phải FAIL)
            │
            ▼
Step 3: Implementation
        (Viết mã nguồn nghiệp vụ tối thiểu để giải quyết vấn đề)
            │
            ▼
Step 4: Autonomous Self-Healing Loop
        ┌──► Chạy lại test suite
        │    │
        │    ├── FAIL / Traceback?
        │    │   └── Agent tự đọc stderr, sửa code, lặp lại (Zero hand-holding)
        │    │
        └─── └── PASS (Exit code 0)?
                 │
                 ▼
Step 5: Clean State & Git Commit
        (Commit sạch sẽ, lưu bài học kinh nghiệm vào Mem0)
```

---

## Chi tiết Thực thi

### 1. Test-First Formulation
- Trước khi chỉnh sửa bất kỳ file logic nào, tạo một test case phản ánh chính xác yêu cầu hoặc tái hiện bug (reproduction test).
- Sử dụng test framework có sẵn của dự án (`pytest`, `vitest`, `jest`, `cargo test`, `go test`, v.v.).

### 2. Negative Baseline Verification
- Chạy test ngay lập tức bằng `run_command`.
- **Tiêu chuẩn bắt buộc**: Test **phải FAIL**. Nếu test pass ngay từ đầu, test case đó là no-op hoặc không kiểm chứng đúng logic cần sửa.

### 3. Implementation
- Triển khai giải pháp mã nguồn. Tránh over-engineering; chỉ viết đủ code để thỏa mãn assertion của test case.

### 4. Self-Healing Loop (Không cần người dùng can thiệp)
- Tự động chạy lại test suite sau mỗi lần sửa.
- Nếu gặp lỗi:
  - Đọc kỹ `stderr`, exception type và line traceback.
  - Định vị chính xác lỗi logic hoặc cú pháp.
  - Sử dụng `replace_file_content` hoặc `ast-grep` để sửa code.
  - Tự động chạy lại test ngay lập tức.
- **Circuit Breaker**: Giới hạn tối đa 5 vòng lặp tự sửa liên tiếp. Nếu sau 5 lần vẫn chưa pass, Agent tổng hợp chi tiết phân tích lỗi và đề xuất phương án cho người dùng.

### 5. Hoàn tất & Atomic Git Commit
- Khi test đã PASS hoàn toàn:
  - Chạy lại toàn bộ test suite để đảm bảo không gây regression (lỗi liên đới).
  - Sử dụng `git add` và `git commit` tạo commit rõ ràng, mang tính nguyên tử (atomic).
  - Tự động gọi `mem0_add` ghi nhận cách giải quyết bài toán nếu đó là bug tinh vi hoặc quy ước kiến trúc mới.
