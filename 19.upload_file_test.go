package golang_web

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func UploadForm(w http.ResponseWriter, r *http.Request) {
	err := myTemplates.ExecuteTemplate(w, "upload.form.gohtml", nil)
	if err != nil {
		panic(err)
	}
}

func Upload(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(50 << 20) // max 50 MB
	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		panic(err)
	}

	fileDestination, err := os.Create("./resources/" + fileHeader.Filename)
	if err != nil {
		panic(err)
	}

	_, err = io.Copy(fileDestination, file)
	if err != nil {
		panic(err)
	}

	name := r.PostFormValue("name")
	myTemplates.ExecuteTemplate(w, "upload.success.gohtml", map[string]interface{}{
		"Name": name,
		"File": "/static/" + fileHeader.Filename,
	})
}

func TestUploadForm(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/form", UploadForm)
	mux.HandleFunc("/upload", Upload)
	mux.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./resources"))))

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

//go:embed resources/golang.png
var uploadFileTest []byte

// UNIT TEST UPLOAD
func TestUploadUnit(t *testing.T) {
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	writer.WriteField("name", "Lukman Sanjaya")

	file, _ := writer.CreateFormFile("file", "NewImageGolang.png")
	file.Write(uploadFileTest)
	writer.Close()

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", body)
	request.Header.Set("Content-Type", writer.FormDataContentType())
	recoder := httptest.NewRecorder()

	Upload(recoder, request)

	bodyResponse, _ := io.ReadAll(recoder.Result().Body)
	fmt.Println(string(bodyResponse))
}
