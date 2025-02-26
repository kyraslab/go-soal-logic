package logic_03

import "github.com/kyraslab/go-soal-logic/utils"

func SoalNo13(n int) [][]int {
	slice := utils.CreateSliceTwoDimensional(n)
	num := 1
	midpoint := (n - 1) / 2

	for i := 0; i <= midpoint; i++ {
		for j := 0; j < n; j++ {
			if i+j >= midpoint && i-j <= midpoint && j-i <= midpoint && i+j <= n-1+midpoint {
				slice[j][i] = num
				slice[j][n-1-i] = num

				if (j%2 == 0 && i%2 == 1) || (j%2 == 1 && i%2 == 0) {
					slice[j][i] = 0
					slice[j][n-1-i] = 0
				}
			}

		}
		num += 2
	}

	return slice
}
