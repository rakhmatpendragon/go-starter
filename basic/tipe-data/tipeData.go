package main

import "fmt"

func main() {
	/* Numeric non-decimal data type */
	var positiveNumber uint8 = 89
	var negativeNumber = -1243423644

	fmt.Println("=== Non-Decimal Number ===")
	fmt.Printf("bilangan positif: %d\n", positiveNumber)
	fmt.Printf("bilangan negatif: %d\n", negativeNumber)

	/* Numeric decimal data type */
	var decimalNumber = 2.62

	fmt.Println("\n=== Decimal Number ===")
	fmt.Printf("bilangan decimal: %f\n", decimalNumber)
	fmt.Printf("bilangan decimal: %3.f\n", decimalNumber)

	/* Boolean data type */
	var exist bool = true

	fmt.Println("\n=== Boolean ===")
	fmt.Printf("exist? %t \n", exist)

	/* String data type */
	var message string = "Halo"
	var msg string = `Nama saya "Rakhmat Malik".
Salam kenal.
	Mari belajar "Golang"`

	fmt.Println("\n=== String ===")
	fmt.Printf("message: %s \n", message)
	fmt.Println(msg)
}
