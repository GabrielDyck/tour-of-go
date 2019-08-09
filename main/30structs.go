package main

/*
Structs
A struct is a collection of fields.
 */
import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Println(Vertex{1, 2})
}
