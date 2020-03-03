/*“写一个带有如下函数签名的函数CountingWriter，传入一个io.Writer接口类型，返回一个把原来的Writer封装在里面的新的Writer类型和一个表示新的写入字节数的int64类型指针”*/
package counter

import "io"

//implement a new Writer that can count the number of bytes writen to the Writer when write
type countWriter struct {
	w      io.Writer
	writen int
}

func (cw *countWriter) Write(p []byte) (int, error) {
	n, err := cw.w.Write(p)
	cw.writen += n
	return n, err
}

//CountingWriter return the writer
func CountingWriter(w io.Writer) (io.Writer, *int64) {
	cw := countWriter{w: w}
	n := int64(cw.writen)
	return &cw, &n
}
