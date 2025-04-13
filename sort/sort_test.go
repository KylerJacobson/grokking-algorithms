package sort

import (
	"reflect"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{input: []int{5, 3, 6, 2, 10}, expected: []int{2, 3, 5, 6, 10}},
		{input: []int{1, 2, 3, 4, 5}, expected: []int{1, 2, 3, 4, 5}},
		{input: []int{5, 4, 3, 2, 1}, expected: []int{1, 2, 3, 4, 5}},
		{input: []int{10, -1, 2, 5, 0}, expected: []int{-1, 0, 2, 5, 10}},
		{input: []int{}, expected: []int{}},
	}

	for _, test := range tests {
		result := SelectionSort(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("SelectionSort(%v) = %v; want %v", test.input, result, test.expected)
		}
	}
}
func TestSelectionSortInPlace(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{input: []int{5, 3, 6, 2, 10}, expected: []int{2, 3, 5, 6, 10}},
		{input: []int{1, 2, 3, 4, 5}, expected: []int{1, 2, 3, 4, 5}},
		{input: []int{5, 4, 3, 2, 1}, expected: []int{1, 2, 3, 4, 5}},
		{input: []int{10, -1, 2, 5, 0}, expected: []int{-1, 0, 2, 5, 10}},
	}

	for _, test := range tests {
		inputCopy := append([]int(nil), test.input...) // Create a copy to avoid modifying the original input
		SelectionSortInPlace(inputCopy)
		if !reflect.DeepEqual(inputCopy, test.expected) {
			t.Errorf("SelectionSortInPlace(%v) = %v; want %v", test.input, inputCopy, test.expected)
		}
	}
}
