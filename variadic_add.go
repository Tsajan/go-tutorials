package main

import "fmt"

func add(nums ...int) (z int) {
    for _, num := range nums {
        z += num
    }
    return
}

func main() {
    x := add(1, 2, 3, 4, 5)
    fmt.Println("Sum of first 5 numbers: ", x)
}
