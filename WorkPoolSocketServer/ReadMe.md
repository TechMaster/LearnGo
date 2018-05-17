Một số lưu ý khi lập trình jQuery đẩy dữ liệu lên iRIS

1. Các key trong JSON phải viết hoa ký tự đầu tiên
2. Phải sử dụng JSON.stringify để convert đối tượng ra string
```
$.ajax({
    url: '/task',
    type: 'post',
    dataType: 'json',
    success: successCallBack,
    data: JSON.stringify(task)
});
```