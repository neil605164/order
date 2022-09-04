package schedule

// LoadSchedule 載入所有排程
func (c *CronJob) LoadSchedule() (jobs []*CronJob) {

	// 載入所有排程
	jobs = []*CronJob{
		// 範例
		// {
		// 	Name:     "印出 hello world", // 排程名稱
		// 	Spec:     "@every 10s",     // 排程時間
		// 	FuncName: c.HelloWorld,  // 對應的 func 名稱
		// 	isRetry:  true,             // 是否可重複執行
		// },
	}

	return
}
