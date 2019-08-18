package main
/*
Nil interface values
A nil interface value holds neither value nor concrete type.

Calling a method on a nil interface is a run-time error because there is no type inside the interface tuple to indicate which concrete method to call.
 */
import "fmt"

type I64 interface {
	M64()
}

func main() {
	var i I64
	describe64(i)
	i.M64()
}

func describe64(i I64) {
	fmt.Printf("(%v, %T)\n", i, i)
}
