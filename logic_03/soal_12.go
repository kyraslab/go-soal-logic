package logic_03

import "github.com/kyraslab/go-soal-logic/utils"

func SoalNo12(n int) [][]int {
	slice := utils.CreateSliceTwoDimensional(n)
	num := 1
	midpoint := (n - 1) / 2

	for i := 0; i <= midpoint; i++ {
		for j := 0; j < n; j++ {
			if i+j == n-1 {
				slice[i][j] = i*2 + 1
				slice[n-1-i][j] = i*2 + 1
			}

			if j >= midpoint && i+j >= n-1 {
				if i%2 == 0 {
					slice[i][n-1-(2*n-2-i-j)] = num
					slice[n-1-i][n-1-(2*n-2-i-j)] = num
				} else {
					slice[i][n-1-j] = num
					slice[n-1-i][n-1-j] = num
				}
				num += 2
			}
		}
	}

	return slice
}
