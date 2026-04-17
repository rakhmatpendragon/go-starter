package main

import "fmt"

func main() {
	/* If else condition */
	var point = 7

	if point == 10 {
		fmt.Println("lulus dengan nilai sempurna")
	} else if point > 5 {
		fmt.Println("lulus")
	} else if point < 5 {
		fmt.Println("Gagal")
	} else {
		fmt.Printf("Silahkah diulang. nilai anda %d\n", point)
	}

	/* Variable temporary */
	var poin = 8840.0

	if percent := poin / 100; percent >= 100 {
		fmt.Printf("%.1f%s perfect!\n", percent, "%")
	} else if percent >= 70 {
		fmt.Printf("%.1f%s good\n", percent, "%")
	} else {
		fmt.Printf("%.1f%s not bad\n", percent, "%")
	}

	/*
		Switch case condition.
		In go no need to use break.
		If condition have been fullfilled it will break automatically.
	*/
	var pnt = 6

	switch pnt {
	case 8:
		fmt.Println("perfect")
	case 7:
		fmt.Println("awesome")
	case 6, 5:
		fmt.Println("not bad")
	default:
		{
			fmt.Println("fail")
			fmt.Println("please try again")
			fmt.Println("you can do better")
		}
	}

	/*
		fallthrough will continue to go next case without consider the condition is true or not but only one case.
		fallthrough can put more than one.
	*/
	switch {
	case point == 8:
		fmt.Println("perfect")
	case (point < 8) && (point > 6):
		fmt.Println("awesome")
		fallthrough
	case point < 5:
		fmt.Println("you need to try again")
		fallthrough
	case point < 4:
		fmt.Println("you need to study more")
	case point < 3:
		fmt.Println("you need to take this more serious")
	default:
		{
			fmt.Println("fail")
			fmt.Println("you need to learn more")
		}
	}
}
