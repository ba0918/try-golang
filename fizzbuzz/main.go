package main

import (
	"fmt"
)

func main() {
	fizzbuzz(1, 100)
}

func fizzbuzz(start int, end int) {
	for i := start; i < end; i++ {
		a := i%3 == 0
		b := i%5 == 0
		switch {
		case a && b:
			fmt.Println("fizzbuzz")
		case a:
			fmt.Println("fizz")
		case b:
			fmt.Println("buzz")
		default:
			fmt.Printf("%d\n", i)
		}
	}
}
