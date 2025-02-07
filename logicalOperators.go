package main

import "fmt"

func main() {
    a := 10
    b := 20

    // logical AND operator
    if a != b && a <= b {
        fmt.Println("True")
    }

    // logical OR operator
    if a != b || a >= b {
        fmt.Printf("True\n")
    }

    // logical NOT operator
    if !(a == b) {
        fmt.Println("a not equals b")
    }

    // bitwise AND operator
    x := 8 ^ 1
    fmt.Println(x)
}
