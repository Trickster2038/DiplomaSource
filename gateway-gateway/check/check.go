package check

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gateway/request"
	"io/ioutil"
	"net/http"
	"strings"
)

// TODO:
// - check if already solved
// - write to stats

type RequestFrame struct {
	UserID  int    `json:"user_id"`
	LevelID int    `json:"level_id"`
	Answer  string `json:"answer"`
}

type LevelsBriefRequest struct {
	MetaInfo request.MetaInfo `json:"metainfo"`
	Data     struct {
		ID int `json:"id"`
	} `json:"data"`
}

// === Analyzer DataFormats ===

type AnalyzerResponseFrame struct {
	StatusStr  string      `json:"status_str"`
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	IsCorrect  bool        `json:"is_correct"`
	Data       interface{} `json:"data"`
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
		return res
	}

	return res
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

	var level_brief LevelsBriefRequest
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

			// var data interface{}

			//TODO:
			if level_type_name == "program" {
				// TODO:
				var data CodeRequest
				err = json.Unmarshal([]byte(user_answer), data.Data.UserSignals)
				err = json.Unmarshal([]byte(level_answer), data.Data.CorrectSignals)
			} else if level_type_name == "singlechoice_test" {
				var data SingleChoiceTestRequest
				var keyval_int_json map[string]int

				data.Type = level_type_name

				json.Unmarshal([]byte(user_answer), &keyval_int_json)
				data.Data.UserAnswerID = keyval_int_json["user_answer_id"]

				err = json.Unmarshal([]byte(level_question), &data.Data.Task)
				json.Unmarshal([]byte(level_answer), &keyval_int_json)
				data.Data.Task.CorrectAnswerID = keyval_int_json["correct_answer_id"]

				payload, _ = json.Marshal(data)
				var res AnalyzerResponseFrame

				// panic(string(payload))

				res = analyze(payload)
				w.WriteHeader(res.StatusCode)
				json.NewEncoder(w).Encode(res)

				// panic(fmt.Sprintf("%v", data))
			} else if level_type_name == "multichoice_test" {
				// TODO:
				// var data MultiChoiceTestRequest
			} else if level_type_name == "text" {
				panic("Cannot check level of text type")
			} else {
				panic("Unknown level type name")
			}

			// err = json.Unmarshal(reqBody, data)

			// if err != nil {
			// 	defer func() {
			// 		if r := recover(); r != nil {
			// 			response.StatusStr = "error"
			// 			response.StatusCode = 400
			// 			response.Message = err.Error()
			// 			w.WriteHeader(response.StatusCode)
			// 			json.NewEncoder(w).Encode(response)
			// 		}
			// 	}()
			// 	panic("JSON (user answer) parsing error")
			// }

			// //TODO: further processing
			// panic(level_type_name + "|||" + level_question + "|||" + level_answer)
		}

	}

}
