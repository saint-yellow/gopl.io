package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "ex2.2: %v\n", err)
			os.Exit(1)
		}

		m, f := meter(t), foot(t)
		fmt.Printf("%s = %s, %s = %s\n", m, mTof(m), f, fTom(f))
	}
}
