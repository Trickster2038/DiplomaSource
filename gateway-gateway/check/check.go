package check

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gateway/request"
	"io/ioutil"
	"net/http"
)

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
		panic("JSON parsing error")
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

	var resRF request.ResponseFrame
	json.NewDecoder(resp.Body).Decode(&resRF)
	if resRF.StatusStr != "ok" {
		panic(fmt.Sprintf("CRUD-microservice.LevelsBrief error: %s", &resRF.Message))
	} else {
		level_type_name := resRF.Data.(map[string]interface{})["level_type_name"].(string)

		//TODO: further processing
		panic(level_type_name)
	}

}
