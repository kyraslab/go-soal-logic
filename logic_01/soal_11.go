package logic_01

func SoalNo11(n int) []interface{} {
	slice := make([]interface{}, n)
	num := 1

	for i := 0; i < n; i++ {
		if i%2 == 0 {
			slice[i] = "buzz"
		} else {
			slice[i] = num

			if i < 3 {
				num *= 3
			} else {
				num *= 2
			}
		}
	}

	return slice
}
