package main

import "net/http"

func main() {
	fs := http.FileServer(http.Dir("/home/pawelhenek/Pulpit/playground/languages/go/go-web-examples/src/com/github/pawelhenek/playground/assets/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":8081", nil)
}
