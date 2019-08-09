package main

/*
Más funciones

Cuando dos parámetros o más tienen el mismo tipo, puedes omitirlo en todos excepto el último.

En este ejemplo, hemos acortado

x int, y int
a

x, y int
 */
import "fmt"

func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}