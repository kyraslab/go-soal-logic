package logic_01

func SoalNo10(n int) []interface{} {
	slice := make([]interface{}, n)
	num := 2

	for i := 0; i < n; i++ {
		if i%2 == 1 {
			slice[i] = "fizz"
		} else {
			slice[i] = num
			num *= 2
		}
	}

	return slice
}
