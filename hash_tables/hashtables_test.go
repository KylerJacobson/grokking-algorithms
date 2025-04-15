package hashtables

import (
	"testing"
)

func TestHashMap_Get(t *testing.T) {
	tests := []struct {
		name     string
		hashSize int
		setKey   string
		setValue string
		getKey   string
		want     string
		found    bool
	}{
		{
			name:     "Get existing key",
			hashSize: 10,
			setKey:   "test_key",
			setValue: "test_value",
			getKey:   "test_key",
			want:     "test_value",
			found:    true,
		},
		{
			name:     "Get non-existing key",
			hashSize: 10,
			setKey:   "test_key",
			setValue: "test_value",
			getKey:   "non_existing_key",
			want:     "",
			found:    false,
		},
		{
			name:     "Get with same hash but different key",
			hashSize: 1, // Force collision
			setKey:   "key1",
			setValue: "value1",
			getKey:   "key2",
			want:     "",
			found:    false,
		},
		{
			name:     "Get when multiple values in same bucket",
			hashSize: 1, // Force collision
			setKey:   "key1",
			setValue: "value1",
			getKey:   "key1",
			want:     "value1",
			found:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := NewHashMap(tt.hashSize)
			h.Set(tt.setKey, tt.setValue)

			// For collision test where we need two values in same bucket
			if tt.name == "Get when multiple values in same bucket" {
				h.Set("key2", "value2")
			}

			got, found := h.Get(tt.getKey)
			if found != tt.found {
				t.Errorf("HashMap.Get() found = %v, want %v", found, tt.found)
			}
			if got != tt.want {
				t.Errorf("HashMap.Get() got = %v, want %v", got, tt.want)
			}
		})
	}
}
