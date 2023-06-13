package logger

import "github.com/sirupsen/logrus"

func Tracef(format string, args ...interface{}) {
	Log.Logf(logrus.TraceLevel, format, args...)
}

func Debugf(format string, args ...interface{}) {
	Log.Logf(logrus.DebugLevel, format, args...)
}

func Infof(format string, args ...interface{}) {
	Log.Logf(logrus.InfoLevel, format, args...)
}

func Warnf(format string, args ...interface{}) {
	Log.Logf(logrus.WarnLevel, format, args...)
}

func Warningf(format string, args ...interface{}) {
	Log.Warnf(format, args...)
}

func Errorf(format string, args ...interface{}) {
	Log.Logf(logrus.ErrorLevel, format, args...)
}

func Fatalf(format string, args ...interface{}) {
	Log.Logf(logrus.FatalLevel, format, args...)
	Log.Exit(1)
}

func Panicf(format string, args ...interface{}) {
	Log.Logf(logrus.PanicLevel, format, args...)
}

func Trace(args ...interface{}) {
	Log.Log(logrus.TraceLevel, args...)
}

func Debug(args ...interface{}) {
	Log.Log(logrus.DebugLevel, args...)
}

func Info(args ...interface{}) {
	Log.Log(logrus.InfoLevel, args...)
}

func Warn(args ...interface{}) {
	Log.Log(logrus.WarnLevel, args...)
}

func Warning(args ...interface{}) {
	Log.Warn(args...)
}

func Error(args ...interface{}) {
	Log.Log(logrus.ErrorLevel, args...)
}

func Fatal(args ...interface{}) {
	Log.Log(logrus.FatalLevel, args...)
	Log.Exit(1)
}

func Panic(args ...interface{}) {
	Log.Log(logrus.PanicLevel, args...)
}

func TraceFn(fn logrus.LogFunction) {
	Log.LogFn(logrus.TraceLevel, fn)
}

func DebugFn(fn logrus.LogFunction) {
	Log.LogFn(logrus.DebugLevel, fn)
}

func InfoFn(fn logrus.LogFunction) {
	Log.LogFn(logrus.InfoLevel, fn)
}

func WarnFn(fn logrus.LogFunction) {
	Log.LogFn(logrus.WarnLevel, fn)
}

func WarningFn(fn logrus.LogFunction) {
	Log.WarnFn(fn)
}

func ErrorFn(fn logrus.LogFunction) {
	Log.LogFn(logrus.ErrorLevel, fn)
}

func FatalFn(fn logrus.LogFunction) {
	Log.LogFn(logrus.FatalLevel, fn)
	Log.Exit(1)
}

func PanicFn(fn logrus.LogFunction) {
	Log.LogFn(logrus.PanicLevel, fn)
}

func Traceln(args ...interface{}) {
	Log.Logln(logrus.TraceLevel, args...)
}

func Debugln(args ...interface{}) {
	Log.Logln(logrus.DebugLevel, args...)
}

func Infoln(args ...interface{}) {
	Log.Logln(logrus.InfoLevel, args...)
}

func Warnln(args ...interface{}) {
	Log.Logln(logrus.WarnLevel, args...)
}

func Warningln(args ...interface{}) {
	Log.Warnln(args...)
}

func Errorln(args ...interface{}) {
	Log.Logln(logrus.ErrorLevel, args...)
}

func Fatalln(args ...interface{}) {
	Log.Logln(logrus.FatalLevel, args...)
	Log.Exit(1)
}

func Panicln(args ...interface{}) {
	Log.Logln(logrus.PanicLevel, args...)
}
