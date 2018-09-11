package log

var logger Logger = &StandardLogger{}

func Debug(msg string, data ...interface{}) {
	logger.Debug(msg, data)
}
func Info(msg string, data ...interface{}) {
	logger.Info(msg, data)
}
func Warn(msg string, data ...interface{}) {
	logger.Warn(msg, data)
}
func Error(msg string, data ...interface{}) {
	logger.Error(msg, data)
}
func Panic(msg string, data ...interface{}) {
	logger.Panic(msg, data)
}
func Fatal(msg string, data ...interface{}) {
	logger.Fatal(msg, data)
}

func UseLog(lg Logger) {
	logger = lg
}

func SetLevel(lvl Level) {
	logger.SetLevel(int8(lvl))
}