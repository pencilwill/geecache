package geecache

// A ByteView holds an immutable view of bytes. 一个 ByteView 持有一个不可改变的字节视图。
type ByteView struct {
	b []byte
}

// Len returns the view's length Len 返回视图的长度
func (v ByteView) Len() int {
	return len(v.b)
}

// ByteSlice returns a copy of the data as a byte slice. ByteSlice 以字节片的形式返回数据的副本。
func (v ByteView) ByteSlice() []byte {
	return cloneBytes(v.b)
}

// String returns the data as a string, making a copy if necessary. String 将数据作为字符串返回，必要时进行复制。
func (v ByteView) String() string {
	return string(v.b)
}

func cloneBytes(b []byte) []byte {
	c := make([]byte, len(b))
	copy(c, b)
	return c
}
