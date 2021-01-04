package main

import "fmt"

func main() {
	fmt.Print("Enter the temperature in F: ")
	var far float64
	var celcius float64
	fmt.Scanf("%f", &far)

	celcius = (far - 32) * 5 / 9

	fmt.Println(celcius)
}
