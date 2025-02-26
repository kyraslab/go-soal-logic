package logic_03

import "github.com/kyraslab/go-soal-logic/utils"

func SoalNo05(n int) [][]int {
	slice := utils.CreateSliceTwoDimensional(n)
	num := 1
	midpoint := (n - 1) / 2

	for i := 0; i <= midpoint; i++ {
		for j := 0; j <= i; j++ {
			if i%2 == 0 {
				slice[i][j] = num
				slice[i][n-1-j] = num
				slice[n-1-i][j] = num
				slice[n-1-i][n-1-j] = num
			} else {
				slice[i][i-j] = num
				slice[i][n-1-i+j] = num
				slice[n-1-i][i-j] = num
				slice[n-1-i][n-1-i+j] = num
			}

			num += 2
		}
	}

	return slice
}
