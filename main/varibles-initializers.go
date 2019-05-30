package main

import "fmt"

var i, j  int =  1, 2

/**
变量定义包含初始值，每个变量对应一个
如果初始化是使用表达式，可以省略类型；变量从初始值中获得类型
 */
func main()  {
	var c, python, java  = true, false, true
	fmt.Println(i, j, c, python, java)
}
