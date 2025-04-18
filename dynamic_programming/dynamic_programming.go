package dynamicprogramming

func LongestSubstring(s1, s2 string) string {

	longest := 0
	if len(s1) > len(s2) {
		longest = len(s1)
	} else {
		longest = len(s2)
	}
	matrix := make([][]int, longest)
	for i := 0; i < len(matrix); i++ {
		matrix[i] = make([]int, longest)
	}

	maxRow, maxVal := 0, 0
	for i, v1 := range s1 {
		for j, v2 := range s2 {
			if v1 == v2 {
				if i-1 >= 0 && j-1 >= 0 {
					matrix[i][j] = matrix[i-1][j-1] + 1
				} else {
					matrix[i][j] = 1
				}
				if matrix[i][j] > maxVal {
					maxVal = matrix[i][j]
					maxRow = i
				}
			} else {
				matrix[i][j] = 0
			}
		}
	}

	runes := []rune(s1)

	result := []rune{}
	for i := 0; i < maxVal; i++ {
		index := maxRow - maxVal + 1 + i
		result = append(result, runes[index])
	}
	return string(result)
}
