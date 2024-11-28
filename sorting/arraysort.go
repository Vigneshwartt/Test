package sorting

import (
	"fmt"
)

func SortingNumbers() {
	var size int

	fmt.Println("Enter how many numbers you want to sort:")
	_, err := fmt.Scanln(&size)
	if err != nil {
		fmt.Println("Enter a proper number")
		return
	}

	if size == 0 {
		fmt.Println("Enter the number above zero")
		return
	}

	numbers := make([]int, size)

	for i := 0; i < size; i++ {
		fmt.Printf("Enter number %d: ", i+1)

		_, err = fmt.Scanln(&numbers[i])
		if err != nil {
			fmt.Println("Enter a proper number", err)
			return
		}
	}

	sortedChannel := make(chan []int)

	go sortNumbers(numbers, sortedChannel)

	sortedNumbers := <-sortedChannel
	fmt.Println("Sorted numbers:", sortedNumbers)
}

func sortNumbers(arr []int, in chan []int) {
	numberSorting(arr)
	in <- arr
	close(in)
}

func numberSorting(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}
