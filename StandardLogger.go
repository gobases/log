package log

import (
	"log"
	"os"
)

type StandardLogger struct {
	log *log.Logger
	level Level
}

func(s *StandardLogger) Initialize(conf *Config) {
	s.log = log.New(os.Stderr, "", log.LstdFlags)
	if conf != nil {
		s.level = conf.Level
		//s.log.SetOutput(&rotate.Logger{
		//	Filename:   "foo.log",
		//	MaxSize:    1, // megabytes
		//	MaxBackups: 3,
		//	MaxAge:     28, // days
		//	Compress:   true, // disabled by default
		//})
	}
}

func(s *StandardLogger) getLog() *log.Logger {
	if s.log == nil {
		initialize(nil)
	}
	return s.log
}

func(s *StandardLogger) SetLevel(lvl int8) {
	s.level = Level(lvl)
}

func(s *StandardLogger) Debug(msg string, data []interface{}) {
	if s.level <= DebugLevel {
		s.getLog().SetPrefix("DEBUG ")
		s.getLog().Printf(msg, data...)
	}
}

func(s *StandardLogger) Info(msg string, data []interface{}) {
	if s.level <= InfoLevel {
		s.getLog().SetPrefix("INFO  ")
		s.getLog().Printf(msg, data...)
	}
}

func(s *StandardLogger) Warn(msg string, data []interface{}) {
	if s.level <= WarnLevel {
		s.getLog().SetPrefix("WARN  ")
		s.getLog().Printf(msg, data...)
	}
}

func(s *StandardLogger) Error(msg string, data []interface{}) {
	if s.level <= ErrorLevel {
		s.getLog().SetPrefix("ERROR ")
		s.getLog().Printf(msg, data...)
	}
}

func(s *StandardLogger) Panic(msg string, data []interface{}) {
	if s.level <= PanicLevel {
		s.getLog().SetPrefix("PANIC ")
		s.getLog().Printf(msg, data...)
		panic(msg)
	}
}

func(s *StandardLogger) Fatal(msg string, data []interface{}) {
	if s.level <= FatalLevel {
		s.getLog().SetPrefix("FATAL ")
		s.getLog().Printf(msg, data...)
		os.Exit(1)
	}
}