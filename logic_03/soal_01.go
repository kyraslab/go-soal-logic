package logic_03

import "github.com/kyraslab/go-soal-logic/utils"

func SoalNo01(n int) [][]int {
	slice := utils.CreateSliceTwoDimensional(n)
	num := 1

	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			if i%2 == 0 {
				slice[i][j] = num
			} else {
				slice[i][i-j] = num
			}

			num += 2
		}
	}

	return slice
}
