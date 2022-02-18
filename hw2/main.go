package main

import (
	"fmt"

	"github.com/zethuman/go-spring/week2/functions"
)

func main() {
	functions.Breaks("Functions")
	functions.Breaks("Task 1")
	fmt.Printf("\n%10s | %8s | %8s | %8s | %8s\n",
		"Name", "Score 1", "Score 2", "Score 3", "Average")
	functions.Breaks("")
	functions.ScoreSummary("Jermaine", 95.4, 82.3, 74.6)
	functions.ScoreSummary("Jessie", 79.3, 99.1, 82.5)
	functions.ScoreSummary("Lamar", 82.2, 95.4, 77.6)

	functions.Breaks("Task 2")

	var sum float64
	sum += functions.Perimeter(8.2, 10)
	sum += functions.Perimeter(5, 5.4)
	sum += functions.Perimeter(6.2, 4.5)
	fmt.Printf("\nYou'll need %f meters of fencing\n", sum)

	functions.Breaks("Task 3")

	p, err := functions.SmartPerimeter(8.2, -10)
	if err != nil {
		fmt.Printf("\n%s\n", err.Error())
		return
	}
	fmt.Println(p)
}
