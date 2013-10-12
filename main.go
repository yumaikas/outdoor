package main

import (
	"fmt"
	//"io"
	"io/ioutil"
	"net/http"
	"os"
	//"encoding/csv"
)

var (
	wd string
	)

func init() {
	var err error
	wd, err = os.Getwd()
	if err != nil {
		panic("wd failed")
	}
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
	css, err := ioutil.ReadFile(wd + `/stylesheet.css`)
	if err != nil {
		panic("css file not found")
	}
	w.Write(css)
}
func index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("index")
	w.WriteHeader(200)
	buf, err := ioutil.ReadFile(wd + `/index.html`)
	if err != nil {
		panic("Index not found!")
	}
	w.Write(buf)
}

func acceptPost(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	name := r.PostFormValue("GET-name")
	phone := r.PostFormValue("GET-phone")
	email := r.PostFormValue("email")
	c_email := r.PostFormValue("confirm-email")
	fmt.Println(name, phone, email, c_email)
	http.Redirect(w, r, "/", 303)
}
