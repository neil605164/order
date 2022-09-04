package swagger

// errorResp 錯誤回傳
type errorResp struct {
	ErrorCode int         `json:"error_code" example:"120103"`
	ErrorMsg  string      `json:"error_msg" example:"SERVICE NOT REGISTER(120103)"`
	LogID     string      `json:"log_id" example:"43bce937fc5d8bac1ae266f29f3217ac"`
	Result    interface{} `json:"result"`
}
