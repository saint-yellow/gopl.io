package main

import (
	"bufio"
	"bytes"
	"io"
)

type wordCounter struct {
	writer  io.Writer
	written int64
	words   int64
}

func (c *wordCounter) Write(p []byte) (int, error) {
	s := bufio.NewScanner(bytes.NewBuffer(p))
	s.Split(bufio.ScanWords)
	for s.Scan() {
		c.words += 1
	}
	c.written = int64(len(p))
	return len(p), nil
}

func countingWriter(w io.Writer) (io.Writer, *int64) {
	c := &wordCounter{
		writer: w,
	}
	return c, &c.written
}
