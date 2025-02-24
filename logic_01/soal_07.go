package logic_01

func SoalNo07(n int) []int {
	slice := make([]int, n)
	midpointIndex := n / 2
	num := 1

	for i := 0; i < n; i++ {
		slice[i] = num
		if i == midpointIndex-1 && n%2 == 0 {
			continue
		} else if i < midpointIndex {
			num += 2
		} else {
			num -= 2
		}
	}

	return slice
}
