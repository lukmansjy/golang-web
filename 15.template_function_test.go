package golang_web

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
	t := template.Must(template.New("FUNCTION").Parse(`{{ .SayHello "Budi" }}`))
	t.ExecuteTemplate(w, "FUNCTION", MyPage{Name: "Lukman"})
}

func TestTemplateFunction(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recoder := httptest.NewRecorder()

	TemplateFunction(recoder, request)

	body, _ := io.ReadAll(recoder.Result().Body)
	fmt.Println(string(body))
}

// Menggunakan global function di template
func TemplateFunctionGlobal(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.New("FUNCTION").Parse(`{{ len .Name }}`))
	t.ExecuteTemplate(w, "FUNCTION", MyPage{Name: "Lukman Sanjaya"})
}

func TestTemplateFunctionGlobal(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recoder := httptest.NewRecorder()

	TemplateFunctionGlobal(recoder, request)

	body, _ := io.ReadAll(recoder.Result().Body)
	fmt.Println(string(body))
}

// Membuat golbal function pada template
func TemplateFunctionCreateGlobal(w http.ResponseWriter, r *http.Request) {
	t := template.New("FUNCTION")
	t = t.Funcs(map[string]interface{}{
		"upper": func(value string) string {
			return strings.ToUpper(value)
		},
		"lower": func(value string) string {
			return strings.ToLower(value)
		},
	})
	t = template.Must(t.Parse(`upper: {{ upper .Name }} lower: {{ lower .Name }}`))
	t.ExecuteTemplate(w, "FUNCTION", MyPage{Name: "Lukman Sanjaya"})
}

func TestTemplateFunctionCreateGlobal(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recoder := httptest.NewRecorder()

	TemplateFunctionCreateGlobal(recoder, request)

	body, _ := io.ReadAll(recoder.Result().Body)
	fmt.Println(string(body))
}

// Function pipeline pada template
func TemplateFunctionPipeline(w http.ResponseWriter, r *http.Request) {
	t := template.New("FUNCTION")
	t = t.Funcs(map[string]interface{}{
		"sayHello": func(value string) string {
			return "Hello " + value
		},
		"upper": func(value string) string {
			return strings.ToUpper(value)
		},
	})
	t = template.Must(t.Parse(`{{ sayHello .Name | upper }}`))
	t.ExecuteTemplate(w, "FUNCTION", MyPage{Name: "Lukman Sanjaya"})
}

func TestTemplateFunctionPipeline(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recoder := httptest.NewRecorder()

	TemplateFunctionPipeline(recoder, request)

	body, _ := io.ReadAll(recoder.Result().Body)
	fmt.Println(string(body))
}
