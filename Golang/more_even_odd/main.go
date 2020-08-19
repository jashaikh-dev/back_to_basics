package main

import "fmt"

func main() {
	num := 20
	for i := 1; i <= num; i++ {
		if i%2 == 0 {
			fmt.Println(i, "is even")
		} else {
			fmt.Println(i, "is odd")
		}
	}
}

// `> go run main.go` to execute the above code
