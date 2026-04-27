package main

import (
	"fmt"
	"math/rand/v2"
	"strings"
)

func main() {
	var names = []string{"Rakhmat", "Malik"}
	printMessage("halo", names)

	/* Function with return value */
	fmt.Println("\n=== Function return value ===")

	var randomValue int

	randomValue = randomWithRange(2, 10)
	fmt.Println("Random number :", randomValue)

	randomValue = randomWithRange(2, 10)
	fmt.Println("Random number :", randomValue)

	randomValue = randomWithRange(2, 10)
	fmt.Println("Random number :", randomValue)

	/* Keyword return for stop processing function */
	divideNumber(10, 2)
	divideNumber(4, 0)
	divideNumber(8, -4)
}

func printMessage(message string, arr []string) {
	var nameString = strings.Join(arr, " ")
	fmt.Println(message, nameString)
}

func randomWithRange(min, max int) int {
	var value = rand.Int()%(max-min+1) + min
	return value
}

func divideNumber(m, n int) {
	if n == 0 {
		fmt.Printf("Invalid divider. %d cannot divided by %d\n", m, n)
		return
	}

	var res = m / n
	fmt.Printf("%d / %d = %d\n", m, n, res)
}
