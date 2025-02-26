package logic_03

import "github.com/kyraslab/go-soal-logic/utils"

func SoalNo08(n int) [][]int {
	slice := utils.CreateSliceTwoDimensional(n)
	num := 1
	midpoint := (n - 1) / 2

	for i := 0; i <= midpoint; i++ {
		for j := 0; j < n; j++ {
			if i+j >= midpoint && i-j <= midpoint && j-i <= midpoint && i+j <= n-1+midpoint {
				if i%2 == 0 {
					slice[n-1-j][i] = num
					slice[n-1-j][n-1-i] = num
				} else {
					slice[j][i] = num
					slice[n-1-j][i] = num
				}

				num += 2
			}
		}
	}

	return slice
}
