package controllers

import (
	"log"
	"net/http"
    // "github.com/yuta82644/go-todo_app/app/app/controllers"
	"github.com/yuta82644/go-todo_app/app/models"
)

// 新規登録
func signup(w http.ResponseWriter, r *http.Request) {
	//formの表示
	if r.Method == "GET" {
		_, err := session(w, r)
		if err != nil {
			generateHTML(w, nil, "layout", "public_navbar", "signup")
		} else {
			http.Redirect(w, r, "/todos", http.StatusFound)
		}
		//Postの場合　登録
	} else if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
		}
		user := models.User{
			Name:     r.PostFormValue("name"),
			Email:    r.PostFormValue("email"),
			Password: r.PostFormValue("password"),
		}
		if err := user.CreateUser(); err != nil {
			log.Println(err)
		}

		http.Redirect(w, r, "/", http.StatusFound)
	}
}

//login

func login(w http.ResponseWriter, r *http.Request) {
    _, err := session(w, r)
    if err != nil {
        generateHTML(w, nil, "layout", "public_navbar", "login")
        
    } else {
        http.Redirect(w, r, "/todos", http.StatusFound)
    }

}




// //session
// func authenticate(w http.ResponseWriter, r *http.Request) {
// 	err := r.ParseForm()
//     user, err := models.GetUserByEmail(r.PostFormValue("email"))
// 	if err != nil {
// 		log.Println(err)
// 		http.Redirect(w, r, "/login", http.StatusFound) //302
// 	}
//     if user.Password == models.HashPassword(r.PostFormValue("password")) {
//         session, err := user.CreateSession()
//         if err != nil {
//             log.Println(err)
//         }
//         cookie := http.Cookie{
//             Name:       "_cookie",
//             Value:      session.UUID,
//             HttpOnly:   true,
//         }
//         http.SetCookie(w, &cookie)

//         http.Redirect(w, r, "/", http.StatusFound)
//     } else {
//         http.Redirect(w, r, "/login", http.StatusFound)
//     }
// }
	
func authenticate(w http.ResponseWriter, r *http.Request) {
    err := r.ParseForm()
    if err != nil {
        log.Println(err)
        http.Redirect(w, r, "/login", http.StatusFound)
        return
    }

    email := r.PostFormValue("email")
    password := r.PostFormValue("password")

    user, err := models.GetUserByEmail(email)
    if err != nil {
        log.Println(err)
        http.Redirect(w, r, "/login", http.StatusFound)
        return
    }

    hashedPassword, err := models.HashPassword(password)
    if err != nil {
        log.Println(err)
        http.Redirect(w, r, "/login", http.StatusFound)
        return
    }

    if user.Password != hashedPassword {
        log.Println("パスワードが一致しません")
        http.Redirect(w, r, "/login", http.StatusFound)
        return
    }

    session, err := user.CreateSession()
    if err != nil {
        log.Println(err)
        http.Redirect(w, r, "/login", http.StatusFound)
        return
    }

    cookie := http.Cookie{
        Name:     "_cookie",
        Value:    session.UUID,
        HttpOnly: true,
    }
    http.SetCookie(w, &cookie)

    http.Redirect(w, r, "/login", http.StatusFound)
}

func logout(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("_cookie")
	if err != nil {
		log.Println(err)
	}
	if err != http.ErrNoCookie {
		session := models.Session{UUID: cookie.Value}
		session.DeleteSessionByUUID()
	}
	http.Redirect(w, r, "/login", http.StatusFound)
}
