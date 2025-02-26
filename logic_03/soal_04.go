package logic_03

import "github.com/kyraslab/go-soal-logic/utils"

func SoalNo04(n int) [][]int {
	slice := utils.CreateSliceTwoDimensional(n)
	num := 1

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i+j >= n-1 {
				if i%2 == 0 {
					slice[i][2*n-2-i-j] = num
				} else {
					slice[i][j] = num
				}
				num += 2
			}
		}
	}

	return slice
}
