package golangweb

import (
	"fmt"
	"net/http"
	"testing"
)

func RedirectTo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello Redirect")
}

func RedirectFrom(w http.ResponseWriter, r *http.Request) {
	// logic redirect
	http.Redirect(w, r, "/redirect-to", http.StatusTemporaryRedirect)
}

func RedirectOut(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://andre-maesha-99.web.app/", http.StatusTemporaryRedirect)
}

func TestRedirect(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/redirect-from", RedirectFrom)
	mux.HandleFunc("/redirect-to", RedirectTo)
	mux.HandleFunc("/redirect-out", RedirectOut)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
