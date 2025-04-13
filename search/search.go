package search

import "fmt"

func BinarySearch(values []int, target int) (int, bool) {

	low := 0
	high := len(values) - 1
	mid := (high + low) / 2

	iterations := 0
	for low <= high {
		iterations++
		currentValue := values[mid]
		if currentValue == target {
			fmt.Printf("Found the value %d in %d iterations\n", target, iterations)
			return mid, true
		}
		if currentValue < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
		mid = (high + low) / 2
	}
	fmt.Printf("Could not find the value %d in %d iterations\n", target, iterations)
	return -1, false
}
