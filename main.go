package main

import (
	"fmt"
)

func main() {
	// ######################################### Chapter 1 #########################################
	testSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	pos := binarySearch(testSlice, 12)

	if pos != nil {
		fmt.Printf("Found value at position %d\n", *pos)
	} else {
		fmt.Println("Value is not in the array")
	}

	// ######################################### Chapter 2 #########################################
	sorterArr := selectionSort([]int{1, 4, 2, 9, 10, 5})
	fmt.Println(sorterArr)

	// ######################################### Chapter 3 #########################################
	factorialArg := 5
	fmt.Printf("Factorial %d is: %d\n", factorialArg, RecursiveFactorial(factorialArg))

	// ######################################### Chapter 4 #########################################

	// 4.1 Write a function to recursively add an array
	slc := []int{1, 2, 3, 4, 5, 6}
	fmt.Printf("Array Sum for array %d is: %d\n", slc, RecursiveAdd(slc))

	// 4.2 Write a function to recursively count the elements in an array
	fmt.Printf("Array len for array %d is: %d\n", slc, RecursiveLen(slc))

	// 4.3 Write a function to return the maximum number in a list
	fmt.Printf("The largest value in array %d is: %d\n", slc, RecursiveMax(slc))

	// 4.4 Write recursive binary search
	target := 12
	slc2 := []int{6, 7, 8, 9, 10, 11, 12}
	fmt.Printf("The target value %d is at position %d in array %d", target, RecursiveBinarySearch(slc2, 0, len(slc2)-1, target), slc2)
}
