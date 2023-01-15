package request

type ResponseFrame struct {
	StatusStr  string      `json:"status_str"`
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message,omitempty"`
	Data       interface{} `json:"data,omitempty"`
}

type InRequsestCRUD struct {
	UserID   int `json:"user_id"`
	MetaInfo struct {
		ObjType string `json:"obj_type"`
		Action  string `json:"action"`
	} `json:"metainfo"`
	Data interface{} `json:"data"`
}
