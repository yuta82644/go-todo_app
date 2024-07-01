package controllers

import (
	// "html/template"
	"log"
	"net/http"

	"github.com/yuta82644/go-todo_app/app/models"
	// "log"
	// "github.com/yuta82644/go-todo_app/app/models"
)

func top(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	if err != nil {
		generateHTML(w, nil, "layout", "public_navbar", "top")
	} else {
		http.Redirect(w, r, "/todos", http.StatusFound)
	}
}
  //アクセス制限、ログインしている場合は
  func index(w http.ResponseWriter, r *http.Request) {
    sess, err := session(w, r)
    if err != nil {
        http.Redirect(w, r, "/", http.StatusFound)
        return
    }

    user, err := sess.GetUserBySession()
    if err != nil {
        log.Println(err)
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    todos, _ := user.GetTodosByUser()
    user.Todos = todos

    generateHTML(w, user, "layout", "private_navbar", "index")
}

  func todoNew(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
	} else {
		generateHTML(w, nil, "layout", "private_navbar", "/todo_new")
	}
  }

  	func todoSave(w http.ResponseWriter, r *http.Request) {
		sess, err := session(w, r)
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusFound)
		} else {
			err = r.ParseForm()
			if err != nil {
				log.Println(err)
			}
			user, err := sess.GetUserBySession()
			if err != nil {
				log.Println(err)
			}

			content := r.PostFormValue("content")
			if err := user.CreateTodo(content); err != nil {
				log.Println(err)
			}

			http.Redirect(w, r, "/todos", http.StatusFound)
		}
	}

func todoEdit(w http.ResponseWriter, r *http.Request, id int) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
	} else {
		_, err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		t, err := models.GetTodo(id)
		if err != nil {
			log.Println(err)
		}
		generateHTML(w, t, "layout", "private_navbar", "todo_edit")
	}
}

func todoUpdate(w http.ResponseWriter, r *http.Request, id int) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
	} else {
		err := r.ParseForm()
		if err != nil {
			log.Panicln(err)
		}
		user, err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		content := r.PostFormValue("content")
		t := &models.Todo{ID: id, Content: content, UserID: user.ID}
		if err := t.UpdateTodo(); err != nil {
			log.Println(err)
		}
		http.Redirect(w, r, "/todos", http.StatusFound)
	}
}
