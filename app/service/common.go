package service

// request Teamplus Request 項目
type request struct {
	URL    string            `json:"url"`
	Header map[string]string `json:"header"`
	Param  interface{}       `json:"param"`
}

// content Teamplus 訊息內容
type content struct {
	Req  request     `json:"req"`
	Resp interface{} `json:"resp"`
}
