package main

import (
	"net/http"
	"fmt"
)

func readMe(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Please read me, first.")
	res.Write([]byte("Please read me, first."))
}

func main() {

	// this is a handler function by using function directly
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte("Hello, world!"))
	})

	// this is a handler function by using function name
	http.HandleFunc("/readme", readMe)

	http.ListenAndServe(":8080", nil)
}