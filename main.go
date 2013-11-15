package main

//This is a very panicky piece of golang at the moment. That will need to be fixed
import (
	"fmt"
	//"io"
	"encoding/csv"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
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
	//Switched from hardcoded duplicate functions to hardcoded strings. Some day these will come from a config file.
	//The reason that http.Dir is not used is because that allows too easily much messing things up for this app,
	//which needs to be dirt simple for outside use
	css := "/stylesheet.css"
	idx := "/index.html"
	win := "/success.html"

	http.HandleFunc("/", sendFile(idx, "text/html"))
	http.HandleFunc(css, sendFile(css, "text/css"))
	http.HandleFunc(win, sendFile(win, "text/html"))
	http.HandleFunc("/submit", acceptPost)
	cmd := exec.Command("cmd", "/c", "start", "http://localhost:8080")
	buf, _ := cmd.CombinedOutput()
	fmt.Println(string(buf))
	fmt.Println(http.ListenAndServe(":8080", nil))
	//this is just to start a webserver on load
}

//type refers to the Content-Type of a web file
func sendFile(name, c_type string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		buf, err := ioutil.ReadFile(wd + name)
		if err != nil {
			panic(fmt.Sprintf("%s file not found!", name))
		}
		w.Header().Set("Content-Type", c_type)
		w.Write(buf)
	}
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
