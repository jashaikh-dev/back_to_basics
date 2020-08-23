package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, val := range arr {
		if val%2 == 0 {
			fmt.Println(val, "is even")
		}
	}
}

// `> go run main.go` to execute the above code
