package main

// Chapter 3
func RecursiveFactorial(val int) int {
	//base case
	if val == 1 {
		return 1
	}
	retVal := val * RecursiveFactorial(val - 1)

	return retVal
}