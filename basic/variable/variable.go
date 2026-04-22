package main

import "fmt"

func main() {
	var firstName string = "Rakhmat"
	var lastName string
	lastName = "Ibrahim"
	middleName := "Malik"

	fmt.Printf("Hello Rakhmat Malik Ibrahim!\n")
	fmt.Printf("Hello %s %s %s!\n", firstName, middleName, lastName)
	fmt.Println("Hello", firstName, middleName, lastName+"!")

	middleName = "Maliq"

	fmt.Println("Hello", firstName, middleName, lastName+"!")

	/* Declaration multiple variable */
	var first, second, third string
	first, second, third = "satu", "dua", "tiga"
	var fourth, fifth, sixth string = "empat", "lima", "enam"
	seventh, eight, ninth := "tujuh", "delapan", "sembilan"
	one, isFriday, twoPointTwo, say := 1, true, 2.2, "hello"

	fmt.Println("\n=== Multiple Variable ===")
	fmt.Println(first, second, third)
	fmt.Println(fourth, fifth, sixth)
	fmt.Println(seventh, eight, ninth)
	fmt.Println(one, isFriday, twoPointTwo, say)

	/*
		Variable Underscore.
		In Go all variable must be use.
		In Go if variable have been declare but not been used it will be error, except Variable Underscore.
	*/
	_ = "Learning Golang"
	_ = "Golang is easy"
	name, _ := "Rkh", "Mlq"
	fmt.Println("\n=== Variable Underscore ===")
	fmt.Println(name)

	/* Variable keyword new */
	nm := new(string)

	fmt.Println("\n=== Variable keyword new ===")
	fmt.Println(nm)
	fmt.Println(*nm)
}
