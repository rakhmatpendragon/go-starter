package main

import "fmt"

func main() {
	var names [4]string

	names[0] = "learn"
	names[1] = "about"
	names[2] = "golang"
	names[3] = "programming"

	fmt.Println(names[0], names[1], names[2], names[3])

	/* Initialization array with num */
	fmt.Println("\n=== Initialization array with num ===")
	var animals = [4]string{"chicken", "duck", "fish", "bird"}

	var transportations = [4]string{
		"car",
		"train",
		"ship",
		"plane",
	}

	fmt.Println("The name of animals", animals)
	fmt.Println("the name of transportations", transportations)
	fmt.Println("The length of animal is", len(animals), "and transportation is", len(transportations))

	/* Initialization array without num */
	fmt.Println("\n=== Initialization array without num ===")
	var number = [...]int{1, 2, 3, 4, 5}

	fmt.Println("The number is", number)
	fmt.Println("Length of number", len(number))

	/* Array multi dimention */
	fmt.Println("\n=== Array multi dimention ===")
	var array1 = [2][3]int{[3]int{1, 2, 3}, [3]int{3, 4, 5}}
	var array2 = [2][3]int{{3, 4, 5}, {6, 7, 8}}

	fmt.Println("array1", array1)
	fmt.Println("array2", array2)

	/* Loop and Array */
	fmt.Println("\n=== Loop and Array ===")

	for i := 0; i < len(animals); i++ {
		fmt.Printf("Element %d : %s\n", i, animals[i])
	}

	/* Loop range and array */
	fmt.Println("\n=== Loop range and array ===")

	for j, transportation := range transportations {
		fmt.Printf("Transport %d : %s\n", j, transportation)
	}

	/* Loop range, array and underscore */
	fmt.Println("\n=== Loop range, array and underscore ===")

	for _, animal := range animals {
		fmt.Printf("This animal is %s\n", animal)
	}

	for k := range animals {
		fmt.Printf("Indeks no %d\n", k)
	}

	/* Array and make keyword */
	fmt.Println("\n=== Array and make keyword ===")

	var emotion = make([]string, 2)
	emotion[0] = "happy"
	emotion[1] = "angry"

	fmt.Println(emotion)
}
