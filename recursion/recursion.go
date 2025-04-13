package recursion

import "fmt"

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
