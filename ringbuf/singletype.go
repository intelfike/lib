package ringbuf

import (
	"fmt"
	"reflect"
)

type singleRingBuf struct {
	Val   reflect.Value
	Index int
}

// Exsample
// NewSingleType(make([]int, 22))
func NewSingleType(i interface{}) singleRingBuf {
	mrb := *new(singleRingBuf)
	mrb.Val = reflect.ValueOf(i)

	if mrb.Val.Kind() != reflect.Slice {
		panic("Argument of NewSingleType({}interface) is Slice only: Exsample -> make([]int, 2)")
	}
	return mrb
}

// write to ring buffer
func (m *singleRingBuf) Write(i interface{}) {
	m.Val.Index(m.Index).Set(reflect.ValueOf(i))
	m.Index++
	m.Index %= m.Val.Len()
}

// get slice in interface{}
// type assert is yourself
func (m *singleRingBuf) Get() interface{} {
	start := m.Val.Slice(m.Index, m.Val.Len())
	end := m.Val.Slice(0, m.Index)
	return reflect.AppendSlice(start, end).Interface()
}

// Get() to string
func (m singleRingBuf) String() string {
	return fmt.Sprint(m.Get())
}
