package user

import (
	"crud/connection"
)

type User struct {
	ID       int    `json:"id"`
	Nickname string `json:"nickname"`
}

func (user *User) Read(id int) {
	db := connection.Connect_db()
	err := db.QueryRow("SELECT id, nickname FROM Users where id = ?", id).
		Scan(&user.ID, &user.Nickname)
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
}
