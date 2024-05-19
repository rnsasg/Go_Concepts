package main

import "fmt"

func tictactoe(moves [][]int) string {
	res := [4]string{"Draw", "A", "B", "Pending"}
	// n := 3
	// matrix := make([][]int,3)
	// for i:=0;i<n;i++{
	//     matrix[i] = make([]int,3)
	// }

	var matrix [3][3]int

	//fmt.Println(matrix)

	var flip bool = true

	for i := 0; i < len(moves); i++ {
		x := moves[i][0]
		y := moves[i][1]
		if flip == true {
			matrix[x][y] = 1
		} else {
			matrix[x][y] = 2
		}
		flip = !flip
	}

	for i := 0; i < 3; i++ {
		fmt.Println(matrix[i])
	}

	if verdict(matrix, 1) == true {
		return res[1]
	} else if verdict(matrix, 2) == true {
		return res[2]
	} else if isPending(matrix) == true {
		return res[3]
	}

	return res[0]
}

func isPending(matrix [3][3]int) bool {

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if matrix[i][j] == 0 {
				return true
			}
		}
	}
	return false
}

func verdict(matrix [3][3]int, ele int) bool {

	for i := 0; i < 3; i++ {
		// Rows
		if matrix[i][0] == ele && matrix[i][1] == ele && matrix[i][2] == ele {
			fmt.Printf("Rows[%d]", i)
			return true
		}
		// Columns
		if matrix[0][i] == ele && matrix[1][i] == ele && matrix[2][i] == ele {
			fmt.Printf("Column[%d]", i)
			return true
		}
	}
	// Digonal
	if matrix[0][0] == ele && matrix[1][1] == ele && matrix[2][2] == ele {
		fmt.Printf("Digonal")
		return true
	}
	// Anti Digonal
	if matrix[2][0] == ele && matrix[1][1] == ele && matrix[0][2] == ele {
		fmt.Printf("Anti Digonal")
		return true
	}
	return false
}

func main() {
	//res := tictactoe([][]int{{0, 0}, {2, 0}, {1, 1}, {2, 1}, {2, 2}})
	// res := tictactoe([][]int{{0, 0}, {1, 1}, {0, 1}, {0, 2}, {1, 0}, {2, 0}})
	// res := tictactoe([][]int{{0, 0}, {1, 1}, {2, 0}, {1, 0}, {1, 2}, {2, 1}, {0, 1}, {0, 2}, {2, 2}})
	// res := tictactoe([][]int{{0, 0}, {1, 1}})
	res := tictactoe([][]int{{1, 0}, {2, 2}, {0, 0}, {1, 1}, {0, 1}, {2, 0}, {1, 2}})
	fmt.Println(res)
}
