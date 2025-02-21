package tutorials

import "fmt"

func InfiniteForLoop() {
	// Just providing no boolean condition with just the for keyword results in an infinite loop
	for {
		fmt.Println("Printing again and again!")
	}
}
