package main

import (
	"fmt"
    "unicode/utf8"
)

func main() {
    s := "上海自来水"
    fmt.Println(s)

    p := []byte(s)
    reverse(p)
    fmt.Println(string(p))
}

func reverse(p []byte) {
    for i, j := 0, len(p)-1; i < j; i, j = i+1, j-1 {
        p[i], p[j] = p[j], p[i]
    }

    // try to decode according to utf8, then fix error
    i := 0
    for i < len(p) {
        var tryTwo, tryThree, tryFour bool
        for {
            r, s := utf8.DecodeRune(p[i:])
            if r != utf8.RuneError {
                i += s
                break
            } else {
                // try two byte length, swap two bytes
                if !tryTwo {
                    tryTwo = true
                    p[i], p[i+1] = p[i+1], p[i]
                    continue
                }

                // try three byte length, swap three bytes
                if !tryThree {
                    // cancel tryTwo side effect
                    p[i], p[i+1] = p[i+1], p[i]
                    tryThree = true
                    p[i], p[i+2] = p[i+2], p[i]
                    continue
                }

                // try four byte length, swap four bytes
                if !tryFour {
                    // cancel tryThree side effect
                    p[i], p[i+1], p[i+2] = p[i+2], p[i+1], p[i]

                    tryFour = true
                    p[i], p[i+1], p[i+2], p[i+3] = p[i+3], p[i+2], p[i+1], p[i]
                    continue
                }
            }
        }
    }
}
