package preset

import "log"

type logger interface {
	Logf(format string, args ...interface{})
}
type debuglogger struct{}
type noplogger struct{}

var _dlog logger = &debuglogger{}
var _noplog logger = &noplogger{}
var _log = _noplog

func (t *noplogger) Logf(format string, args ...interface{}) {}
func (t *debuglogger) Logf(format string, args ...interface{}) {
	log.Printf(format, args...)
}

func GetDebugLogger() logger {
	return _dlog
}
func GetNopLogger() logger {
	return _noplog
}

func GetLogger() logger {
	return _log
}

func SetLogger(l logger) {
	_log = l
}
