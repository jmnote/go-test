package myhttp3

import (
	"io"
	"net/http"
)

func Run() {
	http.HandleFunc("/v1/login", loginEndpoint)
	http.HandleFunc("/v1/submit", submitEndpoint)
	http.HandleFunc("/v1/read", readEndpoint)
	http.HandleFunc("/v2/login", loginEndpoint)
	http.HandleFunc("/v2/submit", submitEndpoint)
	http.HandleFunc("/v2/read", readEndpoint)
	http.ListenAndServe(":8080", nil)
}

func loginEndpoint(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "pong")
}

func submitEndpoint(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "pong")
}

func readEndpoint(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "pong")
}
