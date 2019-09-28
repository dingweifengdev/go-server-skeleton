package ext

import (
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/sirupsen/logrus"
)

var (
	hooks  []func()
	mu     sync.Mutex
	chQuit chan struct{}
)

func init() {
	chQuit = make(chan struct{})
	chSig := make(chan os.Signal, 1)

	// sigterm signal sent from kubernetes
	signal.Notify(chSig, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	logrus.Debug("Init shutdown hook")

	go func() {
		<-chSig

		logrus.Info("Start shutdown...")
		st := time.Now()

		for i := len(hooks) - 1; i >= 0; i-- {
			hooks[i]()
		}

		logrus.Infof("Finish shutdown in %v",
			time.Since(st).Truncate(time.Millisecond))
		close(chQuit)
	}()
}

func OnShutdown(hook func()) {
	mu.Lock()
	hooks = append(hooks, hook)
	mu.Unlock()
}

func WaitShutdown() {
	<-chQuit
}