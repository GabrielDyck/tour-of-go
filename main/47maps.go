package main

/*
Maps
A map maps keys to values.

The zero value of a map is nil. A nil map has no keys, nor can keys be added.

The make function returns a map of the given type, initialized and ready for use.
 */
import "fmt"

type Vertex47 struct {
	Lat, Long float64
}

var m map[string]Vertex47

func main() {
	m = make(map[string]Vertex47)
	m["Bell Labs"] = Vertex47{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
}
