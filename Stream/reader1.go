package main

import (
	"fmt"
	"io"
)

// alphaReader is a simple implementation of an io.Reader
// that streams only alpha chars from its string source.
type alphaReader struct {
	src string
	cur int
}

func newAlphaReader(src string) *alphaReader {
	return &alphaReader{src: src}
}

func alpha(r byte) byte {
	if (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z') {
		return r
	}
	return 0
}

func (a *alphaReader) Read(p []byte) (int, error) {
	if a.cur >= len(a.src) {
		return 0, io.EOF
	}

	x := len(a.src) - a.cur  //x là độ dài từ biến cur đến kết thúc chuỗi src trong a
	n, bound := 0, 0
	if x >= len(p) { //Nếu x lớn hơn kích thước chuỗi p sẽ đọc ra lấy bound bằng kích thước p
		bound = len(p)
	} else if x <= len(p) {
		bound = x
	}

	buf := make([]byte, bound)  //Tạo mảng buf
	for n < bound {  //lặp từ n đến bound
		if char := alpha(a.src[a.cur]); char != 0 {  //Kiểm tra ký tự ở vị trí cur trong a có là alphabet không
			buf[n] = char //nếu đúng copy vào
		}
		n++
		a.cur++
	}
	copy(p, buf)  //copy từ buf vào p
	return n, nil //trả về giá n số ký tự thực sự đạt yêu
}

func main() {
	reader := newAlphaReader("Hello! It's 9am, where is the sun?")
	p := make([]byte, 4)
	for {
		n, err := reader.Read(p)
		if err == io.EOF {
			break
		}
		fmt.Print(string(p[:n]))
	}
	// or use io.Copy
	// io.Copy(os.Stdout, reader)
	fmt.Println()
}
