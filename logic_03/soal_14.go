package logic_03

import "github.com/kyraslab/go-soal-logic/utils"

func SoalNo14(n int) [][]int {
	slice := utils.CreateSliceTwoDimensional(n)
	num := 1
	inc := 0
	res := 0

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				slice[j][i] = 1
				continue
			} else if i == n-1 && j >= n-2 {
				slice[j][i] = res + 2
				res = slice[j][i]
				continue
			}

			res = num + inc
			inc = num
			num = res

			if i%2 == 0 {
				slice[j][i] = res
			} else {
				slice[n-1-j][i] = res
			}

		}
	}

	return slice
}
