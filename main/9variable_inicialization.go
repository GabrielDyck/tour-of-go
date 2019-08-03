package main


/*
Inicialización de variables

Una declaración de variables puede incluir valores iniciales, uno por variable.

Si éstos estan presentes, se puede omitir el tipo y la variable tendrá el tipo del valor inicial utilizado.
 */
import "fmt"

var xx, yy, zz int = 1, 2, 3
var cc, python2, java2 = true, false, "no!"

func main() {
	fmt.Println(xx, yy, zz, cc, python2, java2)
}