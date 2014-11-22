// package toylog provides simple leveled logging.
package toylog

import (
	golog "log"
  "os"
)

type logLevel int

const (
	LevelDebug logLevel = iota
	LevelInfo
	LevelError
	LevelFatal
)

var levels = map[logLevel]string{
	LevelDebug: "D",
	LevelInfo:  "I",
	LevelError: "E",
	LevelFatal: "F",
}

func levelPrefix(level logLevel, space bool) string {
	l := levels[level] + ":"
	if space {
		l = l + " "
	}
	return l
}

func rawPrintLn(level logLevel, msgs ...interface{}) {
	m := append([]interface{}{levelPrefix(level, false)}, msgs...)
	golog.Println(m...)
}

func rawPrintF(level logLevel, msg string, args ...interface{}) {
	msg = levelPrefix(level, true) + msg
	golog.Printf(msg, args...)
}

type logger struct {
	level logLevel
}

func (l *logger) Set(level logLevel) {
	l.level = level
}

func (l *logger) Debugln(v ...interface{}) {
	if l.level == LevelDebug {
		rawPrintLn(LevelDebug, v...)
	}
}

func (l *logger) Debugf(format string, v ...interface{}) {
	if l.level == LevelDebug {
		rawPrintF(LevelDebug, format, v...)
	}
}

func (l *logger) Infoln(v ...interface{}) {
	rawPrintLn(LevelInfo, v...)
}

func (l *logger) Infof(format string, v ...interface{}) {
	rawPrintF(LevelInfo, format, v...)
}

func (l *logger) Errorln(v ...interface{}) {
	rawPrintLn(LevelError, v...)
}

func (l *logger) Errorf(format string, v ...interface{}) {
	rawPrintF(LevelError, format, v...)
}

func (l *logger) Fatalln(v ...interface{}) {
	m := append([]interface{}{levelPrefix(LevelFatal, false)}, v...)
	golog.Fatalln(m...)
}

func (l *logger) Fatalf(format string, v ...interface{}) {
	format = levelPrefix(LevelFatal, true) + format
	golog.Fatalf(format, v...)
}


func Set(level logLevel) {
	defaultLog.level = level
}

func Verbose() {
	Set(LevelDebug)
}

func Debugf(format string, v ...interface{}) {
	defaultLog.Debugf(format, v...)
}

func Debugln(v ...interface{}) {
	defaultLog.Debugln(v...)
}

func Infof(format string, v ...interface{}) {
	defaultLog.Infof(format, v...)
}

func Infoln(v ...interface{}) {
	defaultLog.Infoln(v...)
}

func Errorf(format string, v ...interface{}) {
	defaultLog.Errorf(format, v...)
}

func Errorln(v ...interface{}) {
	defaultLog.Errorln(v...)
}

func Fatalf(format string, v ...interface{}) {
	defaultLog.Fatalf(format, v...)
}

func Fatalln(v ...interface{}) {
	defaultLog.Fatalln(v...)
}

func Stdout() {
  golog.SetOutput(os.Stdout)
}

var defaultLog = logger{level: LevelInfo}

func init() {
	golog.SetFlags(0)
}
