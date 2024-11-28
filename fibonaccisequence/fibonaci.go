package fibonaccisequence

import "fmt"

func FibbonaciSequence() {
	numbers := make(chan int)

	go fibonacciSeries(numbers)

	for a := range numbers {
		fmt.Println("Fibonacii Number", a)
	}
}

func fibonacciSeries(a chan int) {
	var size int
	x, y := 0, 1

	fmt.Println("Enter the size")
	_, err := fmt.Scanln(&size)
	if err != nil {
		fmt.Println("Error occured", err)
		return
	}

	if size == 0 {
		close(a)
		fmt.Println("Enter the number above zero")
		return
	}

	fmt.Println("Fibonacii Number", x, "\n", "Fibonacii Number", y)

	for i := 2; i < size; i++ {
		nextvalue := x + y
		a <- nextvalue

		x = y
		y = nextvalue
	}
	close(a)
}
