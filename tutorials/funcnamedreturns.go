package tutorials

import "fmt"

func multiplyAndDivide(n1, n2 int) (multiplier int, divider int) {
	multiplier = n1 * n2
	divider = n1 / n2
	return // using named returns allows returns multiple return objects without having to specify them in the return statement
}

func FnCallWithNamedReturns() {
	num1, num2 := 40, 5
	fmt.Printf("Numbers are: %v and %v", num1, num2)

	mul, div := multiplyAndDivide(num1, num2)
	fmt.Println("Result: ", mul, div)
}
