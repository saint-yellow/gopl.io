package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	file := os.Args[1]
	data, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ex7.1: %v\n", err)
		os.Exit(1)
	}

	var lc lineCounter
	lc.Write(data)
	var wc wordCounter
	wc.Write(data)

	fmt.Printf("%s: %d line(s), %d word(s)\n", file, lc, wc)
}
