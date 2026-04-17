package main

import "fmt"

func main() {
	const firstName string = "Rakhmat"
	const lastName = "Ibrahim"

	fmt.Print("halo ", firstName, "!\n")
	fmt.Print("nice to meet you ", lastName, "!\n")

	/* Different format Print and Println */
	fmt.Println("Rakhmat Ibrahim")
	fmt.Println("Rakhmat", "Ibrahim")

	fmt.Print("Rakhmat Ibrahim\n")
	fmt.Print("Rakhmat ", "Ibrahim\n")
	fmt.Print("Rakhmat", " ", "Ibrahim\n")

	/* Multi Constant */
	const (
		square         = "kotak"
		isToday  bool  = true
		numeric  uint8 = 1
		floatNum       = 2.2
	)

	const (
		a = "constant"
		b
	)

	const (
		today string = "senin"
		sekarang
		isToday2 = true
	)

	const satu, dua = 1, 2
	const three, four, string = "tiga", "empat"
}
