package main

import "fmt"

func main() {
	/* Implementation */
	fmt.Println("=== Implementation ===")
	var price map[string]int

	/*
		default value of map is nil
		initialize like this to prevent error
	*/
	price = map[string]int{}

	price["mouse"] = 200000
	price["bottle"] = 300000

	fmt.Println("mouse price is", price["mouse"])
	fmt.Println("pencil price is", price["pencil"])

	/* horizontal style */
	var snack = map[string]int{"biscuit": 5000, "cereal": 10000}

	/* vertical style */
	var drink = map[string]int{
		"milk":   20000,
		"coffee": 25000,
		"soda":   15000,
		"water":  7000,
	}

	var food = map[string]int{}
	var desert = make(map[string]int)
	var appetizer = *new(map[string]int)

	fmt.Println("biscuit price is", snack["biscuit"])
	fmt.Println("coffe price is", drink["coffee"])
	fmt.Println(snack)
	fmt.Println(drink)
	fmt.Println(food)
	fmt.Println(desert)
	fmt.Println(appetizer)

	/* for - range on map*/
	fmt.Println("\n=== Implementation for-range on map===")

	for key, val := range drink {
		fmt.Println(key, " \t:", val)
	}

}
