package main

import (
	"fmt"
	"runtime"
)

var (
	i int     // 32 (x86) or 64 (x64) integer; depends on the micro. arch.
	u uint    // 32 (x86) or 64 (x64) UNSIGNED integer; depends on the micro. arch.
	r rune    // 32 bit integer
	b byte    // 8 bit integer
	d float64 // Double-precision floating-point decimal
	f float32 // floating-point decimal
)

/* RANGES
uint8  : 0 to 255
uint16 : 0 to 65535
uint32 : 0 to 4294967295
uint64 : 0 to 18446744073709551615
int8   : -128 to 127
int16  : -32768 to 32767
int32  : -2147483648 to 2147483647
int64  : -9223372036854775808 to 9223372036854775807
*/

func main() {

	fmt.Println("My micro. architecture is", runtime.GOARCH)
	fmt.Println("default values:", i, u, r, b, d, f)

	i = 8008
	u = 255
	r = rune(-i)
	b = 255
	d = 80.08
	f = float32(-d)

	fmt.Println(i, u, r, b, d, f)

	// Conversion
	r = int32(f)
	fmt.Println("Float converted to Rune", r)

	d = float64(i)
	fmt.Println("Int converted to Float", d)

}
