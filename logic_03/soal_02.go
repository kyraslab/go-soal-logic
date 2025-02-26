package logic_03

import "github.com/kyraslab/go-soal-logic/utils"

func SoalNo02(n int) [][]int {
	slice := utils.CreateSliceTwoDimensional(n)
	num := 1

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if j >= i {
				if i%2 == 0 {
					slice[i][j] = num
				} else {
					slice[i][n-1-j+i] = num
				}

				num += 2
			}
		}
	}

	return slice
}
