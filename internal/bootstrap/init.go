package bootstrap

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
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
