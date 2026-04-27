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

	/* Delete on map */
	fmt.Println("\n=== Implementation delete on map===")
	fmt.Println(len(drink))
	fmt.Println(drink)

	delete(drink, "soda")

	fmt.Println(len(drink))
	fmt.Println(drink)

	/* Check item on map */
	fmt.Println("\n=== Implementation check item on map ===")

	var mineralDrink, isMineralExist = drink["water"]
	var juiceDrink, isJuiceExist = drink["orangeJuice"]

	if isMineralExist {
		fmt.Println("The price of mineral water is", mineralDrink)
	} else {
		fmt.Println("Mineral water not on the list menu")
	}

	if isJuiceExist {
		fmt.Println("The price of mineral water is", juiceDrink)
	} else {
		fmt.Println("Juice not on the list menu")
	}

	/* Slice on map */
	fmt.Println("\n=== Implementation slice on map ===")

	var businessProduct = []map[string]string{
		{"name": "Rakhmat", "product": "onigiri"},
		{"name": "Malik", "product": "burger"},
		{"name": "Ibrahim", "product": "daifuku"},
	}

	for _, product := range businessProduct {
		fmt.Println(product["name"], product["product"])
	}

	var products = []map[string]string{
		{"name": "onigiri", "flavour": "chicken teriyaki"},
		{"id": "ong001", "store": "Jakarta"},
		{"ingredient": "rice, chicken"},
	}

	for ky, vl := range products {
		fmt.Println(ky, " \t:", vl)
	}
}
