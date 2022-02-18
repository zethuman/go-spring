package functions

import "fmt"

func ScoreSummary(studentName string, params ...float64) {
	avg := average(params...)
	fmt.Printf("%10s | %8.2f | %8.2f | %8.2f | %8.2f\n",
		studentName, params[0], params[1], params[2], avg)
}

func average(params ...float64) float64 {
	var sum float64
	for _, v := range params {
		sum += v
	}
	return sum / float64(len(params))
}
