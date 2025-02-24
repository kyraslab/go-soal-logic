package logic_01

func SoalNo04(n int) []int {
	slice := make([]int, n)
	num := 19

	for i := 0; i < n; i++ {
		slice[i] = num
		num -= 2
	}

	return slice
}
