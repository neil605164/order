package schedule

import (
	"fmt"
	"order/app/global"
	"order/app/global/errorcode"
	"order/app/global/helper"
	"order/internal/bootstrap"
	"os"

	"github.com/robfig/cron/v3"
)

// Run 執行背景服務
func Run() {
	// 載入排程
	cronIns := SeriesIns()
	jobs := cronIns.LoadSchedule()

	bg := cron.New(cron.WithSeconds())

	// 塞入排程
	for _, job := range jobs {
		job.Init()
		pid, err := bg.AddJob(job.Spec, cron.NewChain(cron.SkipIfStillRunning(cron.DefaultLogger), cron.Recover(cron.DefaultLogger)).Then(job))
		if err != nil {
			_ = helper.ErrorHandle(global.WarnLog, errorcode.Code.CronJobError, err)
		}

		// 設定 pid
		job.SetEntryID(pid)
	}

	// 開始排程
	_ = helper.ErrorHandle(global.SuccessLog, errorcode.Code.CronJobStart, "🔔 crontanb success start 🔔")
	bg.Start()

	// 等待結束訊號
	<-bootstrap.GracefulDown()
	_ = helper.ErrorHandle(global.WarnLog, errorcode.Code.CronJobPrepareStop, "🚦  排程收到訊號囉，等待其他背景完成，準備結束排程 🚦")

	// 停止排程
	bg.Stop()

	// hook
	select {
	// 關閉背景
	case <-bootstrap.WaitOnceSignal():
		fmt.Println("🚦  收到關閉訊號，強制結束 🚦")

		// 等待背景結束
		for _, job := range jobs {
			fmt.Println(job)
			job.Wait()
		}

		_ = helper.ErrorHandle(global.SuccessLog, errorcode.Code.CronJobStop, "🚦  收到關閉訊號，強制結束 🚦")
		os.Exit(2)
	}

	fmt.Println("🚦  已經結束 🚦")
}
