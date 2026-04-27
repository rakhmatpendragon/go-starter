package main

import (
	"fmt"
	"strings"
)

func main() {
	var avg = calculate(3, 4, 3, 5, 4, 3, 3, 5, 5, 3)
	var msg = fmt.Sprintf("Average : %.2f", avg)
	fmt.Println(msg)

	/* use data slice */
	fmt.Println("\n=== use data slice ===")
	var number = []int{3, 4, 3, 5, 4, 3, 3, 5, 5, 3}
	var avgNum = calculate(number...)
	var msgNum = fmt.Sprintf("Average Numb : %.2f", avgNum)

	fmt.Printf(msgNum)

	/* parameter normal and variadic */
	fmt.Println("\n\n=== parameter normal and variadic ===")

	myHobbies("Ibrahim", "watching", "eating")
}

func calculate(numbers ...int) float64 {
	var total int = 0
	for _, number := range numbers {
		total += number
	}

	var avg = float64(total) / float64(len(numbers))
	return avg
}

func myHobbies(name string, hobbies ...string) {
	var hobbiesAsString = strings.Join(hobbies, ", ")

	fmt.Printf("Hello, my name is: %s\n", name)
	fmt.Printf("My hobbies are: %s\n", hobbiesAsString)
}
