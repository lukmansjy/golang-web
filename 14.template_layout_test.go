package golang_web

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TemplateLayout(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles(
		"./templates/header.gohtml",
		"./templates/footer.gohtml",
		"./templates/layout.gohtml",
	))

	t.ExecuteTemplate(w, "layout", map[string]interface{}{
		"Title": "Template Layout",
		"Name":  "Lukman",
	})
}

func TestTemplateLayout(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recoder := httptest.NewRecorder()

	TemplateLayout(recoder, request)

	body, _ := io.ReadAll(recoder.Result().Body)
	fmt.Println(string(body))
}
