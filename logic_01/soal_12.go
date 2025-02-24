package logic_01

func SoalNo12(n int) []int {
	slice := make([]int, n)
	basePattern := []int{1, 3, 5, 7}

	for i := 0; i < n; i++ {
		slice[i] = basePattern[i%len(basePattern)]
	}

	return slice
}
