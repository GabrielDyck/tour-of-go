package main
/*
Methods are functions
Remember: a method is just a function with a receiver argument.

Here's Abs written as a regular function with no change in functionality.
 */
import (
	"fmt"
	"math"
)

type Vertex53 struct {
	X, Y float64
}

func Abs(v Vertex53) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex53{3, 4}
	fmt.Println(Abs(v))
}
