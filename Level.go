package log

type Level int8

const (
	// DebugLevel logs are typically voluminous, and are usually disabled in production.
	DebugLevel Level = iota
	// InfoLevel is the default logging priority.
	InfoLevel
	// WarnLevel logs are more important than Info, but don't need individual
	// human review.
	WarnLevel
	// ErrorLevel logs are high-priority. If an application is running smoothly,
	// it shouldn't generate any error-level logs.
	ErrorLevel
	// DevDebugLevel logs are same as DebugLevel, but panic/fatal will print error like nomal
	DevDebugLevel
	// PanicLevel logs a message, then panics.
	PanicLevel
	// FatalLevel logs a message, then calls os.Exit(1).
	FatalLevel
)

func LevelParse(lvl string) Level {
	switch lvl {
	case "debug":
		return DebugLevel
	case "info":
		return InfoLevel
	case "warn":
		return WarnLevel
	case "dev":
		return DevDebugLevel
	case "panic":
		return PanicLevel
	case "error":
		return ErrorLevel
	case "fatal":
		return FatalLevel
	}
	return DebugLevel
}