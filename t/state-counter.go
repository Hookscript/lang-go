package main

import "fmt"

func Hook(n *int) {
	*n++
	fmt.Print(*n)
}
