package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {

	http.HandleFunc("/", handler)
	http.ListenAndServe("", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/login": // input boxes & button for credentials
		login(w, r)
	case "/loginSubmit": // handle login credentials
		loginSubmit(w, r)
	default:
		fmt.Fprintf(w, "Allan Kama")
	}

}

func login(w http.ResponseWriter, r *http.Request) {

	var fileName = "login.html"
	t, err := template.ParseFiles(fileName)
	if err != nil {
		fmt.Println("Error when parsing file", err)
		return
	}
	err = t.ExecuteTemplate(w, fileName, nil)
	if err != nil {
		fmt.Println("Error When executing template")
		return
	}
}

var userDB = map[string]string{
	"Allan": "allan111",
}

func loginSubmit(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	if userDB[username] == password {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "YOU HAVE LOGGED IN")
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Didn't find your credentials")
	}

}
