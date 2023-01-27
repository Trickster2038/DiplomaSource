package personal

import (
	"fmt"
	"stats/connection"
)

type GeneralProgress struct {
	Actual int `json:"actual"`
	Total  int `json:"total"`
}

func General_progress(user_id int) GeneralProgress {
	db := connection.Connect_db()
	var progress GeneralProgress

	err := db.QueryRow("SELECT sum(cost) FROM SolutionEfforts se, LevelsBrief lb WHERE se.level_id = lb.id AND user_id = ? AND is_successful = 1",
		user_id).
		Scan(&progress.Actual)
	if err != nil {
		panic(fmt.Sprint("Users table error: ", err.Error()))
	}

	err = db.QueryRow("SELECT sum(cost) FROM LevelsBrief").
		Scan(&progress.Total)
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
	return progress
}
