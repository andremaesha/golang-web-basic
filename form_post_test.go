package golangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func FormPost(writer http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()

	if err != nil {
		panic(err)
	}

	// req.PostFormValue() //this can auto parse

	firstname := req.PostForm.Get("first_name")
	lastname := req.PostForm.Get("last_name")

	fmt.Fprintf(writer, "Hello %s %s", firstname, lastname)
}

func TestFormPost(t *testing.T) {
	requestBody := strings.NewReader("first_name=andre&last_name=maesha")
	req := httptest.NewRequest(http.MethodPost, "http://localhost:8080/", requestBody)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	recorder := httptest.NewRecorder()

	FormPost(recorder, req)

	res := recorder.Result()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(string(body))
}
