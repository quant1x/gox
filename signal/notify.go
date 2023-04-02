package signal

import (
	"os"
	"os/signal"
)

// Notify 指定默认监控信号
func Notify() chan os.Signal {
	//创建监听退出chan
	sigs := make(chan os.Signal)
	//监听指定信号 ctrl+c kill
	signal.Notify(sigs, stopSignals...)

	return sigs
}
