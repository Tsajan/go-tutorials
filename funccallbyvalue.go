package main

import "fmt"

func funcCallByValue(x int) {
    // update x to some random value
    x = 49
    fmt.Println("Value of x in fn when passed by value is: ", x)

}

func main() {
    x := 5
    fmt.Println("Value before fn call is: ", x)

    // function call with pass by values
    funcCallByValue(x)

    // print the value after function call by value
    fmt.Println("Value after fn call is: ", x)
}
