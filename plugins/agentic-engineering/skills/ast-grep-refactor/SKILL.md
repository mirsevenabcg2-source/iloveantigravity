---
name: ast-grep-refactor
description: >-
  Structural syntax-aware code search, mass refactoring, and AST rule enforcement
  using ast-grep and ast-grep-mcp. Use when searching code by syntax structure, replacing
  patterns across multiple files, or modernizing deprecated API calls without regex hallucinations.
---

# AST-Grep Structural Refactoring Protocol

## Overview
Thay vì tìm kiếm và thay thế mã nguồn bằng Regular Expressions (Regex) dễ dẫn tới sai sót, ảo giác hoặc phá vỡ cấu trúc cú pháp, Agent sử dụng **ast-grep** để phân tích trực tiếp trên Cây Cú Pháp Trừu Tượng (AST/CST).

---

## Khi nào nên sử dụng
- **Refactor diện rộng**: Đổi tên hàm, thay đổi thứ tự tham số trên toàn bộ codebase.
- **Hiện đại hóa mã nguồn**: Chuyển đổi promise chaining sang `async/await`, migrate API thư viện cũ sang phiên bản mới.
- **Rà soát bảo mật cú pháp**: Tìm kiếm các lời gọi hàm nguy hiểm hoặc thiếu guard clause (ví dụ: `eval()`, SQL query concatenations).

---

## Cú pháp Pattern Cơ bản

Ký hiệu | Ý nghĩa | Ví dụ
:--- | :--- | :---
`$VAR` | Khớp đúng một node cú pháp đơn | `const $A = require('$B')`
`$$$ARGS` | Khớp một danh sách node (0 hoặc nhiều node) | `console.log($$$ARGS)`
`$_` | Node ẩn danh (không cần lưu lại giá trị) | `function $_($$$ARGS) { $$$BODY }`

---

## Quy trình Thực thi An toàn (2-Phase Safety Gate)

```
[Target Pattern & Rewrite Rule]
              │
              ▼
Phase 1: Preview Mode (apply: false)
         Gọi ast_grep_search kiểm tra kết quả khớp, diff preview
              │
              ├── Kết quả khớp chính xác 100%?
              │   └── Không: Tinh chỉnh pattern/selector
              │
              ▼
Phase 2: Mutation Mode (apply: true)
         Áp dụng thay đổi trực tiếp vào filesystem
              │
              ▼
Verification: Chạy compiler / lsp_diagnostics / test runner
```

### Ví dụ Sử dụng qua Tool:

#### 1. Preview thay thế cú pháp (Safe Check):
```json
{
  "pattern": "$A && $A()",
  "rewrite": "$A?.()",
  "lang": "typescript",
  "apply": false,
  "paths": ["src"]
}
```

#### 2. Áp dụng thay đổi thực tế:
```json
{
  "pattern": "$A && $A()",
  "rewrite": "$A?.()",
  "lang": "typescript",
  "apply": true,
  "paths": ["src"]
}
```
Sau khi apply thành công, kết hợp gọi `lsp_diagnostics` hoặc chạy unit test để đảm bảo hệ thống không có lỗi biên dịch.
