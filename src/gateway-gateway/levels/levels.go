package levels

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gateway/request"
	"io/ioutil"
	"net/http"
)

type LevelsCodeData struct {
	ID              int    `json:"id"`
	WideDescription string `json:"wide_description"`
	Code            string `json:"code"`
	Question        struct {
		Src string `json:"src"`
		Tb  string `json:"tb"`
	} `json:"question"`

	Answer string `json:"answer"`
}

type LevelsData struct {
	ID              int    `json:"id"`
	WideDescription string `json:"wide_description"`
	Code            string `json:"code"`
	Question        string `json:"question"`
	Answer          string `json:"answer"`
}

type InRequestLevelsCodeData struct {
	UserID   int              `json:"user_id"`
	MetaInfo request.MetaInfo `json:"metainfo"`
	Data     LevelsCodeData   `json:"data"`
}

type InRequestLevelsData struct {
	UserID   int              `json:"user_id"`
	MetaInfo request.MetaInfo `json:"metainfo"`
	Data     LevelsData       `json:"data"`
}

type OutRequestCompiler struct {
	UserID    int    `json:"user_id"`
	LevelID   int    `json:"level_id"`
	DeviceSrc string `json:"device_src"`
	TbSrc     string `json:"tb_src"`
}

type OutRequestParser struct {
	UserID  int         `json:"user_id"`
	LevelID int         `json:"level_id"`
	Data    interface{} `json:"data"`
}

type OutRequestWavedrom struct {
	Data interface{} `json:"data"`
}

type CodeLevelFlagFrame struct {
	IsCodeLevel bool `json:"is_code_level"`
}

// Returns string(wavedrom.json)
func Handle_code_level_data(user_id int, level_id int, src string, tb string) string {
	var out_req_src OutRequestCompiler
	out_req_src.UserID = user_id
	out_req_src.LevelID = level_id
	out_req_src.DeviceSrc = src
	out_req_src.TbSrc = tb

	payload, _ := json.Marshal(out_req_src)
	resp, err_post := http.Post("http://compiler-microservice:8080/build", "application/json", bytes.NewBuffer(payload))

	if err_post != nil {
		panic(fmt.Sprintf("Accesing Compiler-microservice error: %v", err_post.Error()))
	}

	var res request.ResponseFrame
	json.NewDecoder(resp.Body).Decode(&res)
	if res.StatusStr != "ok" {
		panic("Device synthesis error")
	} else {
		var out_req_parser OutRequestParser

		out_req_parser.UserID = user_id
		out_req_parser.LevelID = level_id
		out_req_parser.Data = res.Data

		payload, _ := json.Marshal(out_req_parser)
		resp, err_post := http.Post("http://parser-microservice:5000/parse", "application/json", bytes.NewBuffer(payload))

		if err_post != nil {
			panic(fmt.Sprintf("Accesing Parser-microservice error: %v", err_post.Error()))
		}

		var res request.ResponseFrame
		json.NewDecoder(resp.Body).Decode(&res)

		if res.StatusStr != "ok" {
			panic("Device parsing error")
		} else {
			var out_req_wd OutRequestWavedrom
			out_req_wd.Data = res.Data

			payload, _ := json.Marshal(out_req_wd)
			resp, err_post := http.Post("http://wavedrom-microservice:8081/wavedrom", "application/json", bytes.NewBuffer(payload))

			if err_post != nil {
				panic(fmt.Sprintf("Accesing Wavedrom-microservice error: %v", err_post.Error()))
			}

			var res request.ResponseFrame
			json.NewDecoder(resp.Body).Decode(&res)

			if res.StatusStr != "ok" {
				panic("Time diagram wavedroming error")
			} else {
				res_marshal, _ := json.Marshal(res.Data)
				return string(res_marshal)
			}
		}
	}

	return "Code handling unknown state"
}

func Crud_levels(w http.ResponseWriter, req *http.Request) {
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
	var dataFrame request.InRequestCRUD

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

	if (dataFrame.MetaInfo.ObjType != "levels_brief") &&
		(dataFrame.MetaInfo.ObjType != "levels_data") {
		panic(fmt.Sprintf("Wrong CRUD-Object Type: %s", dataFrame.MetaInfo.ObjType))
	}

	if (dataFrame.MetaInfo.ObjType == "levels_data") &&
		(dataFrame.MetaInfo.Action == "delete") {
		panic("LevelsData can be only ARCHIVED by archivating LevelsBrief")
	}

	if (dataFrame.MetaInfo.Action == "create") ||
		(dataFrame.MetaInfo.Action == "update") ||
		(dataFrame.MetaInfo.Action == "delete") {
		var payloadFrame request.IdRequestFrame
		payloadFrame.MetaInfo.ObjType = "user"
		payloadFrame.MetaInfo.Action = "read"
		payloadFrame.Data.ID = dataFrame.UserID

		payload, err := json.Marshal(payloadFrame)
		resp, err_post := http.Post("http://crud-microservice:8082/crud", "application/json", bytes.NewBuffer(payload))

		if err_post != nil {
			panic(fmt.Sprintf("Accesing CRUD-microservice.User error: %v", err.Error()))
		}

		var res request.AdminFlagFrame
		json.NewDecoder(resp.Body).Decode(&res)

		if !res.Data.IsAdmin {
			panic("User have no rights to modify levels")
		}
	}

	var code_frame CodeLevelFlagFrame
	err = json.Unmarshal(reqBody, &code_frame)

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

	var payload []byte

	if !code_frame.IsCodeLevel ||
		!(((dataFrame.MetaInfo.Action == "create") || (dataFrame.MetaInfo.Action == "update")) &&
			dataFrame.MetaInfo.ObjType == "levels_data") {
		payload, err = json.Marshal(dataFrame)
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
	} else {
		var code_level_data InRequestLevelsCodeData
		err := json.Unmarshal(reqBody, &code_level_data)
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
			panic("JSON (LevelData) parsing error")
		}

		var code_level_data_str InRequestLevelsData
		code_level_data_str.MetaInfo = code_level_data.MetaInfo
		code_level_data_str.UserID = code_level_data.UserID
		code_level_data_str.Data.ID = code_level_data.Data.ID
		code_level_data_str.Data.WideDescription = code_level_data.Data.WideDescription
		code_level_data_str.Data.Code = code_level_data.Data.Code

		/*
			TODO [NOT]:
			Gateway Request - question: {src, tb}, answer: ""
			CRUD Request - question: tb, answer: wavedrom
		*/
		code_level_data_str.Data.Question = code_level_data.Data.Question.Tb
		code_level_data_str.Data.Answer = Handle_code_level_data(code_level_data.UserID,
			code_level_data.Data.ID,
			code_level_data.Data.Question.Src,
			code_level_data.Data.Question.Tb)

		payload, err = json.Marshal(code_level_data_str)
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
	}

	resp, err_post := http.Post("http://crud-microservice:8082/crud", "application/json", bytes.NewBuffer(payload))

	if err_post != nil {
		panic(fmt.Sprintf("Accesing CRUD-microservice.Levels error: %v", err.Error()))
	}

	var res request.ResponseFrame
	json.NewDecoder(resp.Body).Decode(&res)

	response.StatusStr = res.StatusStr
	response.StatusCode = res.StatusCode
	response.Data = res.Data
	if res.StatusStr == "ok" {
		response.Message = "CRUD operation with Level done"
	} else {
		response.Message = "CRUD-microservice.LevelsX error: " + res.Message
	}
	w.WriteHeader(response.StatusCode)
	json.NewEncoder(w).Encode(response)

}
