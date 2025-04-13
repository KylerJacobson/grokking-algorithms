package search

import "testing"

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		name          string
		values        []int
		target        int
		expectedIndex int
		expectedFound bool
	}{
		{"Target in middle", []int{1, 2, 3, 4, 5}, 3, 2, true},
		{"Target at start", []int{1, 2, 3, 4, 5}, 1, 0, true},
		{"Target at end", []int{1, 2, 3, 4, 5}, 5, 4, true},
		{"Target not in list", []int{1, 2, 3, 4, 5}, 6, -1, false},
		{"Empty list", []int{}, 3, -1, false},
		{"Single element - found", []int{3}, 3, 0, true},
		{"Single element - not found", []int{3}, 1, -1, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			index, found := BinarySearch(tt.values, tt.target)
			if index != tt.expectedIndex || found != tt.expectedFound {
				t.Errorf("BinarySearch(%v, %d) = (%d, %v); want (%d, %v)",
					tt.values, tt.target, index, found, tt.expectedIndex, tt.expectedFound)
			}
		})
	}
}
