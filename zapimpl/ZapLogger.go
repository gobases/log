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

func (z *Logger) getLog() *zap.Logger {
	if z.log == nil {
		z.atomLevel = zap.NewAtomicLevel()
		workDir, _ := filepath.Abs("./")
		var file = filepath.Join(workDir, "conf", "log.json")
		_, err := os.Stat(file)
		if err != nil || os.IsNotExist(err) {
			encoderConf := zap.NewProductionEncoderConfig()
			z.log = zap.New(zapcore.NewCore(
				zapcore.NewJSONEncoder(encoderConf),
				zapcore.Lock(os.Stdout),
				z.atomLevel,
			))
		} else { // load config from file then create a logger
			buf, err := ioutil.ReadFile(file)
			if err != nil {
				panic(err)
			}
			var cfg = zap.NewProductionConfig()
			if err := json.Unmarshal(buf, &cfg); err != nil {
				panic(err)
			}
			cfg.Level = z.atomLevel
			logger, err := cfg.Build()
			if err != nil {
				panic(err)
			}
			z.log = logger
		}
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
	z.getLog().Sugar().Debugw(msg, data...)
}

func (z *Logger) Info(msg string, data []interface{}) {
	z.getLog().Sugar().Infow(msg, data...)
}

func (z *Logger) Warn(msg string, data []interface{}) {
	z.getLog().Sugar().Warnw(msg, data...)
}

func (z *Logger) Error(msg string, data []interface{}) {
	z.getLog().Sugar().Errorw(msg, data...)
}

func (z *Logger) Panic(msg string, data []interface{}) {
	z.getLog().Sugar().Panicw(msg, data...)
}

func (z *Logger) Fatal(msg string, data []interface{}) {
	z.getLog().Sugar().Fatalw(msg, data...)
}
