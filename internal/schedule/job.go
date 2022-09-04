package schedule

import (
	"fmt"
	"order/app/global"
	"order/app/global/errorcode"
	"order/app/global/helper"
	"sync"
	"time"

	"github.com/robfig/cron/v3"
)

type function_name func() (apiErr errorcode.Error)
type CronJob struct {
	Name     string          `json:"name"`      // 背景名稱
	Spec     string          `json:"spec"`      // 執行週期
	FuncName function_name   `json:"func_name"` // 函式名稱
	EntryID  cron.EntryID    `json:"entry_id"`  // EntryID
	isRetry  bool            // 重複執行
	mux      *sync.RWMutex   // 讀寫鎖
	wg       *sync.WaitGroup // 等待通道
}

var Singleton *CronJob
var Once sync.Once

// SeriesIns 獲得單例對象
func SeriesIns() *CronJob {
	Once.Do(func() {
		Singleton = &CronJob{}
	})
	return Singleton
}

// Run 自定義 Crob Job 接口
func (c *CronJob) Run() {
	// 加鎖，檢查是否可以執行背景
	c.mux.RLock()
	isRetry := c.isRetry
	// 解鎖
	c.mux.RUnlock()

	// 如果不可重複則跳過
	if !isRetry {
		return
	}

	// todo 背景開關功能，撈 db 檢查

	// 開始前，基本設定
	c.wg.Add(1)

	// 開始執行
	startTime := time.Now()
	apiErr := c.Exec()
	// 執行後，基本設定
	endtime := time.Now()

	c.wg.Done()

	// 紀錄執行時間
	c.RecordJobStatus(c, startTime, endtime, apiErr)
}

// Init 初始化
func (c *CronJob) Init() {
	c.wg = new(sync.WaitGroup)
	c.mux = new(sync.RWMutex)
}

// Wait 等待 wg 結束
func (c *CronJob) Wait() {
	c.wg.Wait()
}

// Exec 開始執行背景
func (c *CronJob) Exec() (apiErr errorcode.Error) {
	if c.FuncName == nil {
		_ = helper.ErrorHandle(global.WarnLog, errorcode.Code.CronJobFuncNotExist, c.FuncName)
	}

	return c.FuncName()
}

// SetEntryID 設定 pid
func (c *CronJob) SetEntryID(entryID cron.EntryID) {
	c.EntryID = entryID
}

// RecordJobStatus 背景執行完畢，會呼叫這個Func來紀錄執行狀態
func (c *CronJob) RecordJobStatus(job *CronJob, startTime, endTime time.Time, apiErr errorcode.Error) {

	msg := ""
	execTime := endTime.Sub(startTime)
	if apiErr != nil {
		msg = fmt.Sprintf("%v error, error reason %v , and totally spent %v", c.Name, apiErr.Error(), execTime)
		_ = helper.ErrorHandle(global.WarnLog, errorcode.Code.CronJobError, msg)
		return
	}

	msg = fmt.Sprintf("%v execute success, and totally spent %v", c.Name, execTime)
	_ = helper.ErrorHandle(global.SuccessLog, errorcode.Code.CronJobSuccessExecute, msg)
}
