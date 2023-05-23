package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma(os.Args[i]))
	}
}


func comma(s string) string {
    n := len(s)
    i := n % 3
    b := bytes.Buffer{}

    if i > 0 {
        fmt.Fprintf(&b, "%s,", s[:i])
    }
    for ; i < n; i += 3 {
        fmt.Fprint(&b, s[i:i+3])
        if i + 3 < n {
            b.WriteRune(',')
        }
    }

    return b.String()
}
