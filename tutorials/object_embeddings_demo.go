package tutorials

import "fmt"

type Person struct {
    firstName string
    lastName string
    age int
    gender byte
    citizenShipAcquired bool
}

// a method on Person type to change the given name; declare a pointer to underlying Person object to reflect write changes
func (person *Person) acquireCitizenShip() {
    person.citizenShipAcquired = true
}

// another method on Person type
func (person Person) greet() {
    fmt.Printf("Hello I'm %s %s", person.firstName, person.lastName)
}

type Student struct {
    Person // embedding or composing or aggregating Person struct into Student object; HAS-A relationship; declaring without variable name person for Person allows promotion
    grade string
    university string
    hasScholarship bool
}

func ObjectEmbeddingDemo() {
    // initialize a Person struct
    p1 := Person { "Sajan", "Maharjan", 32, 'M', false }
    fmt.Println("Person p1 is: ", p1)

    // initialize a Student object
    s1 := Student { p1, "PhD.", "University of Leeds", true } // embedding p1 within s1 
    fmt.Println("Student s1 is: ", s1)

    p1.acquireCitizenShip()
    fmt.Println("After acquiring citizenship, person p1 is: ", p1) // reflects changes in p1
    
    // is the change reflected in s1 too? naturally should be yes; however, this record is not updated via student object yet
    fmt.Println("Student s1 has also updated the original attributes of the underlying component, s1: ", s1)

    s1.acquireCitizenShip() // now changes can be reflected?
    fmt.Println("Student s1 is: ", s1)

    s1.greet() // promotion example 
}
