package tutorials

import "fmt"

// declare a custom type; say a struct for MountainBike
type MountainBike struct {
    model string
    year int
    price int
}

// binding this method to the MountainBike data-type; much like methods bounds to a certain class type in OOP paradigm
// methods cannot be declared within functions in golang!
func (obj MountainBike) getPriceOfModel() {
    md := obj.model 
    cost := obj.price
    fmt.Printf("%s costs %d", md, cost)
}

// if we want to perform an update on the underlying data type using methods; we ought to pass a reference to the object i.e. pointers
// to reflect changes, we need to pass pointer arguments to the receiverType
// if we don't use * operator on the MountainBike object, the changes wont be reflected
func (obj *MountainBike) updateModelName(newModelName string) {
    obj.model = newModelName
}

// not using * on underlying object of method receiver doesn't reflect write changes; example
func (obj MountainBike) updatePrice(newPrice int) {
    obj.price = newPrice
}

func MethodsDemo() {
    fmt.Println("This the methods demo file!")

    // instantiate an object of the struct data type
    mb1 := MountainBike{ "Giant Anthem", 2021, 150000 }
    mb1.getPriceOfModel()

    // instantiate a new MountainBike object
    mb2 := MountainBike { "Trek", 2024, 220000 }
    fmt.Println("mb2 is: ", mb2)
        
    mb2.updateModelName("Trek Excalibur")
    fmt.Println("Changes reflected normally without using pointers: ", mb2)
    
    // multiple changes updated using pointers will reflect to the underlying object 
    point_mb2 := &mb2 // a pointer object to the underlying datatype
    point_mb2.updateModelName("Trek Excalibur II")
    fmt.Println("Changes reflected using pointers; mb2 is: ", mb2)

    mb2.updatePrice(999999) // this change is not reflected because in the corresponding method, the receiver argument is not being passed as pointers
    fmt.Println("Updated price is not reflected: ", mb2)
}


