Ví dụ này demo một workflow khá phổ biến:
1. Web server nhận file upload từ trình duyệt
2. Web server lưu file vào thư mục sau đó xử lý. Ví dụ với file mp4 thì sẽ encoding ra định dạng khác nhau.
3. Quá trình encode là một tác vụ tốn CPU, chạy trong thời gian.

**Công nghệ tôi học hỏi ở đây:**
1. Web framework dùng iris hỗ trợ sẵn web templatesocket, file upload và web
2. Phần xử lý tác vụ dài (long run task) trước mắt sẽ dùng go routine, và channel. 
3. Web socket để thông báo tiến trình chạy long run task

**Chạy thử**
1. ```go run *.go```
