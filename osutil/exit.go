package osutil

import (
	"os"
	"os/signal"
	"syscall"
	"sync"
)

var showdownFns [] func(os.Signal)
var l = &sync.Mutex{}

func WaitSignal(sigs ...os.Signal) {
	if len(sigs) == 0 {
	}

	stopSignals := make(chan os.Signal, 1)
	if len(sigs) == 0 {
		signal.Notify(stopSignals, syscall.SIGINT, syscall.SIGTERM)
	} else {
		signal.Notify(stopSignals, sigs...)
	}

	sig := <-stopSignals
	l.Lock()
	defer l.Unlock()
	for i := len(showdownFns) - 1; i >= 0; i-- {
		showdownFns[i](sig)
	}
}

func RegisterShutDown(fn func(os.Signal)) {
	l.Lock()
	showdownFns = append(showdownFns, fn)
	l.Unlock()
}