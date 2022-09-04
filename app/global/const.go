package global

// DBMaster db master
const DBMaster = "db_master"

// DBSlaver db slave
const DBSlaver = "db_slave"

/** 正則規則 **/
const (
	AdminAccount  = `[a-zA-Z0-9_]{2,30}` // AdminAccount 管理者帳號
	AdminPassword = `[a-zA-Z0-9]{6,20}`  // AdminPassword 管理者密碼
)

/** Redis Cache Time **/
const (
	OrderDetailExpire      = 3600        // 細單快取時間 1小時
	RoundIDTokenExpire     = 3600        // RoundID Token 過期時間 1小時
	RedisCyExpire          = 300         // CY API資料 5分鐘
	RedisDBExpire          = 600         // 資料庫資料 10分鐘
	RedisLoginExpire       = 864000      // 管理者帳號登入 10天
	RedisSendMessageExpire = 1 * 60 * 60 // Token 1小時
)

/** 平台設定 **/
const (
	FrontEnd = "frontend" // 前台
	BackEnd  = "backend"  // 後台
)

/** 類別清單 **/
const (
	Marquee = 1 // 跑馬燈
	News    = 2 // 最新報導
)

/** 錯誤 Log 類型 **/
const (
	WarnLog    = "Warn"    // 警告型 Log
	FatalLog   = "Fatal"   // 嚴重型 Log
	SuccessLog = "Success" // 成功型 Log
)

/** HTTP CURL 設定 **/
const (
	TimeOut = 10 // 連api 10 秒timeout
)

/** 檔案權限管理 **/
const (
	DirPermission  = 0o755 // 資料夾權限
	FilePermission = 0o644 // 檔案權限
)

/** Redis Key **/
const (
	TeamPlus = "Redis:TeamPlus"
)

/** Log 相關用 **/
const (
	LogPath    = "./storage/logs/" // Log 存放位置
	FileSuffix = ".log"            // log 副檔名
)

/** 錯誤訊息 **/
const (
	RecordNotFound = "record not found"
	RedisNotFound  = "redigo: nil returned"
)

/** 審查狀態 **/
const (
	NotVerify  = "Not Verify"  // 未驗證
	Verifying  = "Verifying"   // 驗證中
	Verified   = "Verified"    // 驗證通過
	VerifyFail = "Verify Fail" // 驗證失敗
)

const LetterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	LetterIdxBits = 6                    // 6 bits to represent a letter index
	LetterIdxMask = 1<<LetterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	LetterIdxMax  = 63 / LetterIdxBits   // # of letter indices fitting in 63 bits
)
