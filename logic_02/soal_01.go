package logic_02

func SoalNo01(n int) [][]int {
	slice := make([][]int, n)

	for i := 0; i < n; i++ {
		slice[i] = make([]int, n)
		num := 1

		for j := 0; j < n; j++ {
			slice[i][j] = num
			num += 2
		}
	}

	return slice
}
