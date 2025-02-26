package utils

func CreateSliceOneDimensional(n int) (result []int) {
	result = make([]int, n)
	return result
}

func CreateSliceTwoDimensional(n int) (result [][]int) {
	result = make([][]int, n)
	for i := 0; i < n; i++ {
		result[i] = make([]int, n)
	}
	return result
}

func CreateSliceTwoDimensionalRect(m int, n int) (result [][]int) {
	result = make([][]int, m)
	for i := 0; i < m; i++ {
		result[i] = make([]int, n+1)
	}
	return result
}
