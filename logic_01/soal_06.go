package logic_01

func SoalNo06(n int) []int {
	slice := make([]int, n)
	num := 30

	for i := 0; i < n; i++ {
		slice[i] = num
		num -= 3
	}

	return slice
}
