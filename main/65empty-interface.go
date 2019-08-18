package main
/*
The empty interface
The interface type that specifies zero methods is known as the empty interface:

interface{}
An empty interface may hold values of any type. (Every type implements at least zero methods.)

Empty interfaces are used by code that handles values of unknown type. For example, fmt.Print takes any number of arguments of type interface{}.
 */
import "fmt"

func main() {
	var i interface{}
	describe65(i)

	i = 42
	describe65(i)

	i = "hello"
	describe65(i)
}

func describe65(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
