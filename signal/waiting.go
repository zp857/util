package signal

import (
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
)

func WaitingSignal() {
	sig := make(chan os.Signal)
	// SIGHUP: terminal closed
	// SIGINT: Ctrl+C
	// SIGTERM: program exit
	// SIGQUIT: Ctrl+/
	// SIGKILL: kill
	signal.Notify(sig,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
		syscall.SIGKILL,
	)
	for t := range sig {
		zap.L().Sugar().Errorf("收到信号: %v, 程序退出\n\n", t.String())
		os.Exit(0)
	}
}
