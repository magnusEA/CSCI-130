package main

import (
	"fmt"
	"net/http"
)

func Url_Path(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "localhost:8080/name%v", req.URL.Path)
}


func main() {
	http.HandleFunc("/", Url_Path)
	http.ListenAndServe(":8080", nil)
}
