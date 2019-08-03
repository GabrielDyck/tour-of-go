package main

/*
Funciones

Una función puede tener cero o más argumentos.

En este ejemplo, add tiene dos parámetros de tipo int.

Fíjate que el tipo va después del nombre de la variable.

(Si quieres más información sobre porqué los tipos tienen la apariencia que tienen, consulta el artículo sobre la sintaxis de declaraciones en Go.) https://blog.golang.org/gos-declaration-syntax
 */
import "fmt"

func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}