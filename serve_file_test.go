package golangweb

import (
	_ "embed"
	"fmt"
	"net/http"
	"testing"
)

func ServeFile(writer http.ResponseWriter, req *http.Request) {
	if req.URL.Query().Get("name") != "" {
		http.ServeFile(writer, req, "./resources/ok.html")
	} else {
		http.ServeFile(writer, req, "./resources/notfound.html")
	}
}

func TestServeFileServer(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(ServeFile),
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}

//go:embed resources/ok.html
var resourcesOk string

//go:embed resources/notfound.html
var resourcesNotfound string

func ServeFileEmbed(w http.ResponseWriter, r *http.Request) {
	if r.URL.Query().Get("name") != "" {
		fmt.Fprint(w, resourcesOk)
	} else {
		fmt.Fprint(w, resourcesNotfound)
	}
}

func TestServeFileServerEmbed(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(ServeFileEmbed),
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
