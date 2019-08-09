package main

/*
Constantes
Las constantes se declaran como las variables, poniendo el prefijo const, que es una palabra reservada.

Las constantes pueden ser caracteres, cadenas de caracteres, Booleanos, o valores numéricos.
 */
import "fmt"

const Pi = 3.14

func main() {
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}