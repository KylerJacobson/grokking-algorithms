package sort

import "math"

func SelectionSort(values []int) []int {
	sorted := make([]int, len(values))
	seen := make([]bool, len(values))
	numItems := len(values)
	for i := 0; i < numItems; i++ {
		min := math.MaxInt
		minIndex := -1
		for j := 0; j < numItems; j++ {
			if values[j] < min && !seen[j] {
				min = values[j]
				minIndex = j
			}
		}
		seen[minIndex] = true
		sorted[i] = min
	}
	return sorted
}

func SelectionSortInPlace(values []int) {
	numItems := len(values)

	for i := 0; i < numItems-1; i++ {
		minIndex := i
		for j := i + 1; j < numItems; j++ {
			if values[j] < values[minIndex] {
				minIndex = j
			}
		}
		if minIndex != i {
			values[i], values[minIndex] = values[minIndex], values[i]
		}
	}
}
