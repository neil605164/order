package bootstrap

import (
	"fmt"
	"log"
	"order/app/global"
	"order/app/global/errorcode"
	"order/app/global/helper"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var sig chan os.Signal
var serverClose chan struct{}

// SetupGracefulSignal 設定優雅關閉的信號
func SetupGracefulSignal() {
	sig = make(chan os.Signal, 1)
	serverClose = make(chan struct{})
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sig
		fmt.Println("🔥收到關閉 channel 通知🔥")
		close(serverClose)
	}()
}

// GracefulDown 優雅結束程式
func GracefulDown() <-chan struct{} {
	return serverClose
}

// WaitOnceSignal 等待一次的訊號
func WaitOnceSignal() (sig chan os.Signal) {
	sig = make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	return
}

// Queue 優雅停止
func SetupQueueGracefulSignal() {
	sig = make(chan os.Signal, 1)
	serverClose = make(chan struct{})
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sig
		fmt.Println("🔥收到關閉 channel 通知🔥")
		for i := 10; i > 0; i-- {
			log.Printf("%v 秒後停止服務", i)
			time.Sleep(1 * time.Second)
		}

		fmt.Println("🚦  等待時間到強制結束 🚦")

		_ = helper.ErrorHandle(global.SuccessLog, errorcode.Code.QueueStop, "🚦  收到關閉訊號，強制結束 🚦")
		os.Exit(2)
	}()
}
