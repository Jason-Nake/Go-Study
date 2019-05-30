package main

import "fmt"

type Vertex2 struct {
	X int
	Y int
}

/**
结构体指针
结构体字段可以通过结构体指针来访问。

通过指针间接的访问是透明的。
*/
func main() {
	v := Vertex2{1, 2}
	p := &v
	p.X = 99
	fmt.Println(v)
}
