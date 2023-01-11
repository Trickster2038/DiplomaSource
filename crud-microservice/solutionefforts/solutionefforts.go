package solutionefforts

import (
	"crud/connection"
)

type SolutionEffort struct {
	ID           int  `json:"id"`
	UserID       int  `json:"user_id"`
	LevelID      int  `json:"level_id"`
	IsSuccessful bool `json:"is_successful"`
	Time         int  `json:"time"`
}

func (solution_effort SolutionEffort) CheckSuccess(user_id int, level_id int) bool {
	db := connection.Connect_db()

	var count_success int
	err := db.QueryRow("SELECT count(*) FROM SolutionEfforts where user_id = ? AND level_id = ? AND is_successful = true",
		user_id, level_id).
		Scan(&count_success)
	if err != nil {
		panic(err.Error())
	}

	return (count_success > 0)
}

func (solution_effort SolutionEffort) Create() {
	db := connection.Connect_db()

	if solution_effort.CheckSuccess(solution_effort.UserID, solution_effort.LevelID) {
		panic("Level is already solved")
	}

	_, err := db.Query("INSERT INTO SolutionEfforts (user_id, level_id, is_successful) VALUES (?, ?, ?)",
		solution_effort.UserID,
		solution_effort.LevelID,
		solution_effort.IsSuccessful)
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
}
