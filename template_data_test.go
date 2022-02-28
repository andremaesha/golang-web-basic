package golangweb

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TemplateDataMap(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/name.gohtml"))

	t.ExecuteTemplate(w, "name.gohtml", map[string]interface{}{
		"Title": "Template Data Map",
		"Name":  "Andre Maesha",
		"Address": map[string]interface{}{
			"Street": "JL. PUCUNG III. NO 30",
		},
	})
}

func TestTemplateDataMap(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateDataMap(recorder, req)

	res := recorder.Result()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(string(body))
}

func TestTemplateDataServer(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(TemplateDataMap),
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}

type Address struct {
	Street string
}

type Page struct {
	Title   string
	Name    string
	Address Address
}

func TemplateDataStruct(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/name.gohtml"))

	t.ExecuteTemplate(w, "name.gohtml", Page{
		Title: "Template Data Struct",
		Name:  "Andre Maesha",
		Address: Address{
			Street: "JL. PUCUNG III. NO: 30",
		},
	})
}

func TestTemplateDataStruct(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateDataStruct(recorder, req)

	res := recorder.Result()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(string(body))
}
