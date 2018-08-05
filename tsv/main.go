package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func main() {
	var fp *os.File
	if len(os.Args) < 2 {
		fp = os.Stdin
	} else {
		var err error
		fp, err = os.Open(os.Args[1])
		if err != nil {
			panic(err)
		}
		defer fp.Close()
	}

	reader := csv.NewReader(fp)
	reader.Comma = '\t'
	reader.LazyQuotes = true

	for {
		line, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				panic(err)
			}
		}
		fmt.Println(line)
	}
}
