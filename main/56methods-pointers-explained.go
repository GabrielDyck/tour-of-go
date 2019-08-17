package main
/*
Pointers and functions
Here we see the Abs and Scale methods rewritten as functions.

Again, try removing the * from line 16. Can you see why the behavior changes? What else did you need to change for the example to compile?

(If you're not sure, continue to the next page.)
 */
import (
	"fmt"
	"math"
)

type Vertex56 struct {
	X, Y float64
}

func Abs56(v Vertex56) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Scale56(v *Vertex56, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex56{3, 4}
	Scale56(&v, 10)
	fmt.Println(Abs56(v))
}
