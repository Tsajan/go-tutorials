package tutorials

import "fmt"

func DefaultVariableValues() {
	var f float64
	var i int16
	var b bool
	var s string

	fmt.Printf("%v %v %v %q\n", f, i, b, s)
}
