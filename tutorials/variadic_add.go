package tutorials

import "fmt"

func variadic_add(nums ...int) (z int) {
	for _, num := range nums {
		z += num
	}
	return
}

func VariadicAddDemo() {
	x := variadic_add(1, 2, 3, 4, 5)
	fmt.Println("Sum of first 5 numbers: ", x)
}
