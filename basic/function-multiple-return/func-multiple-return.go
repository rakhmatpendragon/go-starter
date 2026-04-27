package main

import (
	"fmt"
	"math"
)

func calculate(d float64) (float64, float64) {
	var area = math.Pi * math.Pow(d/2, 2)
	var circumference = math.Pi * d

	return area, circumference
}

func squareCalc(s float64) (area float64, circumference float64) {
	area = s * s
	circumference = 4 * s

	return area, circumference
}

func main() {
	var diameter, side float64 = 20, 10
	var area, circumference = calculate(diameter)
	var squareArea, squareCircumference = squareCalc(side)

	fmt.Printf("Circle area\t\t: %.2f \n", area)
	fmt.Printf("Circle circumference\t\t: %.2f \n", circumference)
	fmt.Printf("Square area\t\t: %.2f \n", squareArea)
	fmt.Printf("Square circumference\t\t: %.2f \n", squareCircumference)
}
