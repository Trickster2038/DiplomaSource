package typerecord

import (
	"crud/connection"
)

type Type struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (type_struct *Type) Read() {
	db := connection.Connect_db()
	err := db.QueryRow("SELECT id, Name FROM Types where id = ?", type_struct.ID).
		Scan(&type_struct.ID, &type_struct.Name)
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
}

func (type_struct *Type) ReadByName() {
	db := connection.Connect_db()
	err := db.QueryRow("SELECT id, Name FROM Types where name = ?", type_struct.Name).
		Scan(&type_struct.ID, &type_struct.Name)
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
}
