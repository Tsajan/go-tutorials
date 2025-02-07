package main

import "fmt"

func add(a int, b int) (int) {
    return (a+b)
}

func main() {
    x := add(10, 20)
    fmt.Println(x)
}
