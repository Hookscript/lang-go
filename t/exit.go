package main

import (
	"fmt"
	"net/http"
	"os"
)

func Hook(r *http.Request) {
	switch r.FormValue("try") {
	case "1":
		panic("Try 1")
	case "2":
		fmt.Fprintln(os.Stderr, "Try 2")
		os.Exit(1)
	}
}
