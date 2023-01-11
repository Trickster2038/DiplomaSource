package main

// TODO:
// - Web-CRUD handlers
// - SolutionEfforts ORM

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os/exec"
	"runtime"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"

	"crud/levelsbrief"
	"crud/levelsdata"
	"crud/solutionefforts"
	"crud/typerecord"
	"crud/user"
)

type MetaInfo struct {
	Obj_type string `json:"obj_type"`
	Action   string `json:"action"`
}

type EncapsuledMetaInfo struct {
	MetaInfo MetaInfo `json:"MetaInfo"`
}

type RfSolutionEffort struct {
	MetaInfo MetaInfo                       `json:"MetaInfo"`
	Data     solutionefforts.SolutionEffort `json:"data"`
}

type RfLevelsBrief struct {
	MetaInfo MetaInfo                `json:"MetaInfo"`
	Data     levelsbrief.LevelsBrief `json:"data"`
}

type RfLevelsData struct {
	MetaInfo MetaInfo              `json:"MetaInfo"`
	Data     levelsdata.LevelsData `json:"data"`
}

type RfTypeRecord struct {
	MetaInfo MetaInfo        `json:"MetaInfo"`
	Data     typerecord.Type `json:"data"`
}

type RfUser struct {
	MetaInfo MetaInfo  `json:"MetaInfo"`
	Data     user.User `json:"data"`
}

func (RfSe RfSolutionEffort) Create() {
	RfSe.Data.Create()
}

type ICreatable interface {
	Create()
}

type IUpdatable interface {
	Update()
}

type IDeletable interface {
	Delete()
}

func Create(v ICreatable) {
	v.Create()
}

func Update(v IUpdatable) {
	v.Update()
}

func Delete(v IDeletable) {
	v.Delete()
}

type ResponseFrame struct {
	Status_str  string      `json:"status_str"`
	Status_code int         `json:"status_code"`
	Message     string      `json:"message,omitempty"`
	Data        interface{} `json:"data,omitempty"`
}

func crud(w http.ResponseWriter, req *http.Request) {
	var response ResponseFrame

	defer func() {
		if panicInfo := recover(); panicInfo != nil {
			var response ResponseFrame
			response.Status_str = "error"
			response.Status_code = 400
			response.Message = fmt.Sprintf("Top-level panic: %v", panicInfo)
			w.WriteHeader(response.Status_code)
			json.NewEncoder(w).Encode(response)
		}
	}()

	reqBody, _ := ioutil.ReadAll(req.Body)
	var reqFrame EncapsuledMetaInfo

	err := json.Unmarshal(reqBody, &reqFrame)

	if err != nil {
		defer func() {
			if r := recover(); r != nil {
				response.Status_str = "error"
				response.Status_code = 400
				response.Message = err.Error()
				w.WriteHeader(response.Status_code)
				json.NewEncoder(w).Encode(response)
			}
		}()
		panic("request-JSON parsing error")
	}

	var data interface{}
	if reqFrame.MetaInfo.Obj_type == "solutioneffort" {
		data = &RfSolutionEffort{}
		err := json.Unmarshal(reqBody, data)
		if err != nil {
			defer func() {
				if r := recover(); r != nil {
					response.Status_str = "error"
					response.Status_code = 400
					response.Message = err.Error()
					w.WriteHeader(response.Status_code)
					json.NewEncoder(w).Encode(response)
				}
			}()
			panic("data-JSON parsing error")
		}
	}

	if reqFrame.MetaInfo.Action == "create" {
		// Create(data.(CreatableStruct))
		data.(ICreatable).Create()
	} else if reqFrame.MetaInfo.Action == "update" {
		data.(IUpdatable).Update()
	} else if reqFrame.MetaInfo.Action == "delete" {
		data.(IDeletable).Delete()
	}

	response.Status_str = "ok"
	response.Status_code = 200
	msg, _ := json.Marshal(reqFrame)
	response.Data = string(msg)
	w.WriteHeader(response.Status_code)
	json.NewEncoder(w).Encode(response)
}

func main() {
	if runtime.GOOS == "windows" {
		fmt.Println("Can't Execute this on a windows machine")
	} else {
		fmt.Println("Server start")
		r := mux.NewRouter().StrictSlash(true)
		r.HandleFunc("/crud", crud).Methods("POST")

		// level := LevelsBrief{0, 1, 2, 3, true, "Level1", "Test level"}

		var lvl_br levelsbrief.LevelsBrief
		lvl_br.Read(1)
		fmt.Println("Lvl name: ", lvl_br.Name)

		var level_data levelsdata.LevelsData
		level_data.Read(1)
		res, _ := json.Marshal(level_data)
		fmt.Println(string(res))

		level_data = levelsdata.LevelsData{ID: 2, WideDescription: "1tttt", Code: "2t", Question: "3t", Answer: "4t"}
		// level_data.CreateOrUpdate()

		level_data.Update()

		var eff solutionefforts.SolutionEffort
		eff.UserID = 1
		eff.LevelID = 7
		eff.IsSuccessful = true

		// eff.Create()

		exec.Command("fuser", "-k", "8082/tcp").Output()
		http.ListenAndServe(":8082", r)
		fmt.Println("Server stop")
	}
}
