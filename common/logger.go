package common

import (
	"runtime"

	log "github.com/sirupsen/logrus"
)

type APILogger struct {
	Stdout *log.Logger
	Stderr *log.Logger
}

func NewAPILogger(stdout *log.Logger, stderr *log.Logger) *APILogger {
	return &APILogger{
		Stdout: stdout,
		Stderr: stderr,
	}
}

func (a *APILogger) HandleErrorWithTrace(err error) (b bool) {

	if err != nil {
		// notice that we're using 1, so it will actually log the where
		// the error happened, 0 = this function, we don't want that.
		pc, fn, line, _ := runtime.Caller(1)
		a.Stderr.Errorf("[error] in %s[%s:%d] %v", runtime.FuncForPC(pc).Name(), fn, line, err)
		b = true
	}
	return
}
