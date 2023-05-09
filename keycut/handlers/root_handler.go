package handlers

import (
	"io"
	"net/http"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "This is my website!\n")
}
