package golang_web

import (
	"embed"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

//go:embed templates/*.gohtml
var templatesFile embed.FS

// memeangil template 1x di gobal variabel, supaya disimpan di memory. karena template parsing itu sangat berat
var myTemplates = template.Must(template.ParseFS(templatesFile, "templates/*.gohtml"))

func TemplateCaching(w http.ResponseWriter, r *http.Request) {
	myTemplates.ExecuteTemplate(w, "simple.gohtml", "Hello HTML Template")
}

func TestTemplateCaching(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recoder := httptest.NewRecorder()

	TemplateCaching(recoder, request)

	body, _ := io.ReadAll(recoder.Result().Body)
	fmt.Println(string(body))
}
