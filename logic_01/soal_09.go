package logic_01

func SoalNo09(n int) []int {
	slice := make([]int, n)
	midpointIndex := n / 2
	num := 3

	for i := 0; i < n; i++ {
		slice[i] = num
		if i == midpointIndex-1 && n%2 == 0 {
			continue
		} else if i < midpointIndex {
			num += 3
		} else {
			num -= 3
		}
	}

	return slice
}
