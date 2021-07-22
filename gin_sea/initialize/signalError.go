package initialize

import (
	"gin/global"
	"os"
	"os/signal"
)


func init() {
	global.ServerSignChan = make(chan os.Signal)
}

func ShutDownServer() {
	//接收该Interrupt指令
	global.ServerSignChan <- os.Interrupt
}

func ServerNotify() {
	//阻塞监听是否有Interrupt，如果有Interrupt指令进入，停止阻塞，执行下面代码
	signal.Notify(global.ServerSignChan, os.Interrupt)
	<-global.ServerSignChan
}
