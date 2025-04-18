package search

import "fmt"

func BinarySearch(values []int, target int) (int, bool) {

	low := 0
	high := len(values) - 1
	mid := (high + low) / 2

	iterations := 0
	for low <= high {
		iterations++
		currentValue := values[mid]
		if currentValue == target {
			fmt.Printf("Found the value %d in %d iterations\n", target, iterations)
			return mid, true
		}
		if currentValue < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
		mid = (high + low) / 2
	}
	fmt.Printf("Could not find the value %d in %d iterations\n", target, iterations)
	return -1, false
}

type Person struct {
	Name        string
	Target      bool
	Connections []string
}

func BreadthFirstSearch(people map[string]Person, start string) bool {

	queue := []Person{}
	visited := map[string]bool{}
	firstPerson := people[start]

	for _, v := range firstPerson.Connections {
		queue = append(queue, people[v])
	}

	for len(queue) > 0 {

		// mark visited
		visited[queue[0].Name] = true

		// check goal
		if queue[0].Target {
			return true
		}

		// add connection to queue
		if len(queue[0].Connections) > 0 {
			for _, v := range queue[0].Connections {
				if !visited[v] {
					queue = append(queue, people[v])
				}

			}
		}

		// pop person off
		queue = queue[1:]
	}

	return false
}
