package logic_03

import "github.com/kyraslab/go-soal-logic/utils"

func SoalNo12b(n int) [][]int {
	slice := utils.CreateSliceTwoDimensionalRect(n, n*4/3)
	num := 1
	midpoint := (n - 1) / 2

	for i := 0; i <= midpoint; i++ {
		for j := 0; j < n; j++ {
			if i+j == n-1 {
				slice[i][midpoint+j] = i*2 + 1
				slice[n-1-i][midpoint+j] = i*2 + 1
			}

			if j >= midpoint && i+j >= n-1 {
				if i%2 == 0 {
					slice[i][midpoint+n-1-(2*n-2-i-j)] = num
					slice[i][midpoint-(n-1-(2*n-2-i-j))] = num
					slice[n-1-i][midpoint+n-1-(2*n-2-i-j)] = num
					slice[n-1-i][midpoint-(n-1-(2*n-2-i-j))] = num
				} else {
					slice[i][midpoint+n-1-j] = num
					slice[i][midpoint-n+1+j] = num
					slice[n-1-i][midpoint+n-1-j] = num
					slice[n-1-i][midpoint-n+1+j] = num
				}
				num += 2
			}
		}
	}

	return slice
}
