package main

import (
	"fmt"

	hashtables "github.com/KylerJacobson/grokking-algorithms/hash_tables"
)

func main() {

	x := hashtables.NewHashMap(1024)

	x.Set("name", "Kyler")

	name, ok := x.Get("name")
	if !ok {
		fmt.Println("Key not found")
	} else {
		fmt.Println(name)
	}

	age, ok := x.Get("age")
	if !ok {
		fmt.Println("Key not found")
	} else {
		fmt.Println(age)
	}

}
