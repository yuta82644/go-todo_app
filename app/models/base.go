package models

import (
	"log"
	"fmt"
    "database/sql"
    "github.com/yuta82644/go-todo_app/config"
    _ "github.com/mattn/go-sqlite3" // Import for sqlite3 driver
    "github.com/google/uuid" //UUID
    "errors"
    "golang.org/x/crypto/bcrypt"//password
   
)

var Db *sql.DB
var err error

const (
    tableNameUser = "users"
    tableNameTodo = "todos"
    tableNameSession = "sessions"
)

func init() {
    Db, err = sql.Open(config.Config.SQLDriver, config.Config.DbName)
	if err != nil {
		log.Fatalln(err)
    }

    cmdU := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        uuid STRING NOT NULL UNIQUE,
        name STRING,
        email STRIMG,
        password STRING,
        created_at DATETIME
    )`, tableNameUser)

    Db.Exec(cmdU)

    //todoテーブル
    cmdT := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        content TEXT,
        user_id INTEGER,
        created_at DATETIME)`, tableNameTodo)

    Db.Exec(cmdT)

        //sessionsテーブル
    cmdS := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        uuid STRING NOT NULL UNIQUE,
        email STRIMG,
        user_id INTEGER,
        created_at DATETIME)`, tableNameSession)
    Db.Exec(cmdS)
}

//UUID
func createUUID() (uuidobj uuid.UUID) {
    uuidobj, _= uuid.NewUUID()
    return uuidobj
}
// HashPassword 
func HashPassword(password string) (string, error) {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
      return "", errors.New("パスワードのハッシュ化に失敗しました")
    }
    return string(hashedPassword), nil
  }

