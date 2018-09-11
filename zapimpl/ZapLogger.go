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
	log       *zap.Logger
	atomLevel zap.AtomicLevel
}

func (z *Logger) getLog() *zap.SugaredLogger {
	if z.log == nil {
		var opts = []zap.Option{zap.AddCallerSkip(2)}
		z.atomLevel = zap.NewAtomicLevel()
		workDir, _ := filepath.Abs("./")
		var file = filepath.Join(workDir, "conf", "log.json")
		_, err := os.Stat(file)
		if err != nil || os.IsNotExist(err) {
			conf := zap.NewProductionConfig()
			conf.Level = z.atomLevel
			z.log, _ = conf.Build(opts...)
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
			z.log, _ = conf.Build(opts...)
		}
	}
	return z.log.Sugar()
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
