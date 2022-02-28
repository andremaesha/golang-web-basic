package golangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func ResponseCode(writer http.ResponseWriter, req *http.Request) {
	name := req.URL.Query().Get("name")

	if name == "" {
		writer.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(writer, "name is empty")
	} else {
		fmt.Fprintf(writer, "Hello %s", name)
	}
}

func TestResponseCodeInvalid(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	ResponseCode(recorder, req)

	res := recorder.Result()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(res.Status)
	fmt.Println(res.StatusCode)
	fmt.Println(string(body))
}

func TestResponseCodeValid(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8080/?name=andre", nil)
	recorder := httptest.NewRecorder()

	ResponseCode(recorder, req)

	res := recorder.Result()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(res.Status)
	fmt.Println(res.StatusCode)
	fmt.Println(string(body))
}
