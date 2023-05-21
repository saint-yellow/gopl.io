package main

import (
	"bufio"
	"bytes"
)

type wordCounter int

func (c *wordCounter) Write(p []byte) (int, error) {
	s := bufio.NewScanner(bytes.NewBuffer(p))
	s.Split(bufio.ScanWords)
	for s.Scan() {
		*c += 1
	}
	return len(p), nil
}

type lineCounter int

func (c *lineCounter) Write(p []byte) (int, error) {
	s := bufio.NewScanner(bytes.NewBuffer(p))
	s.Split(bufio.ScanLines)
	for s.Scan() {
		*c += 1
	}
	return len(p), nil
}
