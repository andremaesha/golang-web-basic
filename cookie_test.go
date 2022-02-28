package golangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SetCookie(writer http.ResponseWriter, req *http.Request) {
	cookie := new(http.Cookie)
	cookie.Name = "data-cookie"
	cookie.Value = req.URL.Query().Get("name")
	cookie.Path = "/"

	http.SetCookie(writer, cookie)
	fmt.Fprint(writer, "succes create cookie")
}

func GetCookie(writer http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("data-cookie")

	if err != nil {
		fmt.Fprint(writer, "No Cookie")
	} else {
		name := cookie.Value
		fmt.Fprintf(writer, "Hello %s", name)
	}
}

func TestCookie(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/set-cookie", SetCookie)
	mux.HandleFunc("/get-cookie", GetCookie)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}

func TestSetCookie(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8080/?name=andre", nil)
	recorder := httptest.NewRecorder()

	SetCookie(recorder, req)

	cookies := recorder.Result().Cookies()

	for _, cookie := range cookies {
		fmt.Printf("Cookie %s:%s \n", cookie.Name, cookie.Value)
	}
}

func TestGetCookie(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	cookie := new(http.Cookie)
	cookie.Name = "data-cookie"
	cookie.Value = "andre"
	req.AddCookie(cookie)

	GetCookie(recorder, req)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
