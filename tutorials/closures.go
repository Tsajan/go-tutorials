package tutorials

import "fmt"

func counter() func() int {
	count := 0
	return func() int {
		count += 1
		return count
	}
}

func ClosureDemo() {
	fmt.Println("Initial value of counter is: 0")

	increment := counter()
	fmt.Println("Value of counter is: ", increment())             // increments the value by 1, to 1 from 0, which is retained
	fmt.Println("Value of counter incremented to: ", increment()) // increments the value by 1; to 2, from 1 which was retained
}
