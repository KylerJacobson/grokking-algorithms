package main

// Chapter 3
func RecursiveFactorial(val int) int {
	//base case
	if val == 1 {
		return 1
	}
	retVal := val * RecursiveFactorial(val-1)

	return retVal
}

func RecursiveAdd(slc []int) int {
	if len(slc) == 1 {
		return slc[0]
	}
	return slc[0] + RecursiveAdd(slc[1:])
}

func RecursiveLen(slc []int) int {
	if len(slc) == 1 {
		return 1
	}
	return 1 + RecursiveLen(slc[1:])
}

func RecursiveMax(slc []int) int {
	if len(slc) == 1 {
		return slc[0]
	}
	if slc[0] > RecursiveMax(slc[1:]) {
		return slc[0]
	} else {
		return RecursiveMax(slc[1:])
	}
}

func RecursiveBinarySearch(slc []int, low, high, target int) int {
	if high >= low {
		mid := low + (high - low) / 2
		if slc[mid] == target {
			return mid
		}
		if slc[mid] > target {
			return RecursiveBinarySearch(slc, low, mid-1, target)
		} else {
			return RecursiveBinarySearch(slc, mid+1, high, target)
		}
	}
	return -1
}
