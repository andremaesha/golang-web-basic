package golangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func SayHello(writer http.ResponseWriter, request *http.Request) {
	name := request.URL.Query().Get("name")

	if name == "" {
		fmt.Fprint(writer, "Hello")
	} else {
		fmt.Fprintf(writer, "Hello %s", name)
	}
}

func TestQueryParameter(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello?name=andre", nil)
	recorder := httptest.NewRecorder()

	SayHello(recorder, req)

	res := recorder.Result()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(string(body))
}

func MultipleQueryParameter(writer http.ResponseWriter, request *http.Request) {
	firstname := request.URL.Query().Get("first_name")
	lastname := request.URL.Query().Get("last_name")

	fmt.Fprintf(writer, "Hello %s %s", firstname, lastname)
}

func TestMultipleQueryParameter(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello?first_name=andre&last_name=maesha", nil)
	recorder := httptest.NewRecorder()

	MultipleQueryParameter(recorder, req)

	res := recorder.Result()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(string(body))
}

func MultipleParameterValues(writer http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	names := query["name"]

	fmt.Fprint(writer, strings.Join(names, " "))
}

func TestMultipleParameterValues(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello?name=andre&name=maesha&name=golang", nil)
	recorder := httptest.NewRecorder()

	MultipleParameterValues(recorder, req)

	res := recorder.Result()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(string(body))
}
