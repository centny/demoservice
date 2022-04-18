package util

import (
	"time"

	"github.com/centny/demoservice/deps/go/pgx"
	"github.com/centny/demoservice/deps/go/xlog"
	"github.com/codingeasygo/util/debug"
)

//NamedRunner will run call by delay
func NamedRunner(name string, delay time.Duration, running *bool, call func() error) {
	xlog.Infof("%v is starting", name)
	var finishCount = 0
	runCall := func() error {
		defer func() {
			if perr := recover(); perr != nil {
				xlog.Errorf("%v is panic with %v, callstaick is \n%v", perr, debug.CallStatck())
			}
		}()
		return call()
	}
	for *running {
		err := runCall()
		if err == nil {
			finishCount++
			continue
		}
		if err != pgx.ErrNoRows {
			xlog.Warnf("%v is fail with %v", name, err)
		} else if finishCount > 0 {
			xlog.Debugf("%v is having %v finished", name, finishCount)
		}
		finishCount = 0
		time.Sleep(delay)
	}
	xlog.Warnf("%v is stopped", name)
}
