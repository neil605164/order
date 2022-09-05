package helper

import (
	"crypto/md5"
	"fmt"
	"io"
	"math/rand"
	"order/app/global"
	"order/app/global/errorcode"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
)

type Log struct {
	HasError    bool        `json:"has_error"`
	LogIDentity string      `json:"log_id"`
	Params      interface{} `json:"params"`
	Result      string      `json:"result"`
	Path        string      `json:"path"`
	FuncName    string      `json:"func_name"`
	Skip        int         `json:"skip"`
}

var errorLog Log

// 紀錄 log 的檔案位置和行數
func getFilePath(l *Log) {
	pc, file, line, ok := runtime.Caller(errorLog.Skip)
	if !ok {
		panic("Could not get context info for logger!")
	}

	l.Path = file + ":" + strconv.Itoa(line)
	funcname := runtime.FuncForPC(pc).Name()
	fmt.Println(funcname)
	fmt.Println(strings.LastIndex(funcname, "/"))
	l.FuncName = funcname[strings.LastIndex(funcname, "/")+1:]
}

// ErrorHandle 取錯誤代碼 + 寫錯誤 Log
func ErrorHandle(errorType string, errorCode errorcode.NewErrorCode, errMsg interface{}, param ...interface{}) (apiErr errorcode.Error) {
	var logID string

	// New 一個 Error Interface
	apiErr = errorcode.NewError()

	// 塞入 Error 對應清單
	apiErr.SetErrorCode(errorCode)

	message := fmt.Sprintf("%v: %v", errorCode.ErrorCode, errorCode.ErrorMsg)
	if errMsg != "" {
		message = fmt.Sprintf("%v:%v:%v", errorCode.ErrorCode, errorCode.ErrorMsg, errMsg)
	}

	switch errorType {
	case global.SuccessLog:
		logID = success(message, param)
	case global.WarnLog:
		logID = warn(message, param)
	case global.FatalLog:
		logID = fatal(message, param)
	default:
		logID = fatal(message, param)
	}

	// 存入 Log 識別證
	apiErr.SetLogID(logID)

	return
}

// 設定 log 紀錄的格式 (json or text)
func SetFormatter(f *logrus.JSONFormatter) {
	logrus.SetFormatter(f)
}

// 設定追朔 log 的層級
func SetReportCallerSkip(n int) {
	errorLog.Skip = n
}

// 是否有紀錄 log 錯誤
func HasError() bool {
	return errorLog.HasError
}

// http access log
func Access(c *gin.Context) {
	fileName := "http-access"
	if os.Getenv("LOG_NAME") != "" {
		fileName = os.Getenv("LOG_NAME") + "-access"
	}

	filePath := fmt.Sprintf("%v%v%v", global.LogPath, fileName, global.FileSuffix) + ".%Y%m%d"

	/* 日誌輪轉相關函式
	`WithLinkName` 為最新的日誌建立軟連線
	`WithRotationTime` 設定日誌分割的時間，隔多久分割一次
	`WithMaxAge 和 WithRotationCount二者只能設定一個
	`WithMaxAge` 設定檔案清理前的最長儲存時間
	`WithRotationCount` 設定檔案清理前最多儲存的個數
	*/

	writer, _ := rotatelogs.New(
		filePath,
		rotatelogs.WithRotationCount(30), // 保留 30 份
		rotatelogs.WithRotationTime(time.Hour*24), // 24 小時切割一次
	)

	// 同時寫 log file 跟輸出到 stdout
	mw := io.MultiWriter(os.Stdout, writer)
	logrus.SetOutput(mw)

	// 	寫access Log
	logrus.WithFields(logrus.Fields{
		"LogType":     "[💚 START 💚 ]",
		"ClientIP":    c.ClientIP(),
		"Path":        c.Request.URL.Path,
		"Status":      c.Writer.Status(),
		"Method":      c.Request.Method,
		"Params":      []string{},
		"HTTPReferer": c.GetHeader("Referer"),
	}).Info()
}

// warn log
func warn(result string, params ...interface{}) string {
	fileName := "http-warn"
	if os.Getenv("LOG_NAME") != "" {
		fileName = os.Getenv("LOG_NAME") + "-warn"
	}

	filePath := fmt.Sprintf("%v%v%v", global.LogPath, fileName, global.FileSuffix) + ".%Y%m%d"
	/* 日誌輪轉相關函式
	`WithLinkName` 為最新的日誌建立軟連線
	`WithRotationTime` 設定日誌分割的時間，隔多久分割一次
	`WithMaxAge 和 WithRotationCount二者只能設定一個
	`WithMaxAge` 設定檔案清理前的最長儲存時間
	`WithRotationCount` 設定檔案清理前最多儲存的個數
	*/

	writer, _ := rotatelogs.New(
		filePath,
		rotatelogs.WithRotationCount(30), // 保留 30 份
		rotatelogs.WithRotationTime(time.Hour*24), // 24 小時切割一次
	)

	// 同時寫 log file 跟輸出到 stdout
	mw := io.MultiWriter(os.Stdout, writer)
	logrus.SetOutput(mw)

	errorLog.HasError = true
	// LogIDentity log 唯一碼目前不需要
	errorLog.LogIDentity = md5EncryptionWithTime(RandString(6))
	errorLog.Result = result
	errorLog.Params = params
	getFilePath(&errorLog)

	// write Log
	logrus.WithFields(logrus.Fields{
		"LogType":     "[⚠️ Warn ⚠️ ]",
		"path":        errorLog.Path,
		"funcname":    errorLog.FuncName,
		"logIDentity": errorLog.LogIDentity,
		"params":      errorLog.Params,
	}).Warn(errorLog.Result)

	return errorLog.LogIDentity
}

