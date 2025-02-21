package tutorials

import "fmt"

func VariableReassignment() {
	var a int = 40
	fmt.Println("Value of a is", a)

	// re-assign the value of a
	a = 42
	fmt.Println("Updated value of a is: ", a)

	// print the length of a string
	str := "This is my string"
	l := len(str)
	fmt.Printf("%v has %v characters", str, l)
}
