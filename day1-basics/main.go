package main

import "fmt"

func main() {
	var a int = 10
	name := "ashish"
	var count int
	var done bool
	const maxRetries int = 5

	price := 9.5
	total := float64(price) * float64(a)
	fmt.Println(a, name, count, done, maxRetries)
	fmt.Println("Total price is:", total)
}
