package tutorials

import "fmt"

// globally exported function that would be imported into main package
// also can be altered within the nested function calls
var ExportedVar int

// nested variables (including functions; since functions are first class citizens) would not be exported when the package is exported
func AnonymousFnCall() {

	// golang allows function calls within fn calls; higher order functions;
	// anonymous functions; a function has no name! (GoT reference if you got it!)
	square := func(num int) int {
		return (num * num)
	}

	// though this is Exportable as signified by the capital letter; nested functions do not get exported
	ExportableSquare := func(num int) int {
		return (num * num)
	}

	// one way to circumnavigate this problem would be to store the operations of change into a global exported variable and update its values within the nested function
	ExportedVar = ExportableSquare(11)

	fmt.Println("ExportedVar is: ", ExportedVar)
	fmt.Printf("ExportableSquare of number %d is %d\n", 9, ExportableSquare(9))
	fmt.Printf("Square of number %d is %d\n", 4, square(4))
}
