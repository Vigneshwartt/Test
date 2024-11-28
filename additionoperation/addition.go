package additionoperation

import (
	"fmt"
	"sync"
)

func AdditionOfNumbers() {
	number := make(chan int)
	addition := make(chan int)
	var a sync.WaitGroup
	a.Add(3)

	go getNumbers(number, &a)
	go receiveData(number, addition, &a)
	go printOutput(addition, &a)

	a.Wait()
}

func getNumbers(in chan int, wg *sync.WaitGroup) {
	var size int
	var userInput int

	fmt.Println("Enter how many numbers you want:")
	_, err := fmt.Scanln(&size)
	if err != nil {
		fmt.Println("Entered wrong Number")
		return
	}

	for i := 0; i < size; i++ {
		fmt.Printf("Enter number %d: ", i+1)

		_, err = fmt.Scanln(&userInput)
		if err != nil {
			fmt.Printf("Error: %s", err)
			return
		}

		if userInput == 0 {
			close(in)
			fmt.Println("You have to enter the number above zero")
			return
		}
		in <- userInput

	}

	close(in)
	defer wg.Done()
}

func receiveData(no chan int, out chan int, wg *sync.WaitGroup) {
	var sum int

	for b := range no {
		sum += b
	}

	out <- sum
	defer wg.Done()
	close(out)
}

func printOutput(b chan int, wg *sync.WaitGroup) {
	finalSum := <-b

	fmt.Println("Total sum:", finalSum)
	defer wg.Done()
}
