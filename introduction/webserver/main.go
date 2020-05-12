package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request){
		writer.Write([]byte("Hello World!"))
	})

	if err := http.ListenAndServe(":9000", nil); err != nil {
		panic(err)
	}
}