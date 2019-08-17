package main

/*
Interface values with nil underlying values
If the concrete value inside the interface itself is nil, the method will be called with a nil receiver.

In some languages this would trigger a null pointer exception, but in Go it is common to write methods that gracefully handle being called with a nil receiver (as with the method M in this example.)

Note that an interface value that holds a nil concrete value is itself non-nil.
 */
import "fmt"

type I63 interface {
	M63()
}

type T63 struct {
	S string
}

func (t *T63) M63() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func main() {
	var i I63

	var t *T63
	i = t
	describe63(i)
	i.M63()

	i = &T63{"hello"}
	describe63(i)
	i.M63()
}

func describe63(i I63) {
	fmt.Printf("(%v, %T)\n", i, i)
}
