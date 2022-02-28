package golangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func RequestHeader(writer http.ResponseWriter, req *http.Request) {
	contentType := req.Header.Get("content-type")
	fmt.Fprint(writer, contentType)
}

func TestRequestHeader(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "http://localhost:8080/", nil)
	req.Header.Add("Content-type", "application/json")

	recorder := httptest.NewRecorder()

	RequestHeader(recorder, req)

	res := recorder.Result()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(string(body))
}

func ResponseHeader(writer http.ResponseWriter, req *http.Request) {
	writer.Header().Add("X-Powered-By", "golang")
	fmt.Fprint(writer, "OK")
}

func TestResponseHeader(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "http://localhost:8080/", nil)
	req.Header.Add("Content-type", "application/json")

	recorder := httptest.NewRecorder()

	ResponseHeader(recorder, req)

	res := recorder.Result()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(string(body))
	fmt.Println(res.Header.Get("X-Powered-By"))
}
