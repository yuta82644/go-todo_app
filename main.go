package main

import (
	"fmt"
	
	// "log"
	//"github.com/yuta82644/go-todo_app/config"
	"github.com/yuta82644/go-todo_app/app/models"
	"github.com/yuta82644/go-todo_app/app/controllers"
)

func main() {
	fmt.Println(models.Db)

	controllers.StartMainServer()


	// セッションの確認
	// user, _ := models.GetUserByEmail("test@go.com")
	// fmt.Println(user)
	
	// session, err := user.CreateSession()
	// if err != nil {
	// 	log.Println(err)
	// }
	// fmt.Println(session) 

	// valid, _ := session.CheckSession()
	// fmt.Println(valid)
}