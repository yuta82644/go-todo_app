package controllers

import (
    "fmt"
    "html/template"
    "net/http"
	"log"
    "github.com/yuta82644/go-todo_app/app/models"
    "github.com/yuta82644/go-todo_app/config"
)

func generateHTML(writer http.ResponseWriter, data interface{}, filenames ...string) {
    var files []string
    for _, file := range filenames {
        files = append(files, fmt.Sprintf("app/views/templates/%s.html", file))
    }

    templates, err := template.ParseFiles(files...)
    if err != nil {
        http.Error(writer, err.Error(), http.StatusInternalServerError)
        log.Println(err)
        return
    }

    err = templates.ExecuteTemplate(writer, "layout", data)
    if err != nil {
        http.Error(writer, err.Error(), http.StatusInternalServerError)
        log.Println(err)
    }
}
func session(w http.ResponseWriter, r *http.Request) (sess models.Session, err error) {
    cookie, err := r.Cookie("_cookie")
    if err == nil {
        sess = models.Session{UUID: cookie.Value}
        if ok, _ := sess.CheckSession(); !ok {
            err = fmt.Errorf("Invalid session")
        }
    }
    return sess, err
}

func StartMainServer() error {
    files := http.FileServer(http.Dir(config.Config.Static))
    http.Handle("/static/", http.StripPrefix("/static/", files))

    http.HandleFunc("/", top)
    http.HandleFunc("/signup", signup)
    http.HandleFunc("/login", login)
    http.HandleFunc("/authenticate", authenticate)
	http.HandleFunc("/logout", logout)
    http.HandleFunc("/todos", index)
	http.HandleFunc("/todos/new", todoNew)
	http.HandleFunc("/todos/save", todoSave)
    return http.ListenAndServe(":"+config.Config.Port, nil)
}
