package main

import "fmt"

func main() {
    fmt.Println("This go routine performs matrix multiplication using slices")

    matrixA := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
    matrixB := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}

    fmt.Println("Matrix A is: ", matrixA)
    fmt.Println("Length of Matrix A is: ", len(matrixA))
    fmt.Println("Matrix B is: ", matrixB)
    
    
    res, err := matrixMultiplication(matrixA, matrixB)
    if err != nil {
        fmt.Println("Error: ", err)
    }
    fmt.Println("Result: ", res)

}

func matrixMultiplication(matA [][]int, matB [][]int) ([][]int, error) {

    // create an empty matrix to store the results using the make function
    res := make([][]int, 3)
    for idx := range res {
        res[idx] = make([]int, 3)
    }

    // perform matrix multiplication
    for rowIdx, row := range res {
        for colIdx := range row {
            for k := 0; k < len(res); k++ {
                res[rowIdx][colIdx] += matA[rowIdx][k] * matB[k][colIdx]
            }
        }
    }

    return res, nil
}
