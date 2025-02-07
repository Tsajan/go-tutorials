package main

import "fmt"

func main() {
    testArr := [...]int{1, 2, 3, 4, 5}

    for idx, val := range testArr {
        fmt.Printf("testArr[%d] = %d\n", idx, val)
    }
    arrSum := arraySum(testArr)
    fmt.Println("Sum of elements of array is ", arrSum)
}

func arraySum(arr [5]int) (z int) {
    for _, val := range arr {
        z += val
    }
    return
}
