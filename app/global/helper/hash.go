package helper

import (
	"crypto/md5"
	"fmt"
	"order/app/global"
	"order/app/global/errorcode"
	"strconv"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// Md5Encryption md5加密
func Md5Encryption(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	md5Str := fmt.Sprintf("%x", has)

	return md5Str
}

// Md5EncryptionWithTime md5 加密（加上奈秒時間）
func Md5EncryptionWithTime(str string) string {
	naTime := time.Now().UnixNano()
	data := str + strconv.FormatInt(naTime, 10)
	key := []byte(data)

	token := md5.Sum(key)
	md5Str := fmt.Sprintf("%x", token)

	return md5Str
}

// HashPassword 密碼加密(註冊管理者使用)
func HashPassword(password string) (value string, apiErr errorcode.Error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	if err != nil {
		apiErr = ErrorHandle(global.WarnLog, errorcode.Code.CryptionError, err.Error())
		return string(bytes), apiErr
	}

	return string(bytes), apiErr
}

// CheckPasswordHash 檢查檢查(登入使用))
func CheckPasswordHash(password, dbPwd string) (result bool) {
	err := bcrypt.CompareHashAndPassword([]byte(dbPwd), []byte(password))
	if err == nil {
		return true
	}
	return
}
