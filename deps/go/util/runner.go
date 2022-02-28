package util

import (
	"time"

	"github.com/codingeasygo/util/debug"
	"sxbastudio.com/base/go/pgx"
	"sxbastudio.com/base/go/xlog"
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
