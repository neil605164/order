package errorcode

// NewErrorCode 錯誤代碼格式
type NewErrorCode struct {
	ErrorCode int    `json:"error_code"`
	ErrorMsg  string `json:"error_msg"`
}

type Errorcode struct {
	Success                NewErrorCode
	UnExpectedError        NewErrorCode
	PermissionDenied       NewErrorCode
	DBConnectError         NewErrorCode
	DBTableNotExist        NewErrorCode
	ServiceIsNotExist      NewErrorCode
	CronJobError           NewErrorCode
	CronJobStart           NewErrorCode
	CronJobStop            NewErrorCode
	CronJobPrepareStop     NewErrorCode
	CronJobFuncNotExist    NewErrorCode
	CronJobSuccessExecute  NewErrorCode
	HTTPServerStart        NewErrorCode
	TCPPortDuplicate       NewErrorCode
	PrePareShutDownService NewErrorCode
	ServiceAlreadyShutdown NewErrorCode
	JSONMarshalError       NewErrorCode
	JSONUnMarshalError     NewErrorCode
	GetTimeZoneError       NewErrorCode
	ParseTimeError         NewErrorCode
	CryptionError          NewErrorCode
	BindParamError         NewErrorCode
	ValidateParamError     NewErrorCode
	GetUserListError       NewErrorCode
	CreateUserFail         NewErrorCode
	UpdateUserFail         NewErrorCode
}

var Code = Errorcode{
	Success:                NewErrorCode{0, "Success"},                      // 呼叫API成功
	UnExpectedError:        NewErrorCode{400, "UnExpected Error"},           // 不預期的錯誤
	PermissionDenied:       NewErrorCode{403, "Permission Denied"},          // 權限不足
	ServiceIsNotExist:      NewErrorCode{1000, "Service Is Not Exist"},      // 服務不存在
	DBConnectError:         NewErrorCode{1001, "DB Connect Error"},          // DB 連線錯誤
	DBTableNotExist:        NewErrorCode{1002, "DB Table Not Exist"},        // DB Table 不存在
	CronJobError:           NewErrorCode{1003, "Cron Job Error"},            // 背景錯誤
	CronJobStart:           NewErrorCode{1004, "Cron Job Start"},            // 背景啟動
	CronJobStop:            NewErrorCode{1005, "Cron Job Stop"},             // 背景停止
	CronJobPrepareStop:     NewErrorCode{1006, "Cron Job Prepare Stop"},     // 背景準備停止
	CronJobFuncNotExist:    NewErrorCode{1007, "Cron Job Func Not Exist"},   // 背景 func 尚未註冊
	CronJobSuccessExecute:  NewErrorCode{1008, "Cron Job Success Execute"},  // 背景成功執行
	HTTPServerStart:        NewErrorCode{1009, "Http Server Start"},         // http server 服務啟動
	TCPPortDuplicate:       NewErrorCode{1010, "Tcp Port Duplicate"},        // tcp port 重複
	PrePareShutDownService: NewErrorCode{1011, "PrePareShutDownService"},    // 準備關閉服務
	ServiceAlreadyShutdown: NewErrorCode{1012, "Service Already Shut down"}, // 服務已關閉
	JSONMarshalError:       NewErrorCode{1013, "Json Marshal Error"},        // Json Marshal Error
	JSONUnMarshalError:     NewErrorCode{1014, "Json Un Marshal Error"},     // Json UnMarshal Error
	GetTimeZoneError:       NewErrorCode{1015, "Get Time Zone Error"},       // 取時區錯誤
	ParseTimeError:         NewErrorCode{1016, "Parse Time Error"},          // 時間轉換錯誤
	CryptionError:          NewErrorCode{1017, "Cryption Error"},            // 加密錯誤
	BindParamError:         NewErrorCode{1018, "Bind Param Error"},          // 取參數錯誤
	ValidateParamError:     NewErrorCode{1019, "ValidateParamError"},        // 驗證參數錯誤

	// 2000 會員錯誤代碼
	GetUserListError: NewErrorCode{2000, "Get User List Error"}, // 取會員清單錯誤
	CreateUserFail:   NewErrorCode{2001, "Create User Fail"},    // 建立會員失敗
	UpdateUserFail:   NewErrorCode{2002, "Update User Fail"},    // 更新會員失敗
}
