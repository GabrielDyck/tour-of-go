package main

/*

Map literals
Map literals are like struct literals, but the keys are required.
 */
import "fmt"

type Vertex48 struct {
	Lat, Long float64
}

var m48 = map[string]Vertex48{
	"Bell Labs": Vertex48{
		40.68433, -74.39967,
	},
	"Google": Vertex48{
		37.42202, -122.08408,
	},
}

func main() {
	fmt.Println(m48)
}
