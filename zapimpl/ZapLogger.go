package zapimpl

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"github.com/gobasis/log"
	"path/filepath"
	"os"
	"io/ioutil"
	"encoding/json"
)

type Logger struct {
	log       *zap.SugaredLogger
	atomLevel zap.AtomicLevel
}

func(z *Logger) Initialize(conf *log.Config) {
	var opts = []zap.Option{zap.AddCallerSkip(2)}
	z.atomLevel = zap.NewAtomicLevel()
	workDir, _ := filepath.Abs("./")
	var file = filepath.Join(workDir, "conf", "zap.json")
	_, err := os.Stat(file)
	var log *zap.Logger
	var cfg = zap.NewProductionConfig()
	if err != nil {
		cfg.Level = z.atomLevel
	} else { // load config from file then create a logger
		buf, err := ioutil.ReadFile(file)
		if err != nil {
			panic(err)
		}
		if err := json.Unmarshal(buf, &cfg); err != nil {
			panic(err)
		}
		cfg.Level = z.atomLevel
	}
	if conf != nil { //set config from log.json
		z.SetLevel(int8(conf.Level))
		cfg.OutputPaths = conf.OutputPaths
		cfg.ErrorOutputPaths = conf.ErrorOutputPaths
	}
	//conf.Sampling = nil //disable sampling
	log, _ = cfg.Build(opts...)
	z.log = log.Sugar()
}

func (z *Logger) getLog() *zap.SugaredLogger {
	if z.log == nil {
		z.Initialize(nil)
	}
	return z.log
}

func (z *Logger) SetLevel(lvl int8) {
	if lvl > int8(log.ErrorLevel) {
		z.atomLevel.SetLevel(zapcore.Level(lvl))
	} else {
		z.atomLevel.SetLevel(zapcore.Level(lvl - 1))
	}
}

func (z *Logger) Debug(msg string, data []interface{}) {
	z.getLog().Debugw(msg, data...)
}

func (z *Logger) Info(msg string, data []interface{}) {
	z.getLog().Infow(msg, data...)
}

func (z *Logger) Warn(msg string, data []interface{}) {
	z.getLog().Warnw(msg, data...)
}

func (z *Logger) Error(msg string, data []interface{}) {
	z.getLog().Errorw(msg, data...)
}

func (z *Logger) Panic(msg string, data []interface{}) {
	z.getLog().Panicw(msg, data...)
}

func (z *Logger) Fatal(msg string, data []interface{}) {
	z.getLog().Fatalw(msg, data...)
}
