package main

import "slices"

func QuickSort(slc []int) []int {
	// base case
	if len(slc) < 2 {
		return slc
	}

	// recursive case
	pivot := slc[0]
	smaller := []int{}
	greater := []int{}
	for _, v := range slc[1:] {
		if v <= pivot {
			smaller = append(smaller, v)
		} else {
			greater = append(greater, v)
		}
	}
	
	return slices.Concat(QuickSort(smaller), []int{pivot}, QuickSort(greater))
}
