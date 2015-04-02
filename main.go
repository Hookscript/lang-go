package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"reflect"
	"runtime"
)

func main() {
	var state interface{}
	var stateData []byte
	var err error
	request := initRequest()
	response := initResponse()

	// handy reflection types
	reqType := reflect.TypeOf(&http.Request{})
	resType := reflect.TypeOf(&http.Response{})

	// build a closure to execute user's Hook function
	hook := reflect.ValueOf(Hook)
	hookType := hook.Type()
	numIn := hookType.NumIn()
	args := make([]reflect.Value, 0, numIn)
	for i := 0; i < numIn; i++ {
		argType := hookType.In(i)
		switch argType.Kind() {
		case reflect.Ptr:
		default:
			panic("All arguments of Hook() must be pointers")
		}

		var argValue reflect.Value
		switch argType {
		case reqType:
			argValue = reflect.ValueOf(request)
		case resType:
			argValue = reflect.ValueOf(response)
		default:
			// assume unknown argument contains state
			switch argType.Kind() {
			case reflect.Ptr:
				argValue = reflect.New(argType.Elem())
				state = argValue.Interface()
			}

			stateData, err = ioutil.ReadFile("state")
			if os.IsNotExist(err) {
				// ignore these errors
			} else if err != nil {
				panic(err)
			} else {
				err := json.Unmarshal(stateData, state)
				maybePanic(err)
			}
		}
		args = append(args, argValue)
	}
	f := func() {
		hook.Call(args)
	}

	// capture stdout while hook executes
	out := captureStdout(f)

	generateResponse(out, response)
	saveState(stateData, state)
}

func initRequest() *http.Request {
	f, err := os.Open("request")
	maybePanic(err)
	r, err := http.ReadRequest(bufio.NewReader(f))
	maybePanic(err)
	return r
}

func initResponse() *http.Response {
	r := new(http.Response)
	r.ProtoMajor = 1
	r.ProtoMinor = 1
	r.StatusCode = 200
	r.Header = make(http.Header)
	r.Header.Set("Content-Type", "text/plain")
	r.Header.Set("Server", "hookscript/go-"+runtime.Version())
	return r
}

// technique from http://stackoverflow.com/a/10476304/174463
func captureStdout(f func()) *bytes.Buffer {
	// connect stdout to a pipe
	oldStdout := os.Stdout
	r, w, err := os.Pipe()
	maybePanic(err)
	os.Stdout = w

	// copy output in a separate goroutine so printing can't block forever
	capturedCh := make(chan *bytes.Buffer)
	go func() {
		buf := new(bytes.Buffer)
		_, err := io.Copy(buf, r)
		maybePanic(err)
		capturedCh <- buf
	}()

	f()

	// restore the original stdout
	w.Close()
	os.Stdout = oldStdout // restoring the real stdout
	return <-capturedCh
}

func generateResponse(body *bytes.Buffer, res *http.Response) {
	// quit early if the hook has already generated a response body
	if res.Body != nil {
		return
	}

	// how large is the response body?
	res.ContentLength = int64(body.Len())

	// assume the hook's output on stdout was the body
	res.Body = ioutil.NopCloser(body)

	// write response to disk
	out, err := os.Create("response")
	maybePanic(err)
	defer out.Close()
	err = res.Write(out)
	maybePanic(err)
}

func saveState(original []byte, state interface{}) {
	updated, err := json.Marshal(state)
	maybePanic(err)
	if bytes.Equal(original, updated) {
		// don't change state file
		return
	}

	err = ioutil.WriteFile("state", updated, 0666)
	maybePanic(err)
}

func maybePanic(err error) {
	if err != nil {
		panic(err)
	}
}
