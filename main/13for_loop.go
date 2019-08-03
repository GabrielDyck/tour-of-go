package main

/*
For

Para hacer bucles, Go solo tiene el for.

El for básico tiene el mismo aspecto que el de C o Java, solo que los paréntesis () no hacen falta (no son ni opcionales), y las llaves {} son obligatorias.
 */
import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}