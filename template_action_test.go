package golangweb

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

type user struct {
	Name  string
	Title string
	Age   int
}

func TemplateActionIf(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/if.gohtml"))

	t.ExecuteTemplate(w, "if.gohtml", user{
		Title: "Template Action IF",
		Name:  "Andre Maesha",
		Age:   22,
	})
}

func TestTemplateActionIF(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateActionIf(recorder, req)

	res := recorder.Result()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(string(body))
}

func TestTemplateActionIFServer(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(TemplateActionIf),
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}

func TemplateActionOperator(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/comparator.gohtml"))

	t.ExecuteTemplate(w, "comparator.gohtml", map[string]interface{}{
		"Title":      "Template Action Operator",
		"FinalValue": 70,
	})
}

func TestTemplateActionOperator(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateActionOperator(recorder, req)

	res := recorder.Result()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(string(body))
}

func TemplateActionRange(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/range.gohtml"))

	t.ExecuteTemplate(w, "range.gohtml", map[string]interface{}{
		"Title": "Template Action Range",
		"Hobbies": []string{
			"game",
			"workout",
			"coding",
		},
	})
}

func TestTemplateActionRange(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateActionRange(recorder, req)

	res := recorder.Result()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(string(body))
}

func TemplateActionWith(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/nestedStruct.gohtml"))

	t.ExecuteTemplate(w, "nestedStruct.gohtml", map[string]interface{}{
		"Title": "Template Action With",
		"Name":  "Andre Maesha",
		"Address": map[string]interface{}{
			"Street": "JL. PUCUNG III",
			"City":   "Jakarta Timur",
		},
	})
}

func TestTemplateActionWith(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateActionWith(recorder, req)

	res := recorder.Result()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(string(body))
}

type PersonAddress struct {
	Alamat   string
	Domisili string
	Kota     string
}

type PersonDetail struct {
	Title         string
	Name          string
	AddressDetail PersonAddress
}

func TemplateActionWithSecond(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/nestedStructSecond.gohtml"))
	t.ExecuteTemplate(w, "nestedStructSecond.gohtml", PersonDetail{
		Title: "Template Action With Second",
		Name:  "Rama",
		AddressDetail: PersonAddress{
			Alamat:   "JL. PUCUNG III",
			Domisili: "Jakarta Timur",
			Kota:     "Jakarta",
		},
	})
}

func TestTemplateActionWithSecond(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateActionWithSecond(recorder, req)

	res := recorder.Result()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(string(body))
}
