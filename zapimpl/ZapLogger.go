package zapimpl

import (
	"encoding/json"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io/ioutil"
	"os"
	"path/filepath"
	"github.com/gobasis/log"
)

type Logger struct {
	log       *zap.SugaredLogger
	atomLevel zap.AtomicLevel
}

func (z *Logger) getLog() *zap.SugaredLogger {
	if z.log == nil {
		var opts = []zap.Option{zap.AddCallerSkip(2)}
		z.atomLevel = zap.NewAtomicLevel()
		workDir, _ := filepath.Abs("./")
		var file = filepath.Join(workDir, "conf", "zap.json")
		_, err := os.Stat(file)
		var log *zap.Logger
		if err != nil {
			conf := zap.NewProductionConfig()
			conf.Level = z.atomLevel
			//conf.Sampling = nil //disable sampling
			log, _ = conf.Build(opts...)
		} else { // load config from file then create a logger
			buf, err := ioutil.ReadFile(file)
			if err != nil {
				panic(err)
			}
			var conf = zap.NewProductionConfig()
			if err := json.Unmarshal(buf, &conf); err != nil {
				panic(err)
			}
			conf.Level = z.atomLevel
			log, _ = conf.Build(opts...)
		}
		z.log = log.Sugar()
	}
	return z.log
}

func (z *Logger) SetLevel(lvl int8) {
	if z.log == nil {
		z.getLog()
	}
	if lvl > int8(log.ErrorLevel) {
		z.atomLevel.SetLevel(zapcore.Level(lvl + 1))
	} else {
		z.atomLevel.SetLevel(zapcore.Level(lvl))
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
