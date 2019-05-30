package main

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe bool = false
	MaxInt uint64 = 1<<61 - 1
	z complex128= cmplx.Sqrt(-5 + 12i)

)
/**
golang的基本类型
bool
string
int int8 int16 int32 int64
uint uint8 uint 16 uint32 uint64 uintptr
byte //unit8别名
rune //int32别名 //代表一个unicode码
float float32 float64
complex64 complex128
 */
func main()  {
	const f  =  "%T(%v)\v"
	fmt.Printf(f, ToBe, ToBe)
	fmt.Printf(f, MaxInt, MaxInt)
	fmt.Printf(f, z, z)
}
