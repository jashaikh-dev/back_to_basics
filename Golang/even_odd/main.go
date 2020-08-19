package main

import "fmt"

func main() {
	num := 20
	if num%2 == 0 {
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}
}

// `> go run main.go` to execute the above code
