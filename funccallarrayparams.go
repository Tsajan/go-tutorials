package main

import "fmt"

func main() {
    arr := [...]int{2, 4, 6, 8, 10}

    arr2 := [...]int{10, 20, 30, 40, 50}
    fmt.Println("Original array is: ", arr)

    toggleArray(arr) // updated array element is returned; but since array are passed by value does it update the changes?
    fmt.Println("Original array is: ", arr)


    fmt.Println("Array 2 is: ", arr2)
    toggleArrayUsingDereferenceOperator(&arr2)

    fmt.Println("After pass by pointer, arr2 is: ", arr2)
}

func toggleArray(myArr [5]int) {
    fmt.Println("Argument myArr is: ", myArr)
    fmt.Println("Updating array element at 2nd index to 5")
    myArr[2] = 5

    fmt.Println("Updated myArr parameter is: ", myArr)
}

func toggleArrayUsingDereferenceOperator(myArrPointer *[5]int) [5]int {
    fmt.Println("Parameter myArrPointer dereferenced is: ", *myArrPointer)

    myArrPointer[0] = 100 // updating the pointer referenced value to 100
    return *myArrPointer

}
