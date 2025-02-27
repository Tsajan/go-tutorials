package tutorials

import "fmt"

type Bicycle struct {
	Model       string
	ReleaseYear int
	price       int // price is private since the field is declared with a small-case letter
}

func StructExportsDemo() {
	// initialize a Bicyle structure within local fn; should allow access to all elements (even if some elements are initialized with small case because it is local)
	b1 := Bicycle{"Giant Revel", 2015, 38000}

	fmt.Println("Bicycle b1 is: ", b1)
}
