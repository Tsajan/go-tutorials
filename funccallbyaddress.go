package main

import "fmt"

func replaceXandY(x *int, y *int) {
    // update x and y
    *x = 7
    *y = 18
}


func addAndReplace(x *int, y *int) (z int) {
    z = *x + *y
    *x = 99 
    *y = 100
    return
}
func main() {
    x, y := 5, 10
    fmt.Println("Before pass by address x, y : ", x , y)
    replaceXandY(&x, &y)
    fmt.Println("After pass by address x, y: ", x , y)
    
    z := addAndReplace(&x, &y)
    fmt.Println(x, y, z)
}
