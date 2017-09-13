// +build OMIT

package main

import "fmt"

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}  // от тип Vertex
	v2 = Vertex{X: 1}  // Y:0 е създадено по подразбиране
	v3 = Vertex{}      // X:0 и Y:0 по подразбиране
	p  = &Vertex{1, 2} // от тип *Vertex
)

func main() {
	fmt.Println(v1, p, v2, v3)
}
