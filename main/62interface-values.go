package main

/*
Interface values
Under the hood, interface values can be thought of as a tuple of a value and a concrete type:

(value, type)
An interface value holds a value of a specific underlying concrete type.

Calling a method on an interface value executes the method of the same name on its underlying type.
 */
import (
	"fmt"
	"math"
)

type I62 interface {
	M62()
}

type T62 struct {
	S string
}

func (t *T62) M62() {
	fmt.Println(t.S)
}

type F62 float64

func (f F62) M62() {
	fmt.Println(f)
}

func main() {
	var i I62

	i = &T62{"Hello"}
	describe(i)
	i.M62()

	i = F62(math.Pi)
	describe(i)
	i.M62()
}

func describe(i I62) {
	fmt.Printf("(%v, %T)\n", i, i)
}
