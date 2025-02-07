package main

import "fmt"

func printText() {
    fmt.Println("First print line")
    defer fmt.Println("This will be printed at the end")

    fmt.Println("This will be printed second")
}

func main() {
    defer fmt.Println("This will ultimately be printed in the end")
    printText()
}
