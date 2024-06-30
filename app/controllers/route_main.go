package controllers

import (
	// "html/template"
	"log"
	"net/http"
	// "log"
	// "github.com/yuta82644/go-todo_app/app/models"
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
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/", http.StatusFound)
	} else {
		user, err := sess.GetUserBySession()
		if err != nil{
			log.Println(err)
		}
		todos, _ := user.GetTodosByUser()
		user.Todos = todos
		generateHTML(w, user, "layout", "private_navbar", "index")
	}
  }

  func todoNew(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
	} else {
		generateHTML(w, nil, "layout", "private_navbar", "/todo_new")
	}
}
