// You can edit this code!
// Click here and start typing.
package tutorials

import "fmt"

type Maanche struct {
	Name    string
	Surname string
	Hobbies []string
	id      string
}

func (person *Maanche) GetFullName() string {
	return fmt.Sprintf("Full Name is: %s %s", person.Name, person.Surname)
}

func (person Maanche) SetFullNameUsingValue(name string, surname string) {
	person.Name = name
	person.Surname = surname
	fmt.Printf("Because Maanche is a value, not a pointer, applied changes are not reflected even after setting %s %s\n", name, surname)
}

func (person *Maanche) SetFullNameUsingPointer(name string, surname string) {
	person.Name = name
	person.Surname = surname
	fmt.Printf("Because Maanche is a pointer, applied changes are reflected to the underlying object after setting %s %s\n", name, surname)
}

func ExecuteMethodsWithPointers() {
	p := Maanche{
		Name:    "Sajan",
		Surname: "Maharjan",
		Hobbies: []string{"cycyling", "running", "reading"},
		id:      "294",
	}

	p.SetFullNameUsingValue("Declan", "Arrozi")
	fmt.Println(p.GetFullName())

	p.SetFullNameUsingPointer("Declan", "Arrozi")
	fmt.Println(p.GetFullName())

}
