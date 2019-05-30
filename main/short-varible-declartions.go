package main

import "fmt"

/**
在函数中，':='简洁赋值语句使用在明确类型的地方，可以代替使用var语句
函数外的语句都必须以关键字开始（'var','func'等等），':='结构不能使用在函数外
 **/
func main()  {
	var i, j = 1, 2
	k := 3
	c, python, java := true, false, true
	fmt.Println(i, j, k, c, python, java)
}
