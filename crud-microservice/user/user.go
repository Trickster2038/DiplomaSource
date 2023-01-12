package user

import (
	"crud/connection"
)

type User struct {
	ID       int    `json:"id"`
	Nickname string `json:"nickname"`
	IsAdmin  bool   `json:"is_admin"`
}

func (user *User) Read() {
	db := connection.Connect_db()
	err := db.QueryRow("SELECT id, nickname, is_admin FROM Users where id = ?", user.ID).
		Scan(&user.ID, &user.Nickname, &user.IsAdmin)
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
}
