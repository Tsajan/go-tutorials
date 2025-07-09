package tutorials

import "fmt"

// define an interface
type Speaker interface {
	MakeNoise() string // methods that needs to be implemented
}

// a Dog class
type Dog struct {
	Name string
}

// a Cat class
type Cat struct {
	Name string
}

// Dog class implementation of the MakeNoise interface
func (d Dog) MakeNoise() string {
	return "Raawr! I am a dog and my name is " + d.Name
}

// Cat class implementation of the MakeNoise interface
func (c Cat) MakeNoise() string {
	return "Meow! Meow! I am a cat and my name is " + c.Name
}

func introduce(s Speaker) {
	fmt.Println(s.MakeNoise())
}

func RunInterfaceExample() {
	growlithe := Dog{Name: "Growlithe"}
	meowth := Cat{Name: "Meowth"}

	introduce(growlithe) // implementation for Dog object called
	introduce(meowth)    // implementation for Cat object called

	// alternately, we can store interface in variables and then the interface will contain objects
	var animal Speaker
	animal = growlithe

	// now growlithe becomes an implementation of the Speaker interface since it has implemented all the methods defined for the interface
	fmt.Println(animal.MakeNoise())

	animal = meowth // the interface now contains meowth, the Cat object
	fmt.Println(animal.MakeNoise())

}
