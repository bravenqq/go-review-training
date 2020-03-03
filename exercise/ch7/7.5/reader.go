/* “io包里面的LimitReader函数接收一个io.Reader接口类型的r和字节数n，并且返回另一个从r中读取字节但是当读完n个字节后就表示读到文件结束的Reader。实现这个LimitReader函数”*/

package reader

import (
	"io"
)

//LimitReader reade limit byte of n
type limitReader struct {
	reader io.Reader
	limit  int
	n      int
}

func (r *limitReader) Read(p []byte) (n int, err error) {
	if r.n+len(p) > r.limit {
		d := make([]byte, r.limit-r.n)
		n, err = r.reader.Read(d)
		copy(p, d)
	} else {
		n, err = r.reader.Read(p)
	}

	r.n += n
	if r.n == r.limit {
		err = io.EOF
	}
	return
}

func LimitReader(r io.Reader, n int64) io.Reader {
	lr := limitReader{reader: r, limit: int(n)}
	return &lr
}
