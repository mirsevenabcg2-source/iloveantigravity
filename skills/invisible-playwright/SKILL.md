---
name: invisible-playwright
description: Stealth web automation and scraping using invisible-playwright (anti-detect Firefox C++ engine). Use for browsing websites, bypassing bot detection/Cloudflare/CAPTCHA, testing web UI, capturing screenshots, and automating browser interactions.
---

# Invisible Playwright Stealth Browser Skill

Hệ thống duyệt web tàng hình (Stealth Web Automation) sử dụng trình duyệt Firefox được vá trực tiếp ở tầng C++ (Camoufox engine, giao thức Juggler in-process) để tạo fingerprint phần cứng thật (canvas, WebGL, audio, font metrics, navigator), vượt qua các hệ thống phát hiện bot (Cloudflare Turnstile, DataDome, Akamai, PerimeterX) mà không cần chèn script JavaScript gây lộ `isTrusted: false`.

---

## 1. Nguyên Tắc Vận Hành Cốt Lõi

1. **Bắt buộc gọi `browser_open` đầu tiên**: Không có tool nào tự động mở trình duyệt nếu chưa được khởi tạo.
2. **Quy tắc 2 trình duyệt (`main` & `support`)**:
   - `main`: Trình duyệt làm việc chính, lưu giữ session/cookies/fingerprint. Mọi lệnh không truyền `browser` sẽ tự động trỏ về `main`.
   - `support`: Trình duyệt phụ dùng để mở các trang trung gian, email tạm thời (throwaway mailbox) mà không làm bẩn identity của `main`. **Bắt buộc đóng ngay bằng `browser_close(browser="support")`** sau khi dùng xong để giải phóng RAM.
3. **Thứ tự tương tác chuẩn (Interaction Ladder)**:
   - **Bậc 1 (Selector)**: Dùng selector từ `browser_snapshot` và gọi `browser_click`, `browser_type`, `browser_select_option`, `browser_press_key`.
   - **Bậc 2 (Tọa độ chuột thật)**: Dùng tọa độ viewport `[x, y]` từ `browser_snapshot` và gọi `browser_click_at` (thao tác chuột thật theo đường cong Bezier, hỗ trợ `hold_seconds`).
   - **Bậc 3 (Thị giác AI)**: Chụp ảnh `browser_take_screenshot` để xác định vị trí các phần tử canvas/map/slider rồi click bằng `browser_click_at`.
   - **Bậc 4 (Read-Only JS)**: `browser_evaluate` chỉ dùng để **ĐỌC** dữ liệu (read-only state, computed styles, text ẩn). Tuyệt đối **không** dùng để gán `value`, `checked` hoặc gọi `click()`, `dispatchEvent()`, `submit()` vì sẽ tạo ra sự kiện giả tạo `isTrusted: false` khiến trang phát hiện bot ngay lập tức.

---

## 2. Danh Sách Công Cụ & Thao Tác

