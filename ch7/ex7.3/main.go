package main

import "fmt"

func main() {
	var t *tree
	values := []int{
		1, 9, 2, 8, 3, 7, 4, 6, 5,
	}
	for _, v := range values {
		t = add(t, v)
	}

	fmt.Println(t)
}
