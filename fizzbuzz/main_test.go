package main

import (
	"testing"
)

type testCase struct {
	in  int
	out string
}

var tests = []testCase{
	{1, "1"},
	{3, "fizz"},
	{5, "buzz"},
	{11, "11"},
	{15, "fizzbuzz"},
	{75, "fizzbuzz"},
	{98, "98"},
	{99, "fizz"},
	{100, "buzz"},
}

func TestFizzbuzz(t *testing.T) {
	for i := range tests {
		test := &tests[i]
		out := fizzbuzz(test.in)
		if test.out != out {
			t.Errorf("fizzbuzz(%d) = %q want %q", test.in, out, test.out)
		}
	}
}
