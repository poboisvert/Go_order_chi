package main

import (
	"fmt"
	"net/http"
)

func main() {

	server := &http.Server{
		Addr:    ":3000",
		Handler: http.HandlerFunc(basicHandler),
	}

	err := server.ListenAndServe()

	if err != nil {
		fmt.Println("Error to Listen")
	}
}

func basicHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}
