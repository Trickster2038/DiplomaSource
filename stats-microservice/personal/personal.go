package personal

import (
	"fmt"
	"stats/connection"
)

// TODO: check if user does not exist on GATEWAY

type GeneralProgress struct {
	Actual int `json:"actual"`
	Total  int `json:"total"`
}

type LevelStatus struct {
	ID           int    `json:"id"`
	Seqnum       int    `json:"seqnum"`
	Cost         int    `json:"cost"`
	LevelName    string `json:"level_name"`
	Brief        string `json:"brief"`
	IsSuccessful bool   `json:"is_succesful"`
	LevelType    string `json:"level_type"`
}

type AverageEffortsStruct struct {
	AverageEfforts float64 `json:"avg_efforts"`
}

type MonthlyActivity struct {
	AcceptedTasks int `json: "accepted_tasks"`
	EarnedPoints  int `json: "earned_points"`
	Efforts       int `json: "efforts"`
}

type ActivityBorders struct {
	FirstEffort string `json: "first_efforts"`
	FirstSolved string `json: "first_solved"`
	LastEffort  string `json: "last_efforts"`
	LastSolved  string `json: "last_solved"`
}

func General_progress(user_id int) GeneralProgress {
	db := connection.Connect_db()
	var progress GeneralProgress

	err := db.QueryRow("SELECT sum(cost) FROM SolutionEfforts se, LevelsBrief lb WHERE se.level_id = lb.id AND user_id = ? AND is_successful = 1 AND lb.is_active = 1",
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

func Levels_statuses(user_id int) []LevelStatus {
	var res []LevelStatus
	db := connection.Connect_db()

	results, err := db.Query("SELECT lb.id, lb.seqnum, cost, level_name, brief, IFNULL(is_successful, 0), tp.name as level_type "+
		"FROM (SELECT id, level_type,seqnum, cost, name as level_name, brief FROM LevelsBrief WHERE is_active = 1) lb "+
		"LEFT JOIN (SELECT user_id, is_successful, level_id FROM SolutionEfforts WHERE user_id = ? AND is_successful = 1) se "+
		"ON lb.id = level_id JOIN Types tp ON level_type = tp.id "+
		"ORDER BY SEQNUM", user_id)
	if err != nil {
		panic(fmt.Sprintf("Getting levels statuses error in DB:", err.Error()))
	}

	var r LevelStatus

	for results.Next() {
		err = results.Scan(&r.ID,
			&r.Seqnum,
			&r.Cost,
			&r.LevelName,
			&r.Brief,
			&r.IsSuccessful,
			&r.LevelType)
		if err != nil {
			panic(err.Error())
		}

		res = append(res, r)
	}

	defer db.Close()
	return res
}

func Average_efforts_per_level(user_id int) AverageEffortsStruct {
	var res AverageEffortsStruct
	db := connection.Connect_db()

	err := db.QueryRow("SELECT IFNULL(c1/c2, -1.0) FROM "+
		"(SELECT count(*) as c1 FROM SolutionEfforts WHERE user_id = ?) tb1, "+
		"(SELECT count(*) as c2 FROM SolutionEfforts WHERE user_id = ? AND is_successful = 1) tb2",
		user_id, user_id).
		Scan(&res.AverageEfforts)
	if err != nil {
		panic(fmt.Sprint("Getting avg efforts per level DB error: ", err.Error()))
	}

	defer db.Close()
	return res
}

func Monthly_activity(user_id int) MonthlyActivity {
	var res MonthlyActivity
	db := connection.Connect_db()

	err := db.QueryRow("SELECT count(*) FROM SolutionEfforts WHERE user_id = ? AND time > CURRENT_TIMESTAMP() - 30*24*60*60*1000",
		user_id).
		Scan(&res.Efforts)
	if err != nil {
		panic(fmt.Sprint("Getting monthly activity DB error: ", err.Error()))
	}

	err = db.QueryRow("SELECT count(*), IFNULL(sum(cost), 0) FROM SolutionEfforts se JOIN LevelsBrief lb ON level_id = lb.id "+
		"WHERE is_successful = 1 AND user_id = ? AND time > CURRENT_TIMESTAMP() - 30*24*60*60*1000;",
		user_id).
		Scan(&res.AcceptedTasks, &res.EarnedPoints)
	if err != nil {
		panic(fmt.Sprint("Getting monthly activity DB error: ", err.Error()))
	}

	defer db.Close()
	return res
}

func Activity_borders(user_id int) ActivityBorders {
	var res ActivityBorders
	db := connection.Connect_db()

	err := db.QueryRow("SELECT IFNULL(min(time), -1) FROM SolutionEfforts WHERE user_id = ?",
		user_id).
		Scan(&res.FirstEffort)
	if err != nil {
		panic(fmt.Sprint("Getting activity borders DB error: ", err.Error()))
	}

	err = db.QueryRow("SELECT IFNULL(min(time), -1) FROM SolutionEfforts WHERE is_successful = 1 AND user_id = ?",
		user_id).
		Scan(&res.FirstSolved)
	if err != nil {
		panic(fmt.Sprint("Getting activity borders DB error: ", err.Error()))
	}

	err = db.QueryRow("SELECT IFNULL(max(time), -1) FROM SolutionEfforts WHERE user_id = ?",
		user_id).
		Scan(&res.LastEffort)
	if err != nil {
		panic(fmt.Sprint("Getting activity borders DB error: ", err.Error()))
	}

	err = db.QueryRow("SELECT IFNULL(max(time), -1) FROM SolutionEfforts WHERE is_successful = 1 AND user_id = ?",
		user_id).
		Scan(&res.LastSolved)
	if err != nil {
		panic(fmt.Sprint("Getting activity borders DB error: ", err.Error()))
	}

	defer db.Close()
	return res
}
