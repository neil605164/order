package helper

import (
	"order/app/global"
	"order/app/global/errorcode"

	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
	"syscall"

	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
)

// CheckFileIsExist 檢查檔案 + 路徑是否存在
func CheckFileIsExist(filePath, fileName string, perm os.FileMode) error {
	// 重新設置umask
	syscall.Umask(0)

	// 檢查檔案路徑是否存在
	if _, err := os.Stat(filePath + fileName); os.IsNotExist(err) {
		// 建制資料夾
		if err = os.MkdirAll(filePath, perm); err != nil {
			log.Printf("❌ WriteLog: 建立資料夾錯誤 [%v] ❌ \n", err.Error())
			return err
		}

		//  建制檔案
		_, err = os.Create(filePath + fileName)
		if err != nil {
			log.Printf("❌ WriteLog: 建立檔案錯誤 [%v] ❌ \n", err.Error())
			return err
		}
	}

	return nil
}

// ValidateStruct 驗證struct規則
func ValidateStruct(req interface{}) error {
	validate := validator.New()
	err := validate.Struct(req)
	if err != nil {
		return err
	}
	return nil
}

// ValidateRegex 驗證字串正則規則
func ValidateRegex(req, regex string) bool {
	ok, _ := regexp.MatchString(regex, req)
	return ok
}

// InArray 檢測值是否在陣列內
func InArray(val string, array []string) (exists bool) {
	for _, v := range array {
		if val == v {
			return true
		}
	}
	return false
}

// CatchError 回傳不可預期的錯誤
func CatchError(c *gin.Context) {
	if err := recover(); err != nil {
		// 回傳不可預期的錯誤
		apiErr := ErrorHandle(global.FatalLog, errorcode.Code.UnExpectedError, err)
		c.JSON(http.StatusBadRequest, Fail(apiErr))
	}
}

// IndexOf 搜尋值在陣列的索引位置
func IndexOf(element string, data []string) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1 // not found.
}

// RemoveIndex 刪除陣列元素
func RemoveIndex(array []string, element string) []string {
	index := IndexOf(element, array)

	return append(array[:index], array[index+1:]...)
}

// IsDeveloperEnv 檢查是否為開發環境(local, develop)
func IsDeveloperEnv() bool {

	env := os.Getenv("ENV")

	if isDev := strings.Contains(env, "develop"); isDev {
		return true
	}

	if isLocal := strings.Contains(env, "local"); isLocal {
		return true
	}

	return false
}

// IsRelease 檢查是否是正式環境
func IsRelease() bool {
	return os.Getenv("ENV") == "release"
}

// IsStage 檢查是否是測試環境
func IsStage() bool {
	return os.Getenv("ENV") == "stage"
}

// IsStress 檢查是否是壓測環境
func IsStress() bool {
	return os.Getenv("ENV") == "stress"
}

// IsLocal 檢查是否是正式環境
func IsLocal() bool {
	return os.Getenv("ENV") == "local" || os.Getenv("ENV") == "local-docker"
}
