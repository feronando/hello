package internal

import (
	"fmt"
	"net/http"
)

type handler struct {
}

func NewHandler() http.Handler {
	return &handler{}
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}
