package main

import "fmt"

func main() {
	var myProduct = []string{"HP", "Laptop", "TV", "PwrBnk"}

	var competitor = myProduct[0:3]
	var mitra = myProduct[1:4]

	var agen = competitor[1:2]
	var store = mitra[0:1]

	var market = myProduct[:]
	var liteMarket = myProduct[:2]
	var microMarket = myProduct[2:]

	fmt.Println("=== Before Update ===")
	fmt.Println(myProduct)
	fmt.Println(competitor)
	fmt.Println(mitra)
	fmt.Println(agen)
	fmt.Println(store)
	fmt.Println(market)
	fmt.Println(liteMarket)
	fmt.Println(microMarket)

	/* update value Laptop into AC */
	store[0] = "AC"

	fmt.Println("\n=== After Update ===")
	fmt.Println(myProduct)
	fmt.Println(competitor)
	fmt.Println(mitra)
	fmt.Println(agen)
	fmt.Println(store)
	fmt.Println(market)
	fmt.Println(liteMarket)
	fmt.Println(microMarket)

	/* len and cap */
	fmt.Println("\n=== About Len() and Cap() ===")
	fmt.Println("length", len(myProduct))
	fmt.Println("cap", cap(myProduct))

	fmt.Println("length", len(competitor))
	fmt.Println("cap", cap(competitor))

	fmt.Println("length", len(mitra))
	fmt.Println("cap", cap(mitra))

	/* append */
	fmt.Println("\n=== About append() ===")

	var bigStore = append(store, "Air Purifier")

	fmt.Println(myProduct)
	fmt.Println(store)
	fmt.Println(bigStore)

	/* copy */
	fmt.Println("\n=== About copy() ===")

	dst := make([]string, 3)
	src := []string{"watermelon", "pinnable", "apple", "banana"}
	n := copy(dst, src)

	fmt.Println(dst)
	fmt.Println(src)
	fmt.Println(n)

	veggie := []string{"carrot", "spinach", "potato"}
	fruit := []string{"melon", "durian"}
	veg := copy(veggie, fruit)

	fmt.Println(veggie)
	fmt.Println(veg)

	/* access elmnt slice with 3 index */
	var item = 
}
