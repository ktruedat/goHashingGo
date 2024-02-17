package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
	"html/template"
	"io"
	"net/http"
)

type HashInput struct {
	Text   string
	Method string
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			text := r.FormValue("text")
			method := r.FormValue("method")
			var hash string
			switch method {
			case "md5":
				hash = HashMd5(text)
			case "sha1":
				hash = HashSha1(text)
			case "sha256":
				hash = HashSha256(text)
			}
			w.Write([]byte(hash))
		} else {
			tmpl := template.Must(template.ParseFiles("~/GolandProjects/goHashingGo/internal/templates/index.html"))
			tmpl.Execute(w, nil)
		}
	})

	http.ListenAndServe(":8080", nil)
}

func HashMd5(text string) string {
	h := md5.New()
	io.WriteString(h, text)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func HashSha1(text string) string {
	h := sha1.New()
	io.WriteString(h, text)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func HashSha256(text string) string {
	h := sha256.New()
	io.WriteString(h, text)
	return fmt.Sprintf("%x", h.Sum(nil))
}
