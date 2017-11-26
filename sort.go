package main

import "fmt"

func bubbleSort(arr []int) []int {
	swapCounter := 0
	maxIndex := len(arr) - 1

	for i := 0; i < maxIndex; i++ {
		if arr[i] > arr[i + 1] {
			swapCounter += 1
			arr[i], arr[i + 1] = arr[i + 1], arr[i]
		}

		if i == 3 && swapCounter != 0 {
			i = -1
			swapCounter = 0
		}
	}

	return arr
}

func main() {
	arr := []int{ 10, 2, 9, 4, 1, 5, 3, 6, 8, 7 }
	
	fmt.Println(bubbleSort(arr))
}
