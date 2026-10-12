---
name: git-worktree-delegation
description: >-
  Orchestrate heavy refactoring, complex migrations, and high-risk coding tasks
  by delegating to isolated subagents running in dedicated Git Worktrees (Workspace: 'branch').
  Use when tasks affect multiple modules, risk breaking existing features, or require clean sandbox verification.
---

# Git Worktree Delegation Protocol (Antigravity 2.0)

## Overview
Khi đối mặt với các bài toán refactoring diện rộng, nâng cấp dependencies lớn, hoặc thử nghiệm kiến trúc mới, Agent **BẮT BUỘC** không được can thiệp trực tiếp làm xáo trộn working copy chính. Thay vào đó, áp dụng cơ chế **Git Worktree Delegation**: chia việc cho Subagent chạy trong môi trường phân nhánh cô lập (`Workspace: 'branch'`).

---

## When to Use
- **Refactoring đa tầng**: Thay đổi cấu trúc folder, đổi tên API cốt lõi, migrate framework.
- **Thử nghiệm rủi ro cao**: Nâng cấp breaking packages, viết lại thuật toán lõi.
- **Thực thi song song (Parallel Exploration)**: Thử 2 cách tiếp cận khác nhau để giải quyết 1 bug khó.

---

## Protocol Steps

```
[Main Working Copy]
         │
         ▼
Step 1: Xác định phạm vi & Ranh giới an toàn
         │
         ▼
Step 2: invoke_subagent(TypeName="self", Workspace="branch", Role="Feature Specialist", ...)
         │ (Tạo git worktree cô lập)
         ▼
Step 3: Subagent tự do code, test, fix lỗi trong sandbox
         │ (Không ảnh hưởng working copy chính)
         ▼
Step 4: Subagent hoàn thành & gửi báo cáo kèm Git Diff
         │
         ▼
Step 5: Host Agent đối soát kết quả, chạy verification pass rồi merge
```

### 1. Khởi tạo Subagent với Workspace Cô lập
Gọi công cụ `invoke_subagent` với tham số `Workspace: 'branch'`:
```json
{
  "Subagents": [
    {
      "TypeName": "self",
      "Role": "Refactor Specialist",
      "Workspace": "branch",
      "Prompt": "Thực hiện refactor module X. Yêu cầu: 1) Viết test kiểm chứng trước; 2) Triển khai mã nguồn; 3) Đảm bảo test suite xanh 100% trước khi báo cáo hoàn thành."
    }
  ]
}
```

### 2. Nguyên tắc Không Busy-Poll (Reactive Notification)
- Sau khi khởi chạy subagent, Host Agent dừng gọi tool hoặc tiếp tục xử lý công việc độc lập khác.
- Hệ thống Antigravity sẽ tự động đánh thức Host Agent khi subagent hoàn thành hoặc gửi message phản hồi.

### 3. Kiểm định & Hợp nhất (Verification & Merge)
- Khi nhận được thông báo hoàn thành từ subagent:
  1. Kiểm tra diff mà subagent đã tạo trên nhánh git cô lập.
  2. Xác nhận mọi unit test và regression test đều đã pass.
  3. Tiến hành merge kết quả vào nhánh chính hoặc hướng dẫn người dùng review.
- Nếu subagent thất bại hoặc gây lỗi: Dễ dàng hủy bỏ (kill subagent) mà không để lại bất kỳ file rác nào trong working copy chính.
