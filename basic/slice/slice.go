package main

import "fmt"

func main() {
	var myProduct = []string{"HP", "Laptop", "TV", "PwrBnk"}

	var competitor = myProduct[0:3]
	var mitra = myProduct[1:4]

	var agen = competitor[1:2]
	var store = mitra[0:1]

	fmt.Println("=== Before Update ===")
	fmt.Println(myProduct)
	fmt.Println(competitor)
	fmt.Println(mitra)
	fmt.Println(agen)
	fmt.Println(store)

	// update value Laptop into AC
	store[0] = "AC"

	fmt.Println("\n=== After Update ===")
	fmt.Println(myProduct)
	fmt.Println(competitor)
	fmt.Println(mitra)
	fmt.Println(agen)
	fmt.Println(store)
}
