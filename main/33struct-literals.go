package main

/*
Struct Literals
A struct literal denotes a newly allocated struct value by listing the values of its fields.

You can list just a subset of fields by using the Name: syntax. (And the order of named fields is irrelevant.)

The special prefix & returns a pointer to the struct value.
 */
import "fmt"

type Vertex33 struct {
	X, Y int
}

var (
	v1 = Vertex33{1, 2}  // has type Vertex
	v2 = Vertex33{X: 1}  // Y:0 is implicit
	v3 = Vertex33{}      // X:0 and Y:0
	p  = &Vertex33{1, 2} // has type *Vertex
)

func main() {
	fmt.Println(v1, p, v2, v3)
}
