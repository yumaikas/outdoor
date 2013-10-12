package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

var (
	idx []byte
	css []byte
)

func init() {
	wd, err := os.Getwd()
	buf, err := ioutil.ReadFile(wd + `\index.html`)
	if err != nil {
		panic("Index not found!")
	}
	idx = buf
	b_css, err := ioutil.ReadFile(wd + `\stylesheet.css`)
	if err != nil {
		panic("css file not found")
	}
	css = b_css
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/submit", acceptPost)
	http.HandleFunc("/stylesheet.css", getStyle)
	http.ListenAndServe(":8080", nil)
}

func getStyle(w http.ResponseWriter, r *http.Request){
	fmt.Println("css")
	w.WriteHeader(200)
	w.Write(css)
}
func index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("index")
	w.WriteHeader(200)
	w.Write(idx)
}

func acceptPost(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	name := r.FormValue("GET-Name")
	phone := r.FormValue("GET-phone")
	email := r.FormValue("email")
	c_email := r.FormValue("confirm-email")
	fmt.Println("form values", name, phone, email, c_email)
	http.Redirect(w, r, "/", 303)
}
