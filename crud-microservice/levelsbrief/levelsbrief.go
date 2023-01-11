package levelsbrief

import (
	"crud/connection"
)

type LevelsBrief struct {
	ID         int    `json:"id"`
	Level_type int    `json:"level_type"`
	Seqnum     int    `json:"seqnum"`
	Cost       int    `json:"cost"`
	Is_active  bool   `json:"is_active"`
	Name       string `json:"name"`
	Brief      string `json:"brief"`
}

// returns MaxId (current for creating LevelsData)
func (level_brief LevelsBrief) Create() int {
	db := connection.Connect_db()
	_, err := db.Query("UPDATE LevelsBrief SET seqnum = seqnum + 1 WHERE seqnum >= ?",
		level_brief.Seqnum)
	if err != nil {
		panic(err.Error())
	}

	_, err = db.Query("INSERT INTO LevelsBrief (level_type, seqnum, cost, is_active, name, brief) VALUES (?, ?, ?, ?, ?, ?)",
		level_brief.Level_type,
		level_brief.Seqnum,
		level_brief.Cost,
		level_brief.Is_active,
		level_brief.Name,
		level_brief.Brief)
	if err != nil {
		panic(err.Error())
	}

	var maxId int
	err = db.QueryRow("SELECT max(id) FROM LevelsBrief").Scan(&maxId)
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	return maxId
}

func (level_brief *LevelsBrief) Read(id int) {
	db := connection.Connect_db()
	err := db.QueryRow("SELECT id, level_type, seqnum, cost, is_active, name, brief FROM LevelsBrief where id = ?",
		id).
		Scan(&level_brief.ID,
			&level_brief.Level_type,
			&level_brief.Seqnum,
			&level_brief.Cost,
			&level_brief.Is_active,
			&level_brief.Name,
			&level_brief.Brief)
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
}

func (level_brief LevelsBrief) Update() {
	db := connection.Connect_db()

	var old_level_brief LevelsBrief

	old_level_brief.Read(level_brief.ID) // can call panic

	if level_brief.Seqnum > old_level_brief.Seqnum {
		_, err := db.Query("UPDATE LevelsBrief SET seqnum = seqnum - 1 WHERE seqnum > ? AND seqnum < ?",
			old_level_brief.Seqnum, level_brief.Seqnum)
		if err != nil {
			panic(err.Error())
		}
	} else if level_brief.Seqnum < old_level_brief.Seqnum {
		_, err := db.Query("UPDATE LevelsBrief SET seqnum = seqnum + 1 WHERE seqnum > ? AND seqnum < ?",
			level_brief.Seqnum, old_level_brief.Seqnum)
		if err != nil {
			panic(err.Error())
		}
	}

	_, err := db.Query("UPDATE LevelsBrief SET "+
		"level_type = ?, "+
		"seqnum = ?, "+
		"cost = ?, "+
		"is_active = ?, "+
		"name = ?, "+
		"brief = ? "+
		"WHERE id = ?",
		level_brief.Level_type,
		level_brief.Seqnum,
		level_brief.Cost,
		level_brief.Is_active,
		level_brief.Name,
		level_brief.Brief,
		level_brief.ID)
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
}

// setting is_active = FALSE
func (level_brief LevelsBrief) Delete() {
	db := connection.Connect_db()

	level_brief.Read(level_brief.ID) // can call panic
	if level_brief.Is_active {
		_, err := db.Query("UPDATE LevelsBrief SET seqnum = seqnum - 1 WHERE seqnum > ?",
			level_brief.Seqnum)
		if err != nil {
			panic(err.Error())
		}

		_, err = db.Query("UPDATE LevelsBrief SET is_active = false WHERE id = ?",
			level_brief.ID)
		if err != nil {
			panic(err.Error())
		}
	} else {
		panic("Level is already deleted (archived)")
	}

	defer db.Close()
}

func (level_brief LevelsBrief) ReadAll() []LevelsBrief {

	var levels_brief_array []LevelsBrief

	db := connection.Connect_db()

	results, err := db.Query("SELECT id, level_type, seqnum, cost, is_active, name, brief FROM LevelsBrief")
	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		var level_brief LevelsBrief
		err = results.Scan(&level_brief.ID,
			&level_brief.Level_type,
			&level_brief.Seqnum,
			&level_brief.Cost,
			&level_brief.Is_active,
			&level_brief.Name,
			&level_brief.Brief)
		if err != nil {
			panic(err.Error())
		}
		levels_brief_array = append(levels_brief_array, level_brief)
	}

	defer db.Close()

	return levels_brief_array
}
