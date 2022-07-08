package golang_web

import (
	"fmt"
	"net/http"
	"testing"
)

/*
Ketika browser akses url image, pdf, atau file yg bisa dirender oleh browser biasanya file tersebut akan
secara otatis dirender oleh browser (tidak didownload).
Dengan cara dibawah ini maka file yg diakses akan dipaksa untuk didownload
*/

func DownloadFile(w http.ResponseWriter, r *http.Request) {
	fileName := r.URL.Query().Get("file")
	if fileName == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "BAD REQUEST")
		return
	}

	w.Header().Add("Content-Disposition", "attachment; filename=\""+fileName+"\"")
	http.ServeFile(w, r, "./resources/"+fileName)
}

func TestDownloadFile(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(DownloadFile),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
