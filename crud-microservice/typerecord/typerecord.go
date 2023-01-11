package typerecord

import (
	"crud/connection"
)

type Type struct {
	ID   int    `json:"id"`
	Name string `json:"nickname"`
}

func (type_struct *Type) Read(id int) {
	db := connection.Connect_db()
	err := db.QueryRow("SELECT id, Name FROM Types where id = ?", id).
		Scan(&type_struct.ID, &type_struct.Name)
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
}
