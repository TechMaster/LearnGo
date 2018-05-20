Glide là công cụ quản lý package trong Go tương tự như npm trong Node.js

Để sử dụng code mới nhất trong thư mục trong máy tính thay vì phải pull code từ Github, Gitlab.

Gõ lệnh
```
glide mirror set github.com/TechMaster/funpackage file://~/gogit/TechMaster/FunPackage
```
Chú ý ở đây ~/gogit là symbolic link trỏ đến $GOPATH/src/github.com

Đồng thời xoá thư mục code tương ứng trong thư mục vendor. Trong ví dụ này tôi import
[https://github.com/TechMaster/funpackage](https://github.com/TechMaster/funpackage)

Do đó tôi xoá vendor/github.com/TechMaster/funpackage

Lệnh glide mirror sẽ tạo ra file mirrors.yml ở ~/.glide. Nội dung file sẽ như sau
```
repos:
- original: github.com/TechMaster/funpackage
  repo: file://~/gogit/TechMaster/FunPackage
```