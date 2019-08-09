package main

import "fmt"
/*
Struct Fields
Struct fields are accessed using a dot.
 */
type Vertex31 struct {
	X int
	Y int
}

func main() {
	v := Vertex31{1, 2}
	v.X = 4
	fmt.Println(v.X)
}
