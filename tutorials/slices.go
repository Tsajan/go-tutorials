package tutorials

import "fmt"

func SliceDemo() {
	var mySlice []int
	mySlice = append(mySlice, 2)

	fmt.Println("Slice: ", mySlice)

	slice2 := []int{1, 2, 3, 4, 5} // short-hand declaration of slice
	// slice2[5] = 6 // throws a panic when assigning a value currently out of index; needs to use append
	fmt.Println("Slice 2 is: ", slice2)

	// creating a slice out of slice
	slice3 := slice2[:3] // upto index 2 i.e. the first 3 elements
	fmt.Println("Slice of slice: ", slice3)

	// creating a null value slice
	var slice4 []int
	fmt.Println("A null slice: ", slice4)

	// golang uses nil keyword to check for zero value slices
	fmt.Println("Prints true if slice4 == nil : ", (slice4 == nil))

	// wtf is capacity of a slice
	fmt.Println("Capacity of slice3 is: ", cap(slice3))
}
