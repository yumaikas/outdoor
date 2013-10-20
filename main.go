package main

//This is a very panicky piece of golang at the moment. That will need to be fixed
import (
	"fmt"
	//"io"
	"encoding/csv"
	"io/ioutil"
	"net/http"
	"os"
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
	//These functions are repetitious and hardcoded, but that saves having to lockdown files.
	http.HandleFunc("/", index)
	http.HandleFunc("/submit", acceptPost)
	http.HandleFunc("/stylesheet.css", getStyle)
	http.HandleFunc("/success.html", success)
	fmt.Println(http.ListenAndServe(":8080", nil))
}

func success(w http.ResponseWriter, r *http.Request) {
	fmt.Println("success")
	html, err := ioutil.ReadFile(wd + `/success.html`)
	if err != nil {
		panic("success file not found!")
	}
	w.Header().Set("Content-Type", "text/html")
	w.Write(html)
}

func getStyle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("css")
	css, err := ioutil.ReadFile(wd + `/stylesheet.css`)
	if err != nil {
		panic("css file not found")
	}
	w.Header().Set("Content-Type", "text/css")
	w.Write(css)
}
func index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("index")
	buf, err := ioutil.ReadFile(wd + `/index.html`)
	if err != nil {
		panic("Index not found!")
	}
	w.Write(buf)
}

//TODO: Add extra layer of verification for names, don't rely on JS in browser.
func acceptPost(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	name := r.PostFormValue("GET-name")
	phone := r.PostFormValue("GET-phone")
	email := r.PostFormValue("email")
	c_email := r.PostFormValue("confirm-email")
	path := wd + "names.csv"
	fmt.Println(path)
	var csv_w *csv.Writer
	if _, err := os.Stat(path); err == nil {
		fmt.Printf("file exists; processing...")
		f, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			panic(fmt.Sprintf("Failed to open file! Err:%v", err))
		}
		defer f.Close()
		csv_w = csv.NewWriter(f)
		defer csv_w.Flush()
	} else {
		if f, err := os.Create(path); err == nil {
			defer f.Close()
			csv_w = csv.NewWriter(f)
			fmt.Println("HEX")
			csv_w.Write([]string{"Name", "Phone", "Email"})
			defer csv_w.Flush()
		} else {
			panic("Unable to create file")
		}
	}
	//If we get to this point, no errors have occured
	fmt.Print("")
	csv_w.Write([]string{name, phone, email})
	fmt.Println(name, phone, email, c_email)
	http.Redirect(w, r, "/success.html", 303)
}
