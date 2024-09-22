package main

import (
	"fmt"
)	
func main() {
	testSlice := []int{1,2,3,4,5,6,7,8,9,10,11,12}

	pos := binarySearch(testSlice, 12)
	
	if pos != nil {
		fmt.Printf("Found value at position %d\n", *pos)
	} else {
		fmt.Println("Value is not in the array")
	}

	sorterArr := selectionSort([]int{1,4,2,9,10,5})
	fmt.Println(sorterArr)
}