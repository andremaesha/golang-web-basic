package golangweb

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TemplateAutoEscape(w http.ResponseWriter, r *http.Request) {
	myTemplates.ExecuteTemplate(w, "post.gohtml", map[string]interface{}{
		"Title": "Template Auto Escape",
		"Body":  "<p>This is body<script>alert('your account has been hacked')</script></p>",
	})
}

func TestTemplateAutoEscape(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateAutoEscape(recorder, req)

	res := recorder.Result()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(string(body))
}

func TestTemplateAutoEscapeServer(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(TemplateAutoEscape),
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}

func TemplateAutoEscapeDisabled(w http.ResponseWriter, r *http.Request) {
	myTemplates.ExecuteTemplate(w, "post.gohtml", map[string]interface{}{
		"Title": "Template Auto Escape",
		"Body":  template.HTML("<p>This is body</p>"),
	})
}

func TestTemplateAutoEscapeDisabled(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateAutoEscapeDisabled(recorder, req)

	res := recorder.Result()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(string(body))
}

func TestTemplateAutoEscapeDisabledServer(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(TemplateAutoEscapeDisabled),
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}

func TemplateXss(w http.ResponseWriter, r *http.Request) {
	myTemplates.ExecuteTemplate(w, "post.gohtml", map[string]interface{}{
		"Title": "Template Auto Escape",
		"Body":  template.HTML(r.URL.Query().Get("body")),
	})
}

func TestTemplateXss(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8080/?body=<p>alert</p>", nil)
	recorder := httptest.NewRecorder()

	TemplateXss(recorder, req)

	res := recorder.Result()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(string(body))
}

func TestTemplateXssServer(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(TemplateXss),
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
