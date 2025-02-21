package tutorials

import "fmt"

func StructsDemo() {
	// custom structure declaration of type bicycle
	type bicycle struct {
		model string
		year  int
		price int
	}

	// initializing the struct object; values are passed in the same order in which they are declared
	b1 := bicycle{"Giant Rincon Disc", 2015, 33500}
	fmt.Println("Struct object b1 is: ", b1)

	// initialization with field names overcomes such need for order of declaration
	b2 := bicycle{model: "Giant Anthem", price: 150000, year: 2022}
	fmt.Println("Struct object b2 is: ", b2)

	// zero value struct or empty struct
	var b3 bicycle
	fmt.Println("Bike b3 is empty: ", b3)

	// structs can also be initialized with partial assignment of the fields in such cases; the uninitialized fields hold empty values
	b4 := bicycle{model: "Trek Excalibur", year: 2021}
	fmt.Println("Bike b4 details: ", b4)

	// use the dot operator to access fields of a struct object
	model_name := b1.model
	fmt.Println("Model name is: ", model_name)

	// pointers to a struct is valid as well
	b1pointer := &b1
	fmt.Println("Bike b1 pointer is: ", b1pointer)

	// access the corresponding field from the pointer object as well
	fmt.Println("model is: ", b1pointer.model)

	b1pointer.model = "Updated model"
	fmt.Println("Bike b1 is updated? ", b1)
}
