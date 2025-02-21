package tutorials

import "fmt"

func MatrixScalarProduct() {
	fmt.Println("This go routine performs matrix multiplication using slices")

	matrixA := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	matrixB := [][]int{{10, 20, 30}, {40, 50, 60}, {70, 80, 90}}

	fmt.Println("Matrix A is: ", matrixA)
	fmt.Println("Length of Matrix A is: ", len(matrixA))
	fmt.Println("Matrix B is: ", matrixB)

	res, err := scalarMultiplication(matrixA, matrixB)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Println("Result: ", res)

}

func scalarMultiplication(matA [][]int, matB [][]int) ([][]int, error) {

	// create an empty matrix to store the results using the make function
	res := make([][]int, 3)
	for idx := range res {
		res[idx] = make([]int, 3)
	}

	// perform element wise multiplication
	for rowIdx, row := range res {
		for colIdx := range row {
			res[rowIdx][colIdx] = matA[rowIdx][colIdx] * matB[rowIdx][colIdx]
		}
	}

	return res, nil
}
