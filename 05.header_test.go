package golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func RequsetHeader(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("content-type")
	fmt.Fprint(w, contentType)
}

func TestRequestHeader(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	request.Header.Add("Content-Type", "aplication/json")
	recorder := httptest.NewRecorder()

	RequsetHeader(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}

func ResponseHeader(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("X-Powered-By", "Lukman Sanjaya")
	fmt.Fprint(w, "Ok")
}

func TestResponseHeaderr(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	request.Header.Add("Content-Type", "aplication/json")
	recorder := httptest.NewRecorder()

	ResponseHeader(recorder, request)
	response := recorder.Result()

	poweredBy := response.Header.Get("x-powered-by")
	fmt.Println(poweredBy)
}
