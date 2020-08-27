package main

import (
	"fmt"
	"strconv"
)

func main() {
	var ts float64
	var tm float64
	fmt.Println("Enter Out of Scores")
	fmt.Scan(&tm)
	fmt.Println("Enter Total Scores")
	fmt.Scan(&ts)
	p := (ts / tm) * 100
	if p > 60.00 && p <= 70.00 {
		print("Grade:  D || Percentage: " + floatToString(p))
	} else if p > 70.0 && p <= 80.00 {
		print("Grade: C || Percentage: " + floatToString(p))
	} else if p > 80.0 && p <= 90.00 {
		print("Grade: B || Percentage: " + floatToString(p))
	} else if p > 90.00 {
		print("Grade: A || Percentage: " + floatToString(p))
	} else {
		print("Failed || Percentage: " + floatToString(p))
	}
}

func print(content string) {
	fmt.Println(content)
}

func floatToString(val float64) string {
	return strconv.FormatFloat(val, 'f', 2, 64)
}
