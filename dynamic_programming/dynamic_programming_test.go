package dynamicprogramming

import "testing"

func TestLongestSubstring(t *testing.T) {
	tests := []struct {
		name           string
		s1             string
		s2             string
		expectedOutput string
		wantError      bool
	}{
		{
			name:           "Simple matching substring",
			s1:             "fish",
			s2:             "hish",
			expectedOutput: "ish",
		},
		{
			name:           "No matching substring",
			s1:             "abc",
			s2:             "def",
			expectedOutput: "",
		},
		{
			name:           "Empty strings",
			s1:             "",
			s2:             "",
			expectedOutput: "",
		},
		{
			name:           "One empty string",
			s1:             "hello",
			s2:             "",
			expectedOutput: "",
		},
		{
			name:           "Identical strings",
			s1:             "hello",
			s2:             "hello",
			expectedOutput: "hello",
		},
		{
			name:           "Multiple matching substrings",
			s1:             "abcdefg",
			s2:             "xyzdef",
			expectedOutput: "def",
		},
		{
			name:           "Arbitrary",
			s1:             "artificial intelligence",
			s2:             "intelligent systems",
			expectedOutput: "intelligen",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := LongestSubstring(tt.s1, tt.s2)
			if got != tt.expectedOutput {
				t.Errorf("LongestSubstring(%q, %q) = %q, want %q",
					tt.s1, tt.s2, got, tt.expectedOutput)
			}
		})
	}
}
