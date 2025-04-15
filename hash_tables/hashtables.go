package hashtables

import (
	"container/list"

	"github.com/tenfyzhong/cityhash"
)

type HashMap struct {
	values []list.List
}

func NewHashMap(size int) *HashMap {
	values := make([]list.List, size)
	for i := 0; i < size; i++ {
		values[i] = *list.New()
	}
	return &HashMap{values: values}
}

func (h *HashMap) Get(key string) (string, bool) {
	hashedValue := cityhash.CityHash64([]byte(key))
	index := hashedValue % uint64(len(h.values))

	for element := h.values[index].Front(); element != nil; element = element.Next() {
		if element.Value == nil {
			return "", false
		}
		stringArray := element.Value.([]string)
		if stringArray[0] == key {
			return stringArray[1], true
		}
	}
	return "", false

}

func (h *HashMap) Set(key, value string) {

	hashedValue := cityhash.CityHash64([]byte(key))

	index := hashedValue % uint64(len(h.values))

	// create a linked list if not null
	h.values[index].PushFront([]string{key, value})

}
