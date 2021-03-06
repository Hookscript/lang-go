package main

import (
	"html/template"
	"net/http"
)

func Hook(w http.ResponseWriter, req *http.Request) {
	whom := req.FormValue("whom")
	if whom == "" {
		whom = "world"
	}
	w.Header().Set("Content-Type", "text/html")
	t, err := template.New("page").Parse(page)
	if err != nil {
		panic(err)
	}
	err = t.Execute(w, whom)
	if err != nil {
		panic(err)
	}
}

const page = `
<html>
    <body>
        <h1>Hello, {{.}}!</h1>

        <form method=GET>
            <input type=submit value="Say Hello to" />
            <input type=text name=whom placeholder="world" />
        </form>
    </body>
</html>
`
