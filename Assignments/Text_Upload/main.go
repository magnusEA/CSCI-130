package main

import (
	"fmt"
	"io"
	"net/http"
)


func main() {
	http.HandleFunc("/", upload)
	http.ListenAndServe(":8080", nil)

}

func upload(w http.ResponseWriter, r *http.Request) {
	key := "file"

	file, hdr, err := r.FormFile(key)

	fmt.Println(file, hdr, err)

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	io.WriteString(w, `<form method="POST" enctype="multipart/form-data">
      <input type="file" name="file">
      <input type="submit">
    </form>`)
}
