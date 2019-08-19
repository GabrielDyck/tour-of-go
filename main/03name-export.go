package main

/*
Nombres exportados

Una vez importado un paquete, puedes referirte a los nombres que exporta.

En Go, un nombre es exportado si su primera letra es mayúscula.

Por ejemplo, suponiendo un paquete ficticio, Foo se exportaría, y también FOO, pero foo no.

Intenta ejecutar el código. Luego renombra math.pi a math.Pi e intenta de nuevo.
 */
import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Pi)
}