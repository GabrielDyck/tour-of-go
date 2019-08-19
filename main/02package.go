package main


/*
Paquetes

Todo programa en Go está hecho de paquetes.

Los programas se empiezan a ejecutar en el paquete main.

Este programa utiliza los paquetes con las rutas "fmt" y "math".

La convención es que el nombre del paquete es el mismo que el último elemento de su ruta.
 */

/*
Importar paquetes

Este código agrupa los paquetes que importa en una única sentencia "factorizada". Aunque también puedes importar los paquetes con sentencias separadas:

import "fmt"
import "math"
pero es típico usar la forma "factorizada" para simplificar.
 */
//Sentencia factorizada para simplificar
import (
	"fmt"
	"math"
)
// or
// import "fmt"
// import "math"

func main() {
	fmt.Println("Happy", math.Pi, "Day")
}