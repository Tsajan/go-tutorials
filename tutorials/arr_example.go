package tutorials

import "fmt"

func ArrayExample() {
	// declaring an array
	var intArr [5]int
	intArr = [5]int{1, 2, 3, 4, 5}

	fmt.Println(intArr)

	// using short-hand declaration to declare and initialze an array
	arr2 := [10]string{"sajan", "rajan", "shristi", "saju", "ashamaya", "radheshyam", "rojni", "ryan", "raman", "joel"}
	fmt.Println(arr2)

	// go allows partial assignment of elements in an array
	arr3 := [10]int{1, 2, 3} // partial assignment where only the first three elements are assigned values and the rest is initialized to 0
	fmt.Println(arr3)

	// go allows somehow dynamic allocation of an array size; during declaration no need to specify its length, but during initialization the size of array is determined; use ... ellipses in array declaration to do so
	arr4 := [...]int{8, 9, 10, 11}
	fmt.Println(arr4)

	// go allows initializing array elements at specific indices
	arr5 := [5]int{0: 1, 4: 1} // this should assign values at indices 0 and 4 to values 1; whereas others are assigned default values of 0
	fmt.Println(arr5)

	// also you can assign values for the indices that were initialized during declaration
	arr5[2] = 2
	fmt.Println(arr5)

	// can the values intialized earlier be updated to reflect a new value
	arr5[2] = 1000
	fmt.Println(arr5) // apparently, yes!
}
