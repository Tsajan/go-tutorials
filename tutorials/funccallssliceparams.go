package tutorials

import "fmt"

func FnCallsWithSlice() {
	// declare a slice
	var ogSlice = []int{10, 20, 30, 40, 50}

	fmt.Println("Original slice is: ", ogSlice)

	updateSlice(ogSlice)

	fmt.Println("Original slice after update call is: ", ogSlice)
}

// because slices are passed by references, any changes to the slice will update the original slice
func updateSlice(mySlice []int) {
	fmt.Println("Original slice within function is: ", mySlice)

	// update slice value
	mySlice[2] = 300

}
