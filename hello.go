package main

import "fmt"

func main() {
    var myString string = `Hello\n world`
    var myString2 string = "Hello\n world 2"
    var myBool bool
    fmt.Println(myString)
    fmt.Println(myString2)
    fmt.Println(myBool)

    myBool2 := true
    fmt.Println(myBool2)

    const pi = 3.14
    fmt.Println(pi)
}
