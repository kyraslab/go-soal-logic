package logic_03

import "github.com/kyraslab/go-soal-logic/utils"

func SoalNo09(n int) [][]int {
	slice := utils.CreateSliceTwoDimensional(n)
	midpoint := (n - 1) / 2

	for i := 0; i <= midpoint; i++ {
		num := 1

		for j := 0; j <= midpoint; j++ {
			if i+j >= midpoint && i-j <= midpoint && j-i <= midpoint && i+j <= n-1+midpoint {
				slice[i][n-1-j] = num
				slice[n-1-i][n-1-j] = num
				slice[i][j] = num
				slice[n-1-i][j] = num

				num += 2
			}
		}
	}

	return slice
}
