package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {

	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}
func postHandler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "POST":
		fmt.Fprintf(w, "post\n")
	case "GET":
		fmt.Fprintf(w, "Get\n")
	default:
		fmt.Fprintf(w, "default\n")

	}

}
func main() {

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)
	http.HandleFunc("/api/post", postHandler)

	http.ListenAndServe(":8089", nil)
}
