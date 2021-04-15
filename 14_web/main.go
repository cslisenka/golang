package main

import (
	"fmt"
	"io"
	"net/http"
)

func test(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Received request", req.RequestURI, req.Method)
	res.Header().Set("Content-Type", "text/html")
	io.WriteString(res, "Test")
}

func main() {
	http.HandleFunc("/test", test)
	http.ListenAndServe(":9000", nil)
}
