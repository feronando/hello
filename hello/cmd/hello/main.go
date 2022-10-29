package main

import (
	"net/http"

	"workshop.go/hello/internal"
)

func main() {
	handler := internal.NewHandler() //	http.HandlerFunc(hello)
	http.ListenAndServe(":8080", handler)
}

/*
func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}
*/
