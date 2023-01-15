package request

type ResponseFrame struct {
	StatusStr  string      `json:"status_str"`
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message,omitempty"`
	Data       interface{} `json:"data,omitempty"`
}

type MetaInfo struct {
	ObjType string `json:"obj_type"`
	Action  string `json:"action"`
}

type InRequestCRUD struct {
	UserID   int         `json:"user_id"`
	MetaInfo MetaInfo    `json:"metainfo"`
	Data     interface{} `json:"data"`
}

type IdRequestFrame struct {
	MetaInfo MetaInfo `json:"metainfo"`
	Data     struct {
		ID int `json:"id"`
	} `json:"data"`
}

type AdminFlagFrame struct {
	MetaInfo MetaInfo `json:"metainfo"`
	Data     struct {
		IsAdmin bool `json:"is_admin"`
	} `json:"data"`
}
