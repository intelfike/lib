package main

import (
	"fmt"

	"github.com/intelfike/lib/ringbuf"
)

func main() {

	rb.Clear()
	rb.Write(23)
	rb.Write(10)

	fmt.Println(rb)

	rb2, _ := ringbuf.New(make([]int, 4))
	rb2.Write(999)
	rb2.Write(888)
	rb2.Write(777)
	rb2.Write(666)
	rb2.Write(555)
	fmt.Println(rb2)
}
