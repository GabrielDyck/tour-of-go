package main

/*
Resultados con nombre

En Go, además de poder nombrar los parámetros de una función, también se pueden nombrar los resultados.

Si se nombran los resultados y se utilizan en la función, la instrucción return devolverá sus valores en ese momento.
 */
import "fmt"

func split(sum int) (x, y int) {
	x = sum * 4/9
	y = sum - x
	return
}

func main() {
	fmt.Println(split(17))
}