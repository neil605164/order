package structer

import "time"

// EnvConfig dev.yaml格式
type EnvConfig struct {
	DBMaster DBMaster `yaml:"master"`
	DBSlave  DBSlave  `yaml:"slave"`
	DB       DB       `yaml:"db"`
	Redis    Redis    `yaml:"redis"`
}

// DBMaster 載入db的master環境設定
type DBMaster struct {
	Host     string `yaml:"host"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

// DBSlave 載入db的slave環境設定
type DBSlave struct {
	Host     string `yaml:"host"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

// DB 對DB其他操作的設定
type DB struct {
	Debug bool `yaml:"debug"`
}

// Redis 載入redis設定
type Redis struct {
	RedisHost string `yaml:"redisHost"`
	RedisPort string `yaml:"redisPort"`
	RedisPwd  string `yaml:"redisPassword"`
}

// APIResult 回傳API格式
type APIResult struct {
	Result interface{} `json:"result"`
	Status RespStatus  `json:"status"`
}

type RespStatus struct {
	ErrorCode   int       `json:"errorCode" example:"1000"`
	ErrorMsg    string    `json:"errorMsg" example:"error message"`
	Datetime    time.Time `json:"datetime" example:"2022-07-21T12:19:39-04:00"`
	LogIDentity string    `json:"logID" example:"qw13er65tyui74rg22o"`
}
