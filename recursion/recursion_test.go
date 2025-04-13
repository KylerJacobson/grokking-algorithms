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
