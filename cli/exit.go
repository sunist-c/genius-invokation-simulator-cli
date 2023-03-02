package cli

import (
	"os"
	"time"
)

func Exit(code int) {
	// 退出前给logger留一定时间写入日志
	time.Sleep(time.Millisecond * 100)
	os.Exit(code)
}
