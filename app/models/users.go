package models

import (
    "time"
    "log"
	// "github.com/yuta82644/go-todo_app/app/models"

)

type User struct {
    ID        int
    UUID      string
    Name      string
    Email     string
    Password  string
    CreatedAt time.Time
	Todos	  []Todo
}

type Session struct {
    ID        int
    UUID      string
    Email     string
    UserID    int
    CreatedAt time.Time
}

func (u *User) CreateUser() (err error) {
    hashedPassword, err := HashPassword(u.Password)
    if err != nil {
        return err
    }

    cmd := `INSERT INTO users (
        uuid,
        name,
        email,
        password,
        created_at
    ) VALUES (?, ?, ?, ?, ?)`

    _, err = Db.Exec(cmd,
        createUUID(),
        u.Name,
        u.Email,
        hashedPassword,
        time.Now(),
    )
    if err != nil {
        log.Println(err)
    }
    return err
}

func GetUser(id int) (user User, err error) {
    user = User{}
    cmd := `SELECT id, uuid, name, email, password, created_at FROM users WHERE id = ?`
    err = Db.QueryRow(cmd, id).Scan(
        &user.ID,
        &user.UUID,
        &user.Name,
        &user.Email,
        &user.Password,
        &user.CreatedAt,
    )
    return user, err
}


func (u *User) UpdateUser() (err error) {
    cmd := `UPDATE users SET name = ?, email = ? WHERE id = ?`
    _, err = Db.Exec(cmd, u.Name, u.Email, u.ID)
    if err != nil {
        log.Println(err)
    }
    return err
}

func (u *User) DeleteUser() (err error) {
    cmd := `DELETE FROM users WHERE id = ?`
    _, err = Db.Exec(cmd, u.ID)
    if err != nil {
        log.Println(err)
    }
    return err
}

func GetUserByEmail(email string) (user User, err error) {
    user = User{}
    cmd := `SELECT id, uuid, name, email, password, created_at FROM users WHERE email = ?`
    err = Db.QueryRow(cmd, email).Scan(
        &user.ID,
        &user.UUID,
        &user.Name,
        &user.Email,
        &user.Password,
        &user.CreatedAt,
    )

    return user, err
}

func (u *User) CreateSession() (session Session, err error) {
    session = Session{}

    cmd1 := `insert into sessions (
        uuid,
        email,
        user_id,
        created_at) VALUES (?, ?, ?, ?)`

		_, err = Db.Exec(cmd1, createUUID(), u.Email,  u.ID, time.Now())

    if err != nil {
        log.Println(err)
    }

    cmd2 := `SELECT id, uuid, email, user_id, created_at FROM sessions WHERE user_id = ? and email = ?`
    err = Db.QueryRow(cmd2, u.ID, u.Email).Scan(
        &session.ID,
        &session.UUID,
        &session.Email,
        &session.UserID,
        &session.CreatedAt,
    )
    return session, err
}

func (sess *Session) CheckSession() (valid bool, err error) {
    cmd := `SELECT id, uuid, email, user_id, created_at FROM sessions WHERE uuid = ?`

    err = Db.QueryRow(cmd, sess.UUID).Scan(
        &sess.ID,
        &sess.UUID,
        &sess.Email,
        &sess.UserID,
        &sess.CreatedAt,
    )
    if err != nil {
        valid = false
		return
    }
	if sess.ID != 0 {
		valid =true
	}
	return valid, err
}
//logout
func (sess *Session) DeleteSessionByUUID() error {
    cmd := `DELETE FROM sessions WHERE uuid = ?`
    _, err := Db.Exec(cmd, sess.UUID)
    if err != nil {
        log.Fatalln(err)
    }
    return err
}

func (sess *Session) GetUserBySession() (user User, err error) {
	user = User{}
	cmd := `select id, uuid, name, email, created_at FROM users where id = ?`
	err = Db.QueryRow(cmd, sess.UserID).Scan(
		&user.ID,
		&user.UUID,
		&user.Name,
		&user.Email,
		&user.CreatedAt)
	
		return user, err
}
