package main
/*
Methods
Go does not have classes. However, you can define methods on types.

A method is a function with a special receiver argument.

The receiver appears in its own argument list between the func keyword and the method name.

In this example, the Abs method has a receiver of type Vertex named v.
 */
import (
	"fmt"
	"math"
)

type Vertex52 struct {
	X, Y float64
}

func (v Vertex52) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex52{3, 4}
	fmt.Println(v.Abs())
}