// fatal log
func fatal(result string, params ...interface{}) string {

	fileName := "http-fatal"
	if os.Getenv("LOG_NAME") != "" {
		fileName = os.Getenv("LOG_NAME") + "-fatal"
	}

	filePath := fmt.Sprintf("%v%v%v", global.LogPath, fileName, global.FileSuffix) + ".%Y%m%d"

	/* 日誌輪轉相關函式
	`WithLinkName` 為最新的日誌建立軟連線
	`WithRotationTime` 設定日誌分割的時間，隔多久分割一次
	`WithMaxAge 和 WithRotationCount二者只能設定一個
	`WithMaxAge` 設定檔案清理前的最長儲存時間
	`WithRotationCount` 設定檔案清理前最多儲存的個數
	*/

	writer, _ := rotatelogs.New(
		filePath,
		rotatelogs.WithRotationCount(30), // 保留 30 份
		rotatelogs.WithRotationTime(time.Hour*24), // 24 小時切割一次
	)

	// 同時寫 log file 跟輸出到 stdout
	mw := io.MultiWriter(os.Stdout, writer)
	logrus.SetOutput(mw)

	errorLog.HasError = true
	// LogIDentity log 唯一碼目前不需要
	errorLog.LogIDentity = md5EncryptionWithTime(RandString(6))
	errorLog.Result = result
	errorLog.Params = params
	getFilePath(&errorLog)

	// write Log
	logrus.WithFields(logrus.Fields{
		"LogType":     "[❌ Fatal❌ ]",
		"path":        errorLog.Path,
		"funcname":    errorLog.FuncName,
		"logIDentity": errorLog.LogIDentity,
		"params":      errorLog.Params,
	}).Warn(errorLog.Result)

	return errorLog.LogIDentity
}

// success log
func success(result string, params ...interface{}) string {
	fileName := "http-success"
	if os.Getenv("LOG_NAME") != "" {
		fileName = os.Getenv("LOG_NAME") + "-success"
	}

	filePath := fmt.Sprintf("%v%v%v", global.LogPath, fileName, global.FileSuffix) + ".%Y%m%d"

	/* 日誌輪轉相關函式
	`WithLinkName` 為最新的日誌建立軟連線
	`WithRotationTime` 設定日誌分割的時間，隔多久分割一次
	`WithMaxAge 和 WithRotationCount二者只能設定一個
	`WithMaxAge` 設定檔案清理前的最長儲存時間
	`WithRotationCount` 設定檔案清理前最多儲存的個數
	*/

	writer, _ := rotatelogs.New(
		filePath,
		rotatelogs.WithRotationCount(30), // 保留 30 份
		rotatelogs.WithRotationTime(time.Hour*24), // 24 小時切割一次
	)

	// 同時寫 log file 跟輸出到 stdout
	mw := io.MultiWriter(os.Stdout, writer)
	logrus.SetOutput(mw)

	errorLog.HasError = true
	// LogIDentity log 唯一碼目前不需要
	errorLog.LogIDentity = md5EncryptionWithTime(RandString(6))
	errorLog.Result = result
	errorLog.Params = params
	getFilePath(&errorLog)

	// write Log
	logrus.WithFields(logrus.Fields{
		"LogType":     "[✔️ SUCCESS ✔️ ]",
		"path":        errorLog.Path,
		"funcname":    errorLog.FuncName,
		"logIDentity": errorLog.LogIDentity,
		"params":      errorLog.Params,
	}).Info(errorLog.Result)

	return errorLog.LogIDentity
}

// md5EncryptionWithTime md5 加密（加上奈秒時間）
func md5EncryptionWithTime(str string) string {
	naTime := time.Now().UnixNano()
	data := str + strconv.FormatInt(naTime, 10)
	key := []byte(data)

	token := md5.Sum(key)
	md5Str := fmt.Sprintf("%x", token)

	return md5Str
}

// RandString 生成指定長度的字符串
func RandString(length int) string {
	return stringWithCharset(length, charset)
}

func stringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))
