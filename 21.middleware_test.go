package golang_web

import (
	"fmt"
	"net/http"
	"testing"
)

type LogMiddleware struct {
	Handler http.Handler
}

// mengikuti kontrak ServeHTTP
func (middleware *LogMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Before Execute Handler")
	middleware.Handler.ServeHTTP(w, r)
	fmt.Println("After Execute Handler")
}

type ErrorHandler struct {
	Handler http.Handler
}

func (errorHandler *ErrorHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("RECOVER : ", err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Error : %s", err)
		}
	}()
	errorHandler.Handler.ServeHTTP(w, r)
}

func TestMiddleware(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Handler Executed")
		fmt.Fprintf(w, "Hello Middleware")
	})
	mux.HandleFunc("/foo", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Foo Executed")
		fmt.Fprintf(w, "Hello Middleware")
	})
	mux.HandleFunc("/panic", func(w http.ResponseWriter, r *http.Request) {
		panic("Oppss Error...")
	})

	logMiddleware := new(LogMiddleware)
	logMiddleware.Handler = mux

	errorHandler := new(ErrorHandler)
	errorHandler.Handler = logMiddleware

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: errorHandler,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
