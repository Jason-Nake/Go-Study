package main

import "fmt"

/**
对 slice 切片
slice 可以重新切片，创建一个新的 slice 值指向相同的数组。

表达式

s[lo:hi]
表示从 lo 到 hi-1 的 slice 元素，含两端。因此

s[lo:lo]
是空的，而

s[lo:lo+1]
有一个元素。
*/
func main() {
	p := []int{2, 3, 4, 5, 7, 11, 13}
	fmt.Println("P ==", p)
	fmt.Println("P[1:4] ==", p[0:4])

	//省略下标代表从0开始
	fmt.Println("P[:3] ==", p[:3])

	//省略上标代表到len(s)结束
	fmt.Println("P[4:] ==", p[4:])
}
