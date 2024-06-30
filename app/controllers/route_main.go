package controllers

import (
	// "html/template"
	"net/http"
	// "log"
)

func top(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	if err != nil {
		generateHTML(w, "hello", "layout", "public_navbar", "top")
	} else {
		http.Redirect(w, r, "/todos", http.StatusFound)
	}
}
  //アクセス制限、ログインしている場合は
  func index(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/", http.StatusFound)
	} else {
		generateHTML(w, nil, "layout", "private_navbar", "index")
	}
  }