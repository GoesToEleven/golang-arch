package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"

	"golang.org/x/crypto/bcrypt"
)

var db = map[string][]byte{}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/register", register)
	http.HandleFunc("/login", login)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	errMsg := r.FormValue("msg")

	fmt.Fprintf(w, `<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<meta http-equiv="X-UA-Compatible" content="ie=edge">
		<title>Document</title>
	</head>
	<body>
        <h1>IF THERE IS ANY MESSAGE FOR YOU, HERE IT IS: %s</h1>
        <h1>REGISTER</h1>
		<form action="/register" method="POST">
			<input type="email" name="e">
			<input type="password" name="p">
			<input type="submit">
        </form>
        <h1>LOG IN</h1>
        <form action="/login" method="POST">
            <input type="email" name="e">
			<input type="password" name="p">
			<input type="submit">
        </form>
	</body>
	</html>`, errMsg)
}

func register(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		msg := url.QueryEscape("your method was not post")
		http.Redirect(w, r, "/?msg="+msg, http.StatusSeeOther)
		return
	}

	e := r.FormValue("e")
	if e == "" {
		msg := url.QueryEscape("your email needs to not be empty")
		http.Redirect(w, r, "/?msg="+msg, http.StatusSeeOther)
		return
	}

	p := r.FormValue("p")
	if p == "" {
		msg := url.QueryEscape("your email password needs to not be empty")
		http.Redirect(w, r, "/?msg="+msg, http.StatusSeeOther)
		return
	}

	bsp, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
	if err != nil {
		msg := "there was an internal server error - evil laugh: hahahahaha"
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}
	log.Println("password", e)
	log.Println("bcrypted", bsp)
	db[e] = bsp

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		msg := url.QueryEscape("your method was not post")
		http.Redirect(w, r, "/?msg="+msg, http.StatusSeeOther)
		return
	}

	e := r.FormValue("e")
	if e == "" {
		msg := url.QueryEscape("your email needs to not be empty")
		http.Redirect(w, r, "/?msg="+msg, http.StatusSeeOther)
		return
	}

	p := r.FormValue("p")
	if p == "" {
		msg := url.QueryEscape("your email password needs to not be empty")
		http.Redirect(w, r, "/?msg="+msg, http.StatusSeeOther)
		return
	}

	if _, ok := db[e]; !ok {
		msg := url.QueryEscape("your email or password didn't match")
		http.Redirect(w, r, "/?msg="+msg, http.StatusSeeOther)
		return
	}

	err := bcrypt.CompareHashAndPassword(db[e], []byte(p))
	if err != nil {
		msg := url.QueryEscape("your email or password didn't match")
		http.Redirect(w, r, "/?msg="+msg, http.StatusSeeOther)
		return
	}

	msg := url.QueryEscape("you logged in " + e)
	http.Redirect(w, r, "/?msg="+msg, http.StatusSeeOther)
}
