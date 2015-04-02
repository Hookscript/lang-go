package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func Hook(req *http.Request, res *http.Response) {
	proto := req.FormValue("protocol")
	file := req.FormValue("file")
	url := fmt.Sprintf("%s://storage.googleapis.com/hookscript/%s", proto, file)

	got, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	if got.StatusCode != 200 {
		res.StatusCode = got.StatusCode
		fmt.Printf("Request to %s failed", url)
		return
	}

	_, err = io.Copy(os.Stdout, got.Body)
	if err != nil {
		panic(err)
	}
	got.Body.Close()
}
