package main

import "fmt"
/**
当两个或多个连续的函数命名参数是同一类型，除了最后一个类型之外，其他都可以省略
 */
func add2(x, y int) int {
	return x + y
}
func main()  {
	fmt.Println(add2(4,5))
}
