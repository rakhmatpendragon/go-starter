package main

import "fmt"

func main() {
	/* Method 1 */
	for i := 0; i < 5; i++ {
		fmt.Println("Number", i)
	}

	/* Method 2 */
	var i = 0

	for i < 5 {
		fmt.Println("No", i)
		i++
	}

	/* Method 3 */
	for {
		fmt.Println("Order", i)

		i++
		if i == 10 {
			break
		}
	}

	/* Method 4 */
	var xs = "1234567891011" // string
	for idx, v := range xs {
		fmt.Println("Index=", idx, "value=", v)
	}

	var ys = [5]int{10, 20, 30, 40, 50} // array
	for _, v := range ys {
		fmt.Println("Value", v)
	}

	var zs = ys[0:2] // slice
	for _, v := range zs {
		fmt.Println("Val", v)
	}

	var kvs = map[byte]int{'a': 0, 'b': 1, 'c': 2} // map
	for k, v := range kvs {
		fmt.Println("key=", k, "value=", v)
	}

	for range kvs {
		fmt.Println("Done")
	}

	for i := range 5 {
		fmt.Println(i)
	}

	/* break and continue */
	fmt.Println("\n*** Break and Continue in looping ***")
	for x := 1; x <= 10; x++ {
		if x%2 == 1 {
			fmt.Println("Continue", x)
			continue
		}

		fmt.Println("On Progress", x)
		if x > 8 {
			fmt.Println("Break", x)
			break
		}

		fmt.Println("Numb", x)
	}

	/* Nested */
	fmt.Println("\n=== Nested Loop ===")
	for k := 0; k < 5; k++ {
		for l := k; l < 5; l++ {
			fmt.Print(l, " ")
		}
		fmt.Println()
	}

	/* Matriks */
	fmt.Println("\n=== Matriks ===")
outerLoop:
	for m := 0; m < 5; m++ {
		for n := 0; n < 5; n++ {
			if m == 3 {
				break outerLoop
			}
			fmt.Print("matriks [", m, "][", n, "]\n")
		}
	}
}
