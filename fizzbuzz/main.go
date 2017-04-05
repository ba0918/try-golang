package main

import (
	"flag"
	"fmt"
	"strconv"
)

func main() {
	var start, end int
	flag.IntVar(&start, "s", 1, "start value")
	flag.IntVar(&end, "e", 1, "end value")
	flag.Parse()
	fizzbuzzRange(start, end)
}

func fizzbuzzRange(start int, end int) {
	for i := start; i <= end; i++ {
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
