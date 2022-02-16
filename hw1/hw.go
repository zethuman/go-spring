package main

import "fmt"

func main() {
	fmt.Println("----------------- Homework 1 --------------------")
	fmt.Println("First Task:")
	first()
	fmt.Println("\nSecond Task:")
	second()
	fmt.Println("----------------- End --------------------")
}

func first() {
	var count int = 20
	unitWeight := 0.4
	totalWeight := float64(count) * unitWeight
	// fmt.Printf("%d cans weigh %f kilograms\n", count, totalWeight)
	fmt.Println(count, "cans weigh", totalWeight, "kilograms")
}

func second() {
	var pebbleWeight float64 = 0.1
	var rockWeight float64 = 1.2
	var boulderWeight float64 = 502.4
	var totalWeight float64 = pebbleWeight + rockWeight + boulderWeight
	fmt.Println(totalWeight)
}
