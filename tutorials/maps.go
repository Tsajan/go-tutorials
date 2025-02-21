package tutorials

import "fmt"

func MapsDemo() {
	fmt.Println("Simple demonstration of map objects in Golang")

	var myMap map[int]string // declaring a map variable that maps keys of type into to values of type string

	// myMap[1] = "one" // funny thing here is that this will give compilation error, because just the previous statement is only declaration and not initialization
	// if map is initialized using make keyword; such as `var myMap = make(map[int]string)` allows initialization later or else using shot-hand declaration
	var myMap2 = make(map[int]string)
	myMap2[1] = "one"

	// short hand declaration of map
	myMap3 := map[string]int{"one": 1, "two": 2} // allows declaration as well as initlialization
	fmt.Println("myMap3 is: ", myMap3)

	// can append additional element to the map?
	myMap3["three"] = 3
	fmt.Println("Updated myMap3 is: ", myMap3)

	// empty map can be compared to nil keyword
	if myMap == nil {
		fmt.Println("Variable myMap is empty")
	} else {
		fmt.Println("myMap is not empty")
	}

	if myMap2 == nil {
		fmt.Println("Variable myMap2 is empty")
	} else {
		fmt.Println("myMap2 is not empty")
		fmt.Println(myMap2)
	}

	// iterating over a map using range
	// iteration of elements in map is randomized. so the same order of items will not be printed or iterated each time
	for k, v := range myMap3 {
		fmt.Printf("Key is %s and value is %d\n", k, v)
	}

	// checking if a values exists within a map; using an additional variable called exists that returns a boolean value
	val, exists := myMap3["one"]
	if exists == true {
		fmt.Println("Key exists and its values is: ", val)
	} else {
		fmt.Println("Key doesn't exist, ", exists)
	}

	// updating or adding kyes to an initialized map can be done using assignment operator
	myMap3["number"] = 123
	fmt.Println("Number added: ", myMap3)

	// to delete a key in a map using the delete keyword
	delete(myMap3, "number")
	fmt.Println("myMap3 after number key deletion is: ", myMap3)

	// maps are pass-by-reference; so if a map is copied into another map and changed; changes are reflected in the underlying original map as well
	copiedMap := myMap3
	copiedMap["one"] = 11
	fmt.Println("Original map is: ", myMap3)
}
