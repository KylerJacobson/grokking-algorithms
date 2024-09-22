package main

// Chapter 1
func binarySearch(slc []int, target int) *int {
	low := 0
	high := len(slc) - 1
	mid := (high - low) / 2
	for slc[mid] != target {
		if slc[mid] == target {
			return &mid
		} else if high - low < 1 {
			return nil
		} else if slc[mid] > target {
			high = mid - 1
		} else if slc[mid] < target {
			low = mid + 1
		}
		mid = (high + low) / 2
	}
	return &mid
}