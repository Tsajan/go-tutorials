package tutorials

import "fmt"

func addtwonumbers(a int, b int) int {
	return (a + b)
}

func FnAdd() {
	x := addtwonumbers(10, 20)
	fmt.Println(x)
}
