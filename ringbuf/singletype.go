package ringbuf

import (
	"fmt"
	"reflect"
)

type SingleRingBuf struct {
	Val   reflect.Value
	Index int
}

// Exsample:
// NewSingleType(make([]int, 22))
func NewSingleType(i interface{}) SingleRingBuf {
	srb := *new(SingleRingBuf)
	srb.Val = reflect.ValueOf(i)

	if srb.Val.Kind() != reflect.Slice {
		panic("Argument of NewSingleType({}interface) is Slice only: Exsample -> make([]int, 2)")
	}
	return srb
}

// write to ring buffer.
func (s *SingleRingBuf) Write(i interface{}) {
	s.Val.Index(s.Index).Set(reflect.ValueOf(i))
	s.Index++
	s.Index %= s.Val.Len()
}

// Get older element.
// 0 is the most old.
func (s *SingleRingBuf) IndexOld(i int) interface{} {
	index := s.LoopModLen(i + s.Index)
	return s.Val.Index(index)
}

// Get newer element.
// 0 is the most new.
func (s *SingleRingBuf) IndexNew(i int) interface{} {
	return s.IndexOld(-1 - i)
}

// Get slice in interface{}.
// Type assert is yourself.
func (s *SingleRingBuf) Get() interface{} {
	start := s.Val.Slice(s.Index, s.Val.Len())
	end := s.Val.Slice(0, s.Index)
	return reflect.AppendSlice(start, end).Interface()
}

// Get() to string.
func (s SingleRingBuf) String() string {
	return fmt.Sprint(s.Get())
}

// -1 % 3 != -1
// -1 % 3 == 2
func LoopMod(i, j int) int {
	mod := i % j
	if mod < 0 {
		mod += j
	}
	return mod
}

func (s SingleRingBuf) LoopModLen(i int) int {
	return LoopMod(i, s.Val.Len())
}

func (s SingleRingBuf) Len() int {
	return s.Val.Len()
}
