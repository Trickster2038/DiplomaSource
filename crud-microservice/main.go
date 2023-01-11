package main

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
	MetaInfo MetaInfo `json:"metainfo"`
}

type LevelsList struct {
	Levels []levelsbrief.LevelsBrief `json:"levels"`
}

type EncapsulatedSuccesfull struct {
	Succesful bool `json:"succesfull"`
}

type RfLevelsBrief struct {
	MetaInfo MetaInfo                `json:"MetaInfo"`
	Data     levelsbrief.LevelsBrief `json:"data"`
}

type RfLevelsData struct {
	MetaInfo MetaInfo              `json:"MetaInfo"`
	Data     levelsdata.LevelsData `json:"data"`
}

type RfSolutionEffort struct {
	MetaInfo MetaInfo                       `json:"MetaInfo"`
	Data     solutionefforts.SolutionEffort `json:"data"`
}

type RfTypeRecord struct {
	MetaInfo MetaInfo        `json:"MetaInfo"`
	Data     typerecord.Type `json:"data"`
}

type RfUser struct {
	MetaInfo MetaInfo  `json:"MetaInfo"`
	Data     user.User `json:"data"`
}

func (v RfLevelsBrief) Create() {
	v.Data.Create()
}

func (v RfLevelsBrief) Read() interface{} {
	v.Data.Read()
	return v.Data
}

func (v RfLevelsBrief) ReadAll() interface{} {
	var levels LevelsList
	levels.Levels = v.Data.ReadAll()
	return levels
}

func (v RfLevelsBrief) Update() {
	v.Data.Update()
}

func (v RfLevelsBrief) Delete() {
	v.Data.Delete()
}

func (v RfLevelsData) Create() {
	v.Data.Create()
}

func (v RfLevelsData) Read() interface{} {
	v.Data.Read()
	return v.Data
}

func (v RfLevelsData) Update() {
	v.Data.Update()
}

func (v RfLevelsData) Delete() {
	v.Data.Delete()
}

func (v RfSolutionEffort) Create() {
	v.Data.Create()
}

func (v RfSolutionEffort) CheckSuccessful() interface{} {
	var m EncapsulatedSuccesfull
	m.Succesful = v.Data.CheckSuccessful()
	return m
}

func (v RfTypeRecord) Read() interface{} {
	v.Data.Read()
	return v.Data
}

func (v RfUser) Read() interface{} {
	v.Data.Read()
	return v.Data
}

type ICreatable interface {
	Create()
}

type IReadable interface {
	Read() interface{}
}

type IReadableAll interface {
	ReadAll() interface{}
}

type IUpdatable interface {
	Update()
}

type IDeletable interface {
	Delete()
}

type ICheackableSuccessful interface {
	CheckSuccessful() interface{}
}

func Create(v ICreatable) {
	v.Create()
}

func Read(v IReadable) {
	v.Read()
}

func ReadAll(v IReadableAll) {
	v.ReadAll()
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
	var err error

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

	err = json.Unmarshal(reqBody, &reqFrame)

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

	if reqFrame.MetaInfo.Obj_type == "levels_brief" {
		data = &RfLevelsBrief{}
		err = json.Unmarshal(reqBody, data)
	} else if reqFrame.MetaInfo.Obj_type == "levels_data" {
		data = &RfLevelsData{}
		err = json.Unmarshal(reqBody, data)
	} else if reqFrame.MetaInfo.Obj_type == "type_record" {
		data = &RfTypeRecord{}
		err = json.Unmarshal(reqBody, data)
	} else if reqFrame.MetaInfo.Obj_type == "solution_effort" {
		data = &RfSolutionEffort{}
		err = json.Unmarshal(reqBody, data)
	} else if reqFrame.MetaInfo.Obj_type == "user" {
		data = &RfUser{}
		err = json.Unmarshal(reqBody, data)
	} else {
		panic("Unknown Obj Type")
	}

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

	if reqFrame.MetaInfo.Action == "create" {
		data.(ICreatable).Create()
	} else if reqFrame.MetaInfo.Action == "read" {
		response.Data = data.(IReadable).Read()
	} else if reqFrame.MetaInfo.Action == "update" {
		data.(IUpdatable).Update()
	} else if reqFrame.MetaInfo.Action == "delete" {
		data.(IDeletable).Delete()
	} else if reqFrame.MetaInfo.Action == "read_all" {
		response.Data = data.(IReadableAll).ReadAll()
	} else if reqFrame.MetaInfo.Action == "check_succesful" {
		response.Data = data.(ICheackableSuccessful).CheckSuccessful()
	} else {
		panic("Unknown Action")
	}

	response.Status_str = "ok"
	response.Status_code = 200
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

		exec.Command("fuser", "-k", "8082/tcp").Output()
		http.ListenAndServe(":8082", r)

		fmt.Println("Server stop")
	}
}
