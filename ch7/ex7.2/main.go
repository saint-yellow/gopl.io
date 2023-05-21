package main

import (
	"bytes"
	"fmt"
)

func main() {
	data := []byte("saint-yellow")
	buffer := bytes.Buffer{}
	writer, n := countingWriter(&buffer)
	writer.Write(data)
	fmt.Println(writer)
	fmt.Println(*n)
}
