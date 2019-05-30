package main

import "fmt"

/**
定义两个数相加的函数，注意变量类型在变量名之后
 */
func add(x int, y int)	int {
	return x + y
}
func main()  {
	fmt.Println(add(4,5))
}
