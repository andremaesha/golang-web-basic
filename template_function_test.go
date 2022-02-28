package golangweb

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type MyPage struct {
	Name string
}

func (myPage MyPage) SayHello(name string) string {
	return "Hello " + name + ", My Name is " + myPage.Name
}

func TemplateFunction(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.New("FUNCTION").Parse(`{{.SayHello "Eko"}}`))

	t.ExecuteTemplate(w, "FUNCTION", MyPage{
		Name: "Andre",
	})
}

func TestTemplateFunction(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateFunction(recorder, req)

	res := recorder.Result()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(string(body))
}

func TemplateFunctionGlobal(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.New("FUNCTION").Parse("{{len .Name}}"))

	t.ExecuteTemplate(w, "FUNCTION", MyPage{
		Name: "Andre Maesha",
	})
}

func TestTemplateFunctionGlobal(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateFunctionGlobal(recorder, req)

	res := recorder.Result()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(string(body))
}

// create global function inside golang template
func TemplateFunctionCreateGlobal(w http.ResponseWriter, r *http.Request) {
	t := template.New("FUNCTION")

	t = t.Funcs(map[string]interface{}{
		"upper": func(value string) string {
			return strings.ToUpper(value)
		},
	})

	t = template.Must(t.Parse(`{{upper .Name}}`))

	t.ExecuteTemplate(w, "FUNCTION", MyPage{
		Name: "andre maesha",
	})
}

func TestTemplateFunctionCreateGlobal(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateFunctionCreateGlobal(recorder, req)

	res := recorder.Result()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(string(body))
}

// pipline
func TemplateFunctionCreateGlobalPipline(w http.ResponseWriter, r *http.Request) {
	t := template.New("FUNCTION")
	t = t.Funcs(map[string]interface{}{
		"sayHello": func(name string) string {
			return "Hello " + name
		},
		"upper": func(value string) string {
			return strings.ToUpper(value)
		},
	})

	t = template.Must(t.Parse(`{{sayHello .Name | upper}}`))

	t.ExecuteTemplate(w, "FUNCTION", MyPage{
		Name: "andre maesha",
	})
}

func TestTemplateFunctionCreateGlobalPipline(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateFunctionCreateGlobalPipline(recorder, req)

	res := recorder.Result()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(string(body))
}
