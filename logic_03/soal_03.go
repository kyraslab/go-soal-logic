package logic_03

import "github.com/kyraslab/go-soal-logic/utils"

func SoalNo03(n int) [][]int {
	slice := utils.CreateSliceTwoDimensional(n)
	num := 1

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				slice[i][j] = 2
				continue
			}

			if i+j <= n-1 {
				if i%2 == 0 && i != 0 {
					num += 2
				} else {
					num += 3
				}

				if i%2 == 0 {
					slice[i][j] = num
				} else {
					slice[i][n-1-j-i] = num
				}
			}
		}
	}

	return slice
}
