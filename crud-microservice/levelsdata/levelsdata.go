package levelsdata

import (
	"crud/connection"
)

type LevelsData struct {
	ID              int    `json:"id"`
	WideDescription string `json:"wide_description"`
	Code            string `json:"code"`
	Question        string `json:"question"`
	Answer          string `json:"answer"`
}

func (level_data LevelsData) CreateOrUpdate() {
	db := connection.Connect_db()

	level_data.Delete()

	_, err := db.Query("INSERT INTO LevelsData (id, wide_description, code, question, answer) VALUES (?, ?, ?, ?, ?)",
		level_data.ID,
		level_data.WideDescription,
		level_data.Code,
		level_data.Question,
		level_data.Answer)
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
}

func (level_data *LevelsData) Read(id int) {
	db := connection.Connect_db()
	err := db.QueryRow("SELECT id, wide_description, code, question, answer FROM LevelsData where id = ?",
		id).
		Scan(&level_data.ID,
			&level_data.WideDescription,
			&level_data.Code,
			&level_data.Question,
			&level_data.Answer)
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
}

func (level_data LevelsData) Update() {
	db := connection.Connect_db()

	var old_level_data LevelsData
	old_level_data.Read(level_data.ID) // existence check

	_, err := db.Query("UPDATE LevelsData SET "+
		"wide_description = ?, code = ?, question = ?, answer = ? "+
		"WHERE id = ?",
		level_data.WideDescription,
		level_data.Code,
		level_data.Question,
		level_data.Answer,
		level_data.ID)
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
}

func (level_data LevelsData) Delete() {
	db := connection.Connect_db()

	// level_data.Read(level_data.ID) // check existence

	_, err := db.Query("DELETE FROM LevelsData WHERE id = ?",
		level_data.ID)
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
}
