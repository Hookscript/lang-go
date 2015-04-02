package main

import (
	"fmt"
	"net/http"
	"os"
)

func Hook(r *http.Request) {
	resource := r.FormValue("resource")
	switch resource {
	case "cpu":
		for {
			// do nothing
		}
	case "mem":
		xs := []int{0}
		for {
			xs = append(xs, xs...)
		}
	case "disk":
		f, err := os.Create("junk")
		MaybePanic(err)
		for {
			_, err = f.WriteString("junk\n")
			MaybePanic(err)
		}
	case "output":
		for i := 0; i < 1000000; i++ {
			_, err := fmt.Println("junk")
			MaybePanic(err)
		}
	}
}

func MaybePanic(err error) {
	if err != nil {
		panic(err)
	}
}
