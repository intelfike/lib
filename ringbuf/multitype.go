package ringbuf

import "fmt"

type multiRingBuf struct {
	Buf   []interface{}
	Index int
}

func NewMultiType(size int) multiRingBuf {
	return multiRingBuf{Buf: make([]interface{}, size)}
}
func (r *multiRingBuf) Write(itf interface{}) {
	r.Buf[r.Index] = itf
	r.Index++
	r.Index %= len(r.Buf)
}

// get interface slice
// type assert is yourself(very hard?)
func (r *multiRingBuf) Get() []interface{} {
	return append(r.Buf[r.Index:], r.Buf[:r.Index]...)
}
func (r multiRingBuf) String() string {
	return fmt.Sprint(r.Get())
}
