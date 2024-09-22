package main

// Chapter 2
func selectionSort(unsortedSlc []int) []int {
	var min int
	sortedSlc := []int{}
	for i := 0; i < len(unsortedSlc); i++ {
		//for _, value := range unsortedSlc {
		min = 999999999
		minIndex := 0
		for j, v := range unsortedSlc {
			if v < min {
				min = v
				minIndex = j
			}
		}
		unsortedSlc[minIndex] = 9999999
		sortedSlc = append(sortedSlc, min)
	}
	return sortedSlc
}
