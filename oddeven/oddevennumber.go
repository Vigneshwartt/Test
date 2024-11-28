package oddeven

import "fmt"

func OddEvenNumbers() {
	numbers := make(chan int)
	ouput := make(chan struct{})

	go oddevenNumbers(numbers)
	go printOutput(numbers, ouput)
	<-ouput
}

func oddevenNumbers(inputNumber chan int) {
	var size int
	var userInput int

	fmt.Println("Enter the size")
	_, err := fmt.Scanln(&size)
	if err != nil {
		fmt.Println("Error occured")
		return
	}

	if size == 0 {
		close(inputNumber)
		fmt.Println("you have to enter the number above zero")
	}

	for i := 0; i < size; i++ {
		fmt.Printf("Enter number %d: ", i+1)

		_, err = fmt.Scanln(&userInput)
		if err != nil {
			fmt.Printf("Error: %s", err)
			return
		}

		if userInput == 0 {
			close(inputNumber)
			fmt.Println("you have to enter the number above zero")
			return
		}

		inputNumber <- userInput
	}

	close(inputNumber)
}

func printOutput(in chan int, empty chan struct{}) {
	temporayVariable := []int{}

	for j := range in {
		temporayVariable = append(temporayVariable, j)
	}

	for _, i := range temporayVariable {
		if i%2 == 0 {
			fmt.Println("Even number\t:", i)
		} else {
			fmt.Println("Odd number\t:", i)
		}
	}

	empty <- struct{}{}
}