| Công cụ | Tham số chính | Mục đích & Lưu ý |
| :--- | :--- | :--- |
| `browser_open` | `browser`, `seed`, `proxy`, `profile` | Khởi tạo phiên duyệt. Không truyền tham số sẽ khôi phục identity cũ từ profile. |
| `browser_close` | `browser` | Đóng trình duyệt và giải phóng tiến trình `firefox.exe`. |
| `browser_list` | *(không)* | Xem danh sách trình duyệt đang mở, URL hiện tại và focus. |
| `browser_status` | `browser` | Kiểm tra seed, IP exit, profile directory và URL hiện tại. |
| `browser_navigate` | `url`, `wait_until`, `browser` | Điều hướng URL. `wait_until`: `"domcontentloaded"` (mặc định), `"load"` (chờ ảnh/css), `"networkidle"` (SPA/React/Vue). |
| `browser_snapshot` | `max_chars`, `browser` | Trích xuất toàn bộ phần tử tương tác kèm selector duy nhất và tọa độ `[x, y]`. |
| `browser_read_text` | `selector`, `max_chars`, `browser` | Đọc text thô đã làm sạch. Mặc định `body`, giới hạn 6000 ký tự. |
| `browser_read_html` | `mode`, `browser` | Lấy HTML đã tinh lọc: `mode="form"` (form/inputs), `mode="text"` (văn bản), `mode="full"` (toàn bộ cấu trúc). |
| `browser_take_screenshot` | `browser` | Chụp ảnh màn hình viewport trang web phục vụ kiểm tra hoặc thị giác AI. |
| `browser_watch` | `browser` | Chụp toàn bộ cửa sổ (bao gồm tab bar, address bar) dạng live JPEG. |
| `browser_click` | `selector`, `browser` | Click phần tử theo selector CSS (tự cuộn vào tầm nhìn và chờ clickable). |
| `browser_click_at` | `x`, `y`, `hold_seconds`, `browser` | Di chuyển chuột thật đến tọa độ x, y và click (hoặc ấn giữ slider/captcha). |
| `browser_type` | `selector`, `text`, `browser` | Nhập văn bản theo nhịp gõ phím người thật, tự xóa nội dung cũ. |
| `browser_select_option` | `selector`, `value`, `browser` | Chọn mục trong dropdown `<select>` theo nhãn hoặc giá trị. |
| `browser_press_key` | `key`, `browser` | Bấm phím chức năng (`"Enter"`, `"Tab"`, `"Escape"`, `"PageDown"`, `"ArrowDown"`). |
| `browser_upload_files` | `selector`, `paths`, `browser` | Tải file lên `<input type=file>`. Đường dẫn phải nằm trong thư mục được cấp phép (`INVISIBLE_MCP_UPLOAD_DIRS`). |
| `browser_evaluate` | `expression`, `browser` | Đọc biến môi trường, state hoặc thuộc tính DOM qua JavaScript an toàn. |

---

## 3. Các Tình Huống Nâng Cao & Giải Pháp (Recipes)

### A. Vượt Cookie Banner & Hộp thoại Consent
- Khi vào trang mới, cookie banner thường che phủ toàn bộ viewport khiến các click bên dưới bị chặn (timeout).
- **Quy trình chuẩn**:
  1. Gọi `browser_snapshot`.
  2. Tìm nút "Accept All", "Agree", "Đồng ý", "Allow all".
  3. Gọi `browser_click` với selector của nút đó trước khi tương tác với nội dung chính.

### B. Vượt Cloudflare Turnstile & Checkbox Challenge
- Cloudflare Turnstile thường render trong iframe với checkbox "Verify you are human".
- **Quy trình chuẩn**:
  1. Chụp ảnh `browser_take_screenshot` hoặc lấy tọa độ checkbox từ `browser_snapshot`.
  2. Dùng `browser_click_at(x, y)` để di chuyển chuột thật vào tâm checkbox.
  3. Chờ 3-5 giây cho challenge hoàn tất (Firefox 41 engine tự động xử lý redirect `location.replace`).
  4. Kiểm tra lại URL bằng `browser_list` hoặc `browser_status`.

### C. Xử lý Cuộn Trang (Infinite Scroll / Lazy Load)
- Với các trang dữ liệu dài (Facebook feeds, Twitter, Shopee, tin tức):
  - **Cách 1 (Bàn phím)**: Gọi `browser_press_key("PageDown")` hoặc `browser_press_key("ArrowDown")` nhiều lần, cách nhau 1-2 giây.
  - **Cách 2 (Đọc qua evaluate)**: Để cuộn đến vị trí mong muốn:
    ```javascript
    window.scrollBy({ top: 800, behavior: 'smooth' });
    ```
    (Lưu ý: Scroll là hành vi đọc/hiển thị an toàn, không kích hoạt cơ chế chống can thiệp DOM).

### D. Tải tệp lên biểu mẫu (File Uploads)
- Các file đính kèm phải được đặt trong thư mục:
  `C:\Users\Admin\.gemini\antigravity\scratch`
- Gọi `browser_upload_files(selector="input[type=file]", paths=["C:\\Users\\Admin\\.gemini\\antigravity\\scratch\\tailieu.pdf"])`.

### E. Lưu Trữ Session & Profile Cố Định
- Profile chính được tự động lưu tại:
  `C:\Users\Admin\.gemini\antigravity\profiles\stealthfox`
- Toàn bộ Cookies, LocalStorage và phiên đăng nhập sẽ được duy trì xuyên suốt qua các lần khởi động.
