package main

import (
	"fmt"
	"routines/additionoperation"
	"routines/fibonaccisequence"
	"routines/oddeven"
	"routines/sorting"
)

func main() {
	var inputNumber int

	fmt.Println("Here three tasks are there \n 1-Addition of numbers \n 2-Fibonacci numbers\n 3-Odd or Even numbers\n 4-Array Of Numbers")

	_, err := fmt.Scanln(&inputNumber)
	if err != nil {
		fmt.Println("Error ocuured:", err)
		return
	}

	switch inputNumber {
	case 1:
		additionoperation.AdditionOfNumbers()
	case 2:
		fibonaccisequence.FibbonaciSequence()
	case 3:
		oddeven.OddEvenNumbers()
	case 4:
		sorting.SortingNumbers()
	default:
		fmt.Println("You have entered wrong number")
		return
	}
}
