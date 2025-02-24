package logic_02

func SoalNo03(n int) [][]int {
	slice := make([][]int, n)
	num := 1

	for i := 0; i < n; i++ {
		slice[i] = make([]int, n)

		for j := 0; j < n; j++ {
			slice[i][j] = num
			num += 2
		}
	}

	return slice
}
