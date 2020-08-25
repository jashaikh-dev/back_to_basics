package main

import "fmt"

func main() {
	var num int
	fmt.Println("Enter a number: ")
	_, err := fmt.Scan(&num)
	if err != nil {
		fmt.Println("Something went wrong while reading input")
	}

	if num < 0 {
		fmt.Println(num, "is negative")
	} else if num > 0 {
		fmt.Println(num, "is positive")
	} else {
		fmt.Println(num, "is zero")
	}
}
