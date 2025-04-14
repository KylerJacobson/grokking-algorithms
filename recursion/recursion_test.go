package recursion_test

import (
	"testing"

	"github.com/KylerJacobson/grokking-algorithms/recursion"
)

func Test_Factorial(t *testing.T) {
	tests := []struct {
		Name           string
		NumToTest      int
		ExpectedResult int
		ExpectedError  bool
	}{
		{
			Name:           "Five",
			NumToTest:      5,
			ExpectedResult: 120,
			ExpectedError:  false,
		},
		{
			Name:           "Fifteen",
			NumToTest:      15,
			ExpectedResult: 1307674368000,
			ExpectedError:  false,
		},
		{
			Name:           "One",
			NumToTest:      1,
			ExpectedResult: 1,
			ExpectedError:  false,
		},
		{
			Name:           "Zero",
			NumToTest:      0,
			ExpectedResult: 1,
			ExpectedError:  false,
		},
		{
			Name:           "Negative",
			NumToTest:      -1,
			ExpectedResult: 0,
			ExpectedError:  true,
		},
	}

	for _, test := range tests {
		result, err := recursion.Factorial(test.NumToTest)
		if result != test.ExpectedResult {
			t.Errorf("For test %s got: %d but expected: %d", test.Name, result, test.ExpectedResult)
		}

		if err != nil && !test.ExpectedError {
			t.Errorf("In test %s, got an error when we did not expect one: %v", test.Name, err)
		}
	}
}
func Test_Sum(t *testing.T) {
	tests := []struct {
		Name           string
		ValuesToTest   []int
		ExpectedResult int
	}{
		{
			Name:           "Empty Slice",
			ValuesToTest:   []int{},
			ExpectedResult: 0,
		},
		{
			Name:           "Single Element",
			ValuesToTest:   []int{5},
			ExpectedResult: 5,
		},
		{
			Name:           "Multiple Elements",
			ValuesToTest:   []int{1, 2, 3, 4, 5},
			ExpectedResult: 15,
		},
		{
			Name:           "Negative Numbers",
			ValuesToTest:   []int{-1, -2, -3, -4, -5},
			ExpectedResult: -15,
		},
		{
			Name:           "Mixed Numbers",
			ValuesToTest:   []int{-1, 2, -3, 4, -5},
			ExpectedResult: -3,
		},
	}

	for _, test := range tests {
		result := recursion.Sum(test.ValuesToTest)
		if result != test.ExpectedResult {
			t.Errorf("For test %s got: %d but expected: %d", test.Name, result, test.ExpectedResult)
		}
	}
}
func Test_Length(t *testing.T) {
	tests := []struct {
		Name           string
		ValuesToTest   []int
		ExpectedResult int
	}{
		{
			Name:           "Empty Slice",
			ValuesToTest:   []int{},
			ExpectedResult: 0,
		},
		{
			Name:           "Single Element",
			ValuesToTest:   []int{5},
			ExpectedResult: 1,
		},
		{
			Name:           "Multiple Elements",
			ValuesToTest:   []int{1, 2, 3, 4, 5},
			ExpectedResult: 5,
		},
		{
			Name:           "Negative Numbers",
			ValuesToTest:   []int{-1, -2, -3, -4, -5},
			ExpectedResult: 5,
		},
		{
			Name:           "Mixed Numbers",
			ValuesToTest:   []int{-1, 2, -3, 4, -5},
			ExpectedResult: 5,
		},
	}

	for _, test := range tests {
		result := recursion.Length(test.ValuesToTest)
		if result != test.ExpectedResult {
			t.Errorf("For test %s got: %d but expected: %d", test.Name, result, test.ExpectedResult)
		}
	}
}
func Test_Max(t *testing.T) {
	tests := []struct {
		Name           string
		ValuesToTest   []int
		ExpectedResult int
		ExpectedError  bool
	}{
		{
			Name:           "Empty Slice",
			ValuesToTest:   []int{},
			ExpectedResult: -1,
			ExpectedError:  true,
		},
		{
			Name:           "Single Element",
			ValuesToTest:   []int{5},
			ExpectedResult: 5,
			ExpectedError:  false,
		},
		{
			Name:           "Multiple Elements",
			ValuesToTest:   []int{1, 2, 3, 4, 5},
			ExpectedResult: 5,
			ExpectedError:  false,
		},
		{
			Name:           "Negative Numbers",
			ValuesToTest:   []int{-1, -2, -3, -4, -5},
			ExpectedResult: -1,
			ExpectedError:  false,
		},
		{
			Name:           "Mixed Numbers",
			ValuesToTest:   []int{-1, 2, -3, 4, -5},
			ExpectedResult: 4,
			ExpectedError:  false,
		},
	}

	for _, test := range tests {
		result, err := recursion.Max(test.ValuesToTest)
		if result != test.ExpectedResult {
			t.Errorf("For test %s got: %d but expected: %d", test.Name, result, test.ExpectedResult)
		}

		if (err != nil) != test.ExpectedError {
			t.Errorf("In test %s, got error: %v, but expected error: %v", test.Name, err != nil, test.ExpectedError)
		}
	}
}
func Test_BinarySearch(t *testing.T) {
	tests := []struct {
		Name           string
		ValuesToTest   []int
		Target         int
		ExpectedResult int
		ExpectedError  bool
	}{
		{
			Name:           "Target Found in Middle",
			ValuesToTest:   []int{1, 2, 3, 4, 5},
			Target:         3,
			ExpectedResult: 2,
			ExpectedError:  false,
		},
		{
			Name:           "Target Found at Start",
			ValuesToTest:   []int{1, 2, 3, 4, 5},
			Target:         1,
			ExpectedResult: 0,
			ExpectedError:  false,
		},
		{
			Name:           "Target Found at End",
			ValuesToTest:   []int{1, 2, 3, 4, 5},
			Target:         5,
			ExpectedResult: 4,
			ExpectedError:  false,
		},
		{
			Name:           "Target Not Found",
			ValuesToTest:   []int{1, 2, 3, 4, 5},
			Target:         6,
			ExpectedResult: -1,
			ExpectedError:  true,
		},
		{
			Name:           "Empty Slice",
			ValuesToTest:   []int{},
			Target:         1,
			ExpectedResult: -1,
			ExpectedError:  true,
		},
		{
			Name:           "Single Element Found",
			ValuesToTest:   []int{1},
			Target:         1,
			ExpectedResult: 0,
			ExpectedError:  false,
		},
		{
			Name:           "Single Element Not Found",
			ValuesToTest:   []int{1},
			Target:         2,
			ExpectedResult: -1,
			ExpectedError:  true,
		},
		{
			Name:           "Target Found in Larger Array",
			ValuesToTest:   []int{1, 3, 5, 7, 9, 11, 13, 15},
			Target:         7,
			ExpectedResult: 3,
			ExpectedError:  false,
		},
		{
			Name:           "Target Found in Larger Array",
			ValuesToTest:   []int{1, 3, 5, 7, 9, 11, 13, 15},
			Target:         11,
			ExpectedResult: 5,
			ExpectedError:  false,
		},
	}

	for _, test := range tests {
		result, err := recursion.BinarySearch(test.ValuesToTest, test.Target)
		if result != test.ExpectedResult {
			t.Errorf("For test %s got: %d but expected: %d", test.Name, result, test.ExpectedResult)
		}

		if (err != nil) != test.ExpectedError {
			t.Errorf("In test %s, got error: %v, but expected error: %v", test.Name, err != nil, test.ExpectedError)
		}
	}
}
