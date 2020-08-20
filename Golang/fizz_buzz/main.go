package main

import "fmt"

func main() {
	num := 20
	for i := 1; i <= num; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FIZZBUZZ")
		} else if i%3 == 0 {
			fmt.Println("FIZZ")
		} else if i%5 == 0 {
			fmt.Println("BUZZ")
		} else {
			fmt.Println("-")
		}
	}
}

// `> go run main.go` to execute the above code
