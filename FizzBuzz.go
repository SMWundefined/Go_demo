package main

import (
	"fmt"
)

func main() {
	var limit int
	i := 0
	fmt.Println("Enter the limit")
	fmt.Scanf("%d", &limit)
	for i < limit {
		if i%5&i%3 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}
