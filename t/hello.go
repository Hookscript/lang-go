package main

import (
	"fmt"
	"net/http"
)

func Hook(r *http.Request) {
	whom := r.FormValue("whom")
	if whom == "" {
		whom = "world"
	}
	fmt.Printf("Hello, %s!\n", whom)
}
