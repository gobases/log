package log

type Logger interface {
	Initialize(conf *Config)
	SetLevel(lvl int8)
	Debug(msg string, data []interface{})
	Info(msg string, data []interface{})
	Warn(msg string, data []interface{})
	Error(msg string, data []interface{})
	Panic(msg string, data []interface{})
	Fatal(msg string, data []interface{})
}
