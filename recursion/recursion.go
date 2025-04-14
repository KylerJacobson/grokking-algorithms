package recursion

import (
	"fmt"
)

func Factorial(num int) (int, error) {
	if num < 0 {
		return 0, fmt.Errorf("Factorial is not defined on negative numbers")
	}

	if num <= 1 {
		return 1, nil
	}

	result, err := Factorial(num - 1)

	return num * result, err
}

func Sum(values []int) int {
	if len(values) == 0 {
		return 0
	}

	return values[0] + Sum(values[1:])
}

func Length(values []int) int {
	if len(values) == 0 {
		return 0
	}

	return 1 + Length(values[1:])
}

func Max(values []int) (int, error) {
	if len(values) == 1 {
		return values[0], nil
	}

	if len(values) == 0 {
		return -1, fmt.Errorf("no maximum on an empty slice")
	}

	result, err := Max(values[1:])

	if result > values[0] {
		return result, err
	} else {
		return values[0], err
	}
}

func BinarySearch(values []int, target int) (int, error) {

	if len(values) == 0 {
		return -1, fmt.Errorf("value is not in the array")
	}

	mid := len(values) / 2

	if values[mid] == target {
		return mid, nil
	}

	if values[mid] < target {
		index, err := BinarySearch(values[mid+1:], target)
		if err != nil {
			return -1, err
		}

		return mid + index + 1, nil
	}

	index, err := BinarySearch(values[:mid], target)
	if err != nil {
		return -1, err
	}

	return index, err
}
