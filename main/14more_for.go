package main

/*
Más For

Como en C o Java, puedes dejar vacías las sentencias pre y post.
 */
import "fmt"

func main() {
	sum := 1
	for ; sum < 1000; {
		sum += sum
	}
	fmt.Println(sum)
}