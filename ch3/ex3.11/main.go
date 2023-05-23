package main

import (
    "bytes"
    "strings"
    "fmt"
)

func main() {
    fmt.Println(comma("-1145141919810"))
    fmt.Println(comma("+114514.1919810"))
}

func comma(s string) string {
    b := bytes.Buffer{}

    if strings.HasPrefix(s, "-") || strings.HasPrefix(s, "-") {
        b.WriteString(s[:1])
        s = s[1:]
    }

    if period := strings.LastIndex(s, "."); period != -1 {
        write(s[:period], &b)
        b.WriteString(".")
        b.WriteString(s[period+1:])
    } else {
        write(s, &b)
    }

    return b.String()
}

func write(s string, b *bytes.Buffer) {
    n := len(s)
    w := 3
    i := n % w

    if i > 0 {
        fmt.Fprintf(b, "%s,", s[:i])
    }

    for i < n {
        fmt.Fprint(b, s[i:i+w])
        if i + w < n {
            b.WriteString(",")
        }

        i += w
    }
}
