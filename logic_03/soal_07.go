package logic_03

import "github.com/kyraslab/go-soal-logic/utils"

func SoalNo07(n int) [][]int {
	slice := utils.CreateSliceTwoDimensional(n)
	num := 1
	midpoint := (n - 1) / 2

	for i := 0; i <= midpoint; i++ {
		for j := 0; j < n; j++ {
			if i+j >= midpoint && i-j <= midpoint && j-i <= midpoint && i+j <= n-1+midpoint {
				if i%2 == 0 {
					slice[i][n-1-j] = num
					slice[n-1-i][n-1-j] = num
				} else {
					slice[i][j] = num
					slice[n-1-i][j] = num
				}

				num += 2
			}
		}
	}

	return slice
}
