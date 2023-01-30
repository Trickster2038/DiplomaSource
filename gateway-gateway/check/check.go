package check

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gateway/levels"
	"gateway/request"
	"gateway/user"
	"io/ioutil"
	"net/http"
	"strings"
)

// TODO:
// - read stats

type RequestFrame struct {
	UserID  int    `json:"user_id"`
	LevelID int    `json:"level_id"`
	Answer  string `json:"answer"`
}

type CheckSuccessfulUniversal struct {
	StatusStr  string           `json:"status_str"`
	StatusCode int              `json:"status_code"`
	Message    string           `json:"message"`
	MetaInfo   request.MetaInfo `json:"metainfo"`
	Data       struct {
		UserID       int  `json:"user_id"`
		LevelID      int  `json:"level_id"`
		IsSuccessful bool `json:"is_successful"`
	} `json:"data"`
}

// === Analyzer DataFormats ===

type AnalyzerResponseFrame struct {
	StatusStr       string      `json:"status_str"`
	StatusCode      int         `json:"status_code"`
	Message         string      `json:"message"`
	IsCorrect       bool        `json:"is_correct"`
	IsAlreadySolved bool        `json:"is_already_solved"` // additional, not inherited
	Data            interface{} `json:"data"`
}

type SingleChoiceTask struct {
	CorrectAnswerID int `json:"correct_answer_id"`
	Answers         []struct {
		Text string `json:"text"`
		Hint string `json:"hint"`
	} `json:"answers"`
}

type SingleChoiceTestRequest struct {
	Type string `json:"type"`
	Data struct {
		UserAnswerID int              `json:"user_answer_id"`
		Task         SingleChoiceTask `json:"task"`
	} `json:"data"`
}

type MultiChoiceTestRequest struct {
	Type string `json:"type"`
	Data struct {
		UserAnswers []bool `json:"user_answers"`
		Task        struct {
			CorrectAnswers []bool `json:"correct_answers"`
		} `json:"task"`
	} `json:"data"`
}

type WavedromSignal struct {
	Name string   `json:"name"`
	Wave string   `json:"wave"`
	Data []string `json:"data"`
}

type CodeRequest struct {
	Type string `json:"type"`
	Data struct {
		UserSignals    []WavedromSignal `json:"user_signals"`
		CorrectSignals []WavedromSignal `json:"correct_signals"`
	} `json:"data"`
}

func analyze(payload []byte) AnalyzerResponseFrame {
	var res AnalyzerResponseFrame

	resp, err_post := http.Post("http://analyzer-microservice:8083/check", "application/json", bytes.NewBuffer(payload))

	if err_post != nil {
		panic(fmt.Sprintf("Accesing analyzer-microservice error: %v", err_post.Error()))
	}

	json.NewDecoder(resp.Body).Decode(&res)
	if res.StatusStr != "ok" {
		panic(fmt.Sprintf("analyzer-microservice error: %v", &res.Message))
	} else {
		res.IsAlreadySolved = false
		return res
	}

	return res
}

func level_already_solved(level_id int, user_id int) bool {
	var req CheckSuccessfulUniversal
	req.MetaInfo.ObjType = "solution_effort"
	req.MetaInfo.Action = "check_successful"
	req.Data.LevelID = level_id
	req.Data.UserID = user_id

	payload, _ := json.Marshal(req)

	resp, err_post := http.Post("http://crud-microservice:8082/crud", "application/json", bytes.NewBuffer(payload))

	if err_post != nil {
		panic(fmt.Sprintf("Accesing CRUD-microservice.Users error: %v", err_post.Error()))
	}

	var res CheckSuccessfulUniversal

	json.NewDecoder(resp.Body).Decode(&res)
	if res.StatusStr != "ok" {
		panic(fmt.Sprintf("Error, checking if solution already exists"))
	} else {
		if res.Data.IsSuccessful == true {
			return true
		}
	}

	return false
}

func write_solution_to_stats(level_id int, user_id int, is_correct bool) {
	var req CheckSuccessfulUniversal
	req.Data.LevelID = level_id
	req.Data.UserID = user_id
	req.Data.IsSuccessful = is_correct
	req.MetaInfo.ObjType = "solution_effort"
	req.MetaInfo.Action = "create"

	payload, _ := json.Marshal(req)

	_, err_post := http.Post("http://crud-microservice:8082/crud", "application/json", bytes.NewBuffer(payload))

	if err_post != nil {
		panic(fmt.Sprintf("Accesing CRUD-microservice.SolutionEfforts error: %v", err_post.Error()))
	}
}

