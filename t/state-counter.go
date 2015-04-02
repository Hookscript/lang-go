package main

import "fmt"

type Counter int

func Hook(n *Counter) {
	*n++
	fmt.Print(*n)
}
