package main

import "fmt"

func main() {
	var twoD = [3][4]int{{0, 1, 2, 3}, {4, 5, 6, 7}, {8, 9, 10, 11}}
	for i := 0; i < 3; i++ {
		for j := 0; j < 4; j++ {
			fmt.Printf("twod [%d][%d] = %d\n", i, j, twoD[i][j])
		}
	}
}