func Check(w http.ResponseWriter, req *http.Request) {
	var response request.ResponseFrame

	defer func() {
		if panicInfo := recover(); panicInfo != nil {
			var response request.ResponseFrame
			response.StatusStr = "error"
			response.StatusCode = 400
			response.Message = fmt.Sprintf("Top-level panic: %v", panicInfo)
			w.WriteHeader(response.StatusCode)
			json.NewEncoder(w).Encode(response)
		}
	}()

	reqBody, _ := ioutil.ReadAll(req.Body)
	var dataFrame RequestFrame

	err := json.Unmarshal(reqBody, &dataFrame)

	if err != nil {
		defer func() {
			if r := recover(); r != nil {
				response.StatusStr = "error"
				response.StatusCode = 400
				response.Message = err.Error()
				w.WriteHeader(response.StatusCode)
				json.NewEncoder(w).Encode(response)
			}
		}()
		panic("JSON (Request) parsing error")
	}

	user.Check_user_exists(dataFrame.UserID)
	if level_already_solved(dataFrame.LevelID, dataFrame.UserID) {

		var res AnalyzerResponseFrame
		res.StatusCode = 200
		res.StatusStr = "ok"
		res.IsCorrect = true //FIXME[NOT]: Attention! Not NULL
		res.IsAlreadySolved = true
		res.Data = nil
		w.WriteHeader(res.StatusCode)
		json.NewEncoder(w).Encode(res)
	} else {

		var level_brief request.IdRequestFrame
		level_brief.MetaInfo.ObjType = "levels_brief"
		level_brief.MetaInfo.Action = "read"
		level_brief.Data.ID = dataFrame.LevelID

		payload, _ := json.Marshal(level_brief)
		resp, err_post := http.Post("http://crud-microservice:8082/crud", "application/json", bytes.NewBuffer(payload))

		if err_post != nil {
			panic(fmt.Sprintf("Accesing CRUD-microservice.LevelsBrief error: %v", err_post.Error()))
		}

		var res request.ResponseFrame
		json.NewDecoder(resp.Body).Decode(&res)
		if res.StatusStr != "ok" {
			panic(fmt.Sprintf("CRUD-microservice.LevelsBrief error: %s", &res.Message))
		} else {
			level_type_name := res.Data.(map[string]interface{})["level_type_name"].(string)

			level_brief.MetaInfo.ObjType = "levels_data"

			payload, _ := json.Marshal(level_brief)
			resp, err_post := http.Post("http://crud-microservice:8082/crud", "application/json", bytes.NewBuffer(payload))

			if err_post != nil {
				panic(fmt.Sprintf("Accesing CRUD-microservice.LevelsBrief error: %v", err_post.Error()))
			}

			var res request.ResponseFrame
			json.NewDecoder(resp.Body).Decode(&res)
			if res.StatusStr != "ok" {
				panic(fmt.Sprintf("CRUD-microservice.LevelsData error: %s", &res.Message))
			} else {
				level_question := strings.Replace(
					res.Data.(map[string]interface{})["question"].(string),
					"\\", "", -1)
				level_answer := strings.Replace(
					res.Data.(map[string]interface{})["answer"].(string),
					"\\", "", -1)
				user_answer := strings.Replace(
					dataFrame.Answer,
					"\\", "", -1)

				var payload []byte
				if level_type_name == "program" {
					var data CodeRequest
					data.Type = level_type_name

					user_signals := levels.Handle_code_level_data(dataFrame.UserID,
						dataFrame.LevelID, user_answer, level_question)

					user_signals = strings.Replace(user_signals, "\\", "", -1)

					err = json.Unmarshal([]byte(user_signals), &data.Data.UserSignals)
					if err != nil {
						panic(fmt.Sprintf("Compiling user device error: %v", user_signals))
					}
					err = json.Unmarshal([]byte(level_answer), &data.Data.CorrectSignals)
					if err != nil {
						panic("Parsing correct signals error")
					}

					payload, _ = json.Marshal(data)
				} else if level_type_name == "singlechoice_test" {
					var data SingleChoiceTestRequest
					var keyval_int_json map[string]int

					data.Type = level_type_name

					json.Unmarshal([]byte(user_answer), &keyval_int_json)
					data.Data.UserAnswerID = keyval_int_json["user_answer_id"]

					json.Unmarshal([]byte(level_question), &data.Data.Task)

					json.Unmarshal([]byte(level_answer), &keyval_int_json)
					data.Data.Task.CorrectAnswerID = keyval_int_json["correct_answer_id"]

					payload, _ = json.Marshal(data)
				} else if level_type_name == "multichoice_test" {
					var data MultiChoiceTestRequest

					data.Type = level_type_name
					json.Unmarshal([]byte(level_answer), &data.Data.Task)
					json.Unmarshal([]byte(user_answer), &data.Data)

					payload, _ = json.Marshal(data)
				} else if level_type_name == "text" {
					panic("Cannot check level of text type")
				} else {
					panic("Unknown level type name")
				}

				var res AnalyzerResponseFrame
				res = analyze(payload)

				write_solution_to_stats(dataFrame.LevelID, dataFrame.UserID, res.IsCorrect)

				w.WriteHeader(res.StatusCode)
				json.NewEncoder(w).Encode(res)

			}
		}
	}
}
