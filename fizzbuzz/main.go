package main

import (
	"fmt"
	"strconv"
)

func main() {
	fizzbuzzRange(1, 100)
}

func fizzbuzzRange(start int, end int) {
	for i := start; i < end; i++ {
		fmt.Println(fizzbuzz(i))
	}
}

func fizzbuzz(n int) string {
	a := n%3 == 0
	b := n%5 == 0
	switch {
	case a && b:
		return "fizzbuzz"
	case a:
		return "fizz"
	case b:
		return "buzz"
	default:
		return strconv.Itoa(n)
	}
}
