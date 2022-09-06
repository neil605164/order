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
	IllegalAmount          NewErrorCode
	GetProductByIdError    NewErrorCode
	ProductNotExist        NewErrorCode
	CreateOrderFail        NewErrorCode
	RedisSubscribeStart    NewErrorCode
	RedisSubscribeFail     NewErrorCode
	QueueStop              NewErrorCode
	GetOrderByIdError      NewErrorCode
	CanNotDeleteOrder      NewErrorCode
	DeleteOrderFail        NewErrorCode
	GetOrderListError      NewErrorCode
	UpdateOrderFail        NewErrorCode
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
	RedisSubscribeStart:    NewErrorCode{1010, "Redis Subscribe Start"},     // Redis Queue Start
	RedisSubscribeFail:     NewErrorCode{1011, "Redis Subscribe Fail"},      // Redis Queue Fail
	QueueStop:              NewErrorCode{1012, "Queue Stop"},                // Queue Stop

	// 2000 產品錯誤代碼
	GetProductByIdError: NewErrorCode{2000, "Get Product By Id Error"}, // 透過產品 id 取產品資料錯誤
	ProductNotExist:     NewErrorCode{2001, "Product Not Exist"},       // 產品不存在

	// 3000 訂單錯誤代碼
	IllegalAmount:     NewErrorCode{3000, "Illegal Amount"},        // 不合法的金額
	CreateOrderFail:   NewErrorCode{3001, "Create Order Fail"},     // 建立訂單錯誤
	GetOrderByIdError: NewErrorCode{3002, "Get Order By Id Error"}, // 透過 id 取訂單資料錯誤
	CanNotDeleteOrder: NewErrorCode{3003, "Can Not Delete Order"},  // 禁止取消訂單
	DeleteOrderFail:   NewErrorCode{3004, "Delete Order Fail"},     // 刪除訂單失敗
	GetOrderListError: NewErrorCode{3005, "Get Order List Error"},  // 取訂單清單錯誤
	UpdateOrderFail:   NewErrorCode{3006, "Update Order Fail"},     // 更新訂單錯誤

}
