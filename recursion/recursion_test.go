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
func Test_QuickSort(t *testing.T) {
	tests := []struct {
		Name           string
		ValuesToTest   []int
		ExpectedResult []int
	}{
		{
			Name:           "Empty Slice",
			ValuesToTest:   []int{},
			ExpectedResult: []int{},
		},
		{
			Name:           "Single Element",
			ValuesToTest:   []int{5},
			ExpectedResult: []int{5},
		},
		{
			Name:           "Already Sorted",
			ValuesToTest:   []int{1, 2, 3, 4, 5},
			ExpectedResult: []int{1, 2, 3, 4, 5},
		},
		{
			Name:           "Reverse Sorted",
			ValuesToTest:   []int{5, 4, 3, 2, 1},
			ExpectedResult: []int{1, 2, 3, 4, 5},
		},
		{
			Name:           "Unsorted",
			ValuesToTest:   []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5},
			ExpectedResult: []int{1, 1, 2, 3, 3, 4, 5, 5, 5, 6, 9},
		},
		{
			Name:           "Negative Numbers",
			ValuesToTest:   []int{-3, -1, -4, -1, -5, -9, -2, -6, -5, -3, -5},
			ExpectedResult: []int{-9, -6, -5, -5, -5, -4, -3, -3, -2, -1, -1},
		},
		{
			Name:           "Mixed Numbers",
			ValuesToTest:   []int{-3, 1, -4, 1, 5, -9, 2, -6, 5, -3, 5},
			ExpectedResult: []int{-9, -6, -4, -3, -3, 1, 1, 2, 5, 5, 5},
		},
		{
			Name:           "Large Array",
			ValuesToTest:   []int{42, 17, 99, 3, 28, 56, 71, 83, 12, 33, 5, 62, 41, 88, 7, 21, 14, 50, 94, 39, 10, 67, 25, 73, 81, 19, 30, 55, 2, 70, 45, 8, 91, 22, 36, 60, 77, 16, 85, 29, 52, 64, 97, 11, 75, 34, 48, 66, 89, 4, 27, 57, 80, 15, 38, 63, 23, 92, 6, 47, 74, 31, 53, 86, 20, 44, 69, 9, 58, 82, 18, 35, 90, 1, 24, 51, 76, 32, 65, 96, 13, 43, 61, 87, 26, 49, 72, 37, 59, 84, 46, 68, 93, 54, 79, 98, 40, 78, 95},
			ExpectedResult: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90, 91, 92, 93, 94, 95, 96, 97, 98, 99},
		},
	}

	for _, test := range tests {
		result := recursion.QuickSort(test.ValuesToTest)
		if !equalSlices(result, test.ExpectedResult) {
			t.Errorf("For test %s got: %v but expected: %v", test.Name, result, test.ExpectedResult)
		}
	}
}

func equalSlices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
