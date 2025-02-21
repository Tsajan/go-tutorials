package tutorials

import "fmt"

func VarDeclare() {
	a := 10

	// a := 20 // redeclaring variable with just that single variable gives error

	a, b := 20, 40 // but if we re-declare the same variable with another new variable; it works
	fmt.Println(a, b)
}
