package main

import "fmt"

func main() {
    multidim_slice := [][]string{{"sajan", "rajan", "saju", "raman", "kriti"}, {"anish", "ruchi", "sagar", "binamra", "jyoti"}}

    fmt.Println("Multi-dimensional slice is: ", multidim_slice)

    // appending to a particular dimension within a multi-dimensional array
    multidim_slice[0] = append(multidim_slice[0], "joel") // need to specify which index of the multi-dimensional slice to append to 
    multidim_slice[1] = append(multidim_slice[1], "salil") // same here

    fmt.Println("Updated multi-dimensional slice is: ", multidim_slice)
}
