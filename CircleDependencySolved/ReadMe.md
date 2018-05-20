Đây là hướng sửa chữa lỗi circular dependency ở dự án
[CircleDependency](https://github.com/TechMaster/LearnGo/tree/master/CircleDependency)

Tạo ra một interface 
```
     +--------------------------+
     |          share           |
     +--------------------------+
         ^                ^
         |                |
 +-------------+     +------------+
 |   apackage  |     |  bpackage  |
 +-------------+     +------------+
```