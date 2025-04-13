package main

import (
	"fmt"

	"github.com/KylerJacobson/grokking-algorithms/search"
	"github.com/KylerJacobson/grokking-algorithms/sort"
)

func main() {
	// ######################################### Chapter 1 #########################################
	testSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	index, ok := search.BinarySearch(testSlice, 100)
	if !ok {
		fmt.Println("Value is not in the slice")
	} else {
		fmt.Printf("Value is at position %d\n", index)
	}

	// ######################################### Chapter 2 #########################################
	needsSorted := []int{1, 4, 2, 9, 10, 5}
	sorterArr := sort.SelectionSort(needsSorted)
	sort.SelectionSortInPlace(needsSorted)
	fmt.Println(needsSorted)
	fmt.Println(sorterArr)

}
