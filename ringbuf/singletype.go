package ringbuf

import (
	"errors"
	"fmt"
	"reflect"
)

type RingBuf struct {
	val   reflect.Value
	index int
	count int
}

var _ fmt.Stringer = RingBuf{}

// Exsample:
// New(make([]int, 22))
func New(i interface{}) (*RingBuf, error) {
	srb := new(RingBuf)
	srb.val = reflect.ValueOf(i)

	if srb.val.Kind() != reflect.Slice {
		return nil, errors.New("Argument of NewSingleType({}interface) is Slice only: Exsample -> make([]int, 2)")
	}
	return srb, nil
}

func MustNew(i interface{}) *RingBuf {
	ring, err := New(i)
	if err != nil {
		panic(err)
	}
	return ring
}

// i[0] = the most old
// i[len(i) - 1] = the most new
// write to ring buffer.
func (s *RingBuf) Write(i interface{}) {
	if s.Cap() == 0 {
		return
	}
	s.val.Index(s.index).Set(reflect.ValueOf(i))
	s.index++
	s.index %= s.Cap()
	if s.count < s.Cap() {
		s.count++
	}
}

func (s *RingBuf) Clear() {
	s.count = 0
}

// Get older element.
// 0 is the most old.
func (s RingBuf) IndexOld(i int) interface{} {
	if s.Len() == 0 {
		return nil
	}
	index := LoopMod(i+s.index, s.Len())
	return s.val.Index(index)
}

// Get newer element.
// 0 is the most new.
func (s RingBuf) IndexNew(i int) interface{} {
	return s.IndexOld(-1 - i)
}

// Get slice in interface{}.
// Type assert is yourself.
func (s RingBuf) Get() interface{} {
	return s.getValue().Interface()
}

// reflect.Value
func (s RingBuf) getValue() reflect.Value {
	start := s.val.Slice(s.index, s.Cap())
	end := s.val.Slice(0, s.index)
	return reflect.AppendSlice(start, end).Slice(s.Cap()-s.Len(), s.Cap())
}

// Get() to string.
func (s RingBuf) String() string {
	return fmt.Sprint(s.Get())
}

// writed length.
func (s RingBuf) Len() int {
	return s.count
}

// The return value is the slice size decided by New() Argument.
func (s RingBuf) Cap() int {
	return s.val.Len()
}

// -1 % 3 != -1:
// -1 % 3 == 2
func LoopMod(i, j int) int {
	mod := i % j
	if mod < 0 {
		mod += j
	}
	return mod
}
