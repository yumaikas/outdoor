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
	if err != nil {
		panic("wd failed")
	}
	buf, err := ioutil.ReadFile(wd + `/index.html`)
	if err != nil {
		panic("Index not found!")
	}
	idx = buf
	b_css, err := ioutil.ReadFile(wd + `/stylesheet.css`)
	if err != nil {
		panic("css file not found")
	}
	css = b_css
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/submit", acceptPost)
	http.HandleFunc("/stylesheet.css", getStyle)
	fmt.Println(http.ListenAndServe(":8080", nil))

}

func getStyle(w http.ResponseWriter, r *http.Request){
	fmt.Println("css")
	w.Header().Set("Content-Type", "text/css")
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
	name := r.PostFormValue("GET-name")
	phone := r.PostFormValue("GET-phone")
	email := r.PostFormValue("email")
	c_email := r.PostFormValue("confirm-email")
	fmt.Println("form values", name, phone, email, c_email)
	http.Redirect(w, r, "/", 303)
}
