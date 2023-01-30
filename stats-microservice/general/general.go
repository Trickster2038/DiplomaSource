package general

import (
	"fmt"
	"stats/connection"
)

// select lb.id, lb.name, lb.seqnum, IFNULL(sum(is_successful), 0) as solutions from (SELECT * FROM LevelsBrief WHERE is_active = 1) lb LEFT JOIN SolutionEfforts se on lb.id = se.level_id GROUP BY lb.id ORDER BY seqnum;

type LevelPassed struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Seqnum    int    `json:"seqnum"`
	Solutions int    `json:"solutions"`
}

type AvgEfforts struct {
	ID         int     `json:"id"`
	Name       string  `json:"name"`
	Seqnum     int     `json:"seqnum"`
	AvgEfforts float64 `json:"avg_efforts"`
}

type SolutionCount struct {
	SolutionsCount int `json:"solutions_count"`
	LevelsCount    int `json:"levels_count"`
}

type ActivtyByMonth struct {
	Year           int `json:"year"`
	Month          int `json:"month"`
	SolutionsCount int `json:"solutions_count"`
	EffortsCount   int `json:"efforts_count"`
}

type TopUser struct {
	ID             int    `json:"id"`
	Nickname       string `json:"nickname"`
	SolutionsCount int    `json:"solutions_count"`
	EffortsCount   int    `json:"efforts_count"`
}

func Each_level_passed() []LevelPassed {
	var res []LevelPassed
	db := connection.Connect_db()

	results, err := db.Query("SELECT lb.id, lb.name, lb.seqnum, IFNULL(sum(is_successful), 0) as solutions " +
		"FROM (SELECT * FROM LevelsBrief WHERE is_active = 1) lb " +
		"LEFT JOIN SolutionEfforts se on lb.id = se.level_id " +
		"GROUP BY lb.id ORDER BY seqnum")
	if err != nil {
		panic(fmt.Sprintf("Getting levels statuses error in DB:", err.Error()))
	}

	var r LevelPassed

	for results.Next() {
		err = results.Scan(
			&r.ID,
			&r.Name,
			&r.Seqnum,
			&r.Solutions)
		if err != nil {
			panic(err.Error())
		}

		res = append(res, r)
	}

	defer db.Close()
	return res
}

func Each_level_avg_efforts() []AvgEfforts {
	var res []AvgEfforts
	db := connection.Connect_db()

	results, err := db.Query("SELECT lb.id, lb.name, lb.seqnum, IFNULL(IFNULL(sum(is_successful IS NOT NULL), 0)/IFNULL(sum(is_successful), 0), 0) as solutions " +
		"FROM (SELECT * FROM LevelsBrief WHERE is_active = 1) lb " +
		"LEFT JOIN SolutionEfforts se on lb.id = se.level_id " +
		"GROUP BY lb.id ORDER BY seqnum")
	if err != nil {
		panic(fmt.Sprintf("Getting levels avg efforts error in DB:", err.Error()))
	}

	var r AvgEfforts

	for results.Next() {
		err = results.Scan(
			&r.ID,
			&r.Name,
			&r.Seqnum,
			&r.AvgEfforts)
		if err != nil {
			panic(err.Error())
		}

		res = append(res, r)
	}

	defer db.Close()
	return res
}

func Solutions_count_distribution() []SolutionCount {
	var res []SolutionCount
	db := connection.Connect_db()

	results, err := db.Query("SELECT solutions, count(*) FROM (SELECT lb.id, lb.name, lb.seqnum, IFNULL(sum(is_successful), 0) as solutions " +
		"FROM (SELECT * FROM LevelsBrief WHERE is_active = 1) lb " +
		"LEFT JOIN SolutionEfforts se on lb.id = se.level_id " +
		"GROUP BY lb.id ORDER BY seqnum) tb " +
		"GROUP BY solutions ORDER BY solutions")
	if err != nil {
		panic(fmt.Sprintf("Getting solutions distribution error in DB:", err.Error()))
	}

	var r SolutionCount

	for results.Next() {
		err = results.Scan(
			&r.SolutionsCount,
			&r.LevelsCount)
		if err != nil {
			panic(err.Error())
		}

		res = append(res, r)
	}

	defer db.Close()
	return res
}

func Activity_by_month() []ActivtyByMonth {
	var res []ActivtyByMonth
	db := connection.Connect_db()

	results, err := db.Query("Select YEAR(time), MONTH(time), IFNULL(sum(is_successful), 0) " +
		"FROM SolutionEfforts GROUP BY MONTH(time), YEAR(time) " +
		"ORDER BY YEAR(time), MONTH(time)")
	if err != nil {
		panic(fmt.Sprintf("Getting activity by month error in DB:", err.Error()))
	}

	var r ActivtyByMonth

	for results.Next() {
		err = results.Scan(
			&r.Year, &r.Month, &r.SolutionsCount)
		if err != nil {
			panic(err.Error())
		}

		res = append(res, r)
	}

	i := 0

	results, err = db.Query("Select IFNULL(count(*), 0) " +
		"FROM SolutionEfforts GROUP BY MONTH(time), YEAR(time) " +
		"ORDER BY YEAR(time), MONTH(time)")
	if err != nil {
		panic(fmt.Sprintf("Getting activity by month error in DB:", err.Error()))
	}

	for results.Next() {
		err = results.Scan(
			&res[i].EffortsCount)
		if err != nil {
			panic(err.Error())
		}

		i++
	}

	defer db.Close()
	return res
}

func Top_last_month_active_users() []TopUser {
	var res []TopUser
	db := connection.Connect_db()

	results, err := db.Query("select u.id, u.nickname, IFNULL(sum(is_successful),0) as solutions, IFNULL(count(is_successful),0) " +
		"FROM Users u LEFT JOIN (SELECT * FROM SolutionEfforts WHERE time > CURRENT_TIMESTAMP() - 30*24*60*60*1000) se " +
		"ON u.id = se.user_id GROUP BY u.id ORDER BY solutions DESC;")
	if err != nil {
		panic(fmt.Sprintf("Getting solutions distribution error in DB:", err.Error()))
	}

	var r TopUser

	for results.Next() {
		err = results.Scan(
			&r.ID, &r.Nickname, &r.SolutionsCount, &r.EffortsCount)
		if err != nil {
			panic(err.Error())
		}

		res = append(res, r)
	}

	defer db.Close()
	return res
}
