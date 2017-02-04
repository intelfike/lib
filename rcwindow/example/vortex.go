package main

import (
	"fmt"
	"math/cmplx"
	"time"

	"github.com/intelfike/lib/rcwindow"
)

func main() {
	rc := rcwindow.NewWindow(1, 1, 200)
	rc.SafeConfig(func() { rc.DotSize = 3 })
	c := cmplx.Pow(1i, 1.0/100.0)
	v := 1i
	rc.RedrawTick(1 << 24)
	for n := 0; n <= 2000; n++ {
		if n == 1000 {
			fmt.Println("clear")
		}
		rc.Dot(real(v), imag(v))
		v *= c / 1.001
		time.Sleep(1)
	}
	rc.Wait()
}
