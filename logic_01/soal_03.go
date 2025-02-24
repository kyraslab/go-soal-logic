package logic_01

func SoalNo03(n int) []int {
	slice := make([]int, n)
	num := 3

	for i := 0; i < n; i++ {
		slice[i] = num
		num += 3
	}

	return slice
}
