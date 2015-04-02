package main

import (
	"fmt"
	"net/http"
)

type State *map[string]string

func Hook(r *http.Request, state State) {
	value := ""
	register := r.FormValue("register")
	if *state == nil {
		*state = make(map[string]string)
	}
	registers := *state

	switch r.Method {
	case "GET":
		var ok bool
		value, ok = registers[register]
		if !ok {
			value = fmt.Sprintf("unknown register: %s", register)
		}
	case "POST":
		value = r.FormValue("value")
		if value == "" {
			value = "default value"
		}
		if register == "death" {
			panic("I am dead")
		}
		registers[register] = value
	}

	fmt.Print(value)
}
