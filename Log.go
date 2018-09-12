package log

import (
	"path/filepath"
	"os"
	"io/ioutil"
	"encoding/json"
)

type Config struct {
	Use    	       string `json:"use" yaml:"use"`
	level          Level
	LevelName      string `json:"level" yaml:"level"`
	//Rotate RotateConf `json:"rotate" yaml:"rotate"`
}
var conf Config
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
	SetLevel(conf.level)
}

func SetLevel(lvl Level) {
	conf.level = lvl
	logger.SetLevel(int8(lvl))
}

func init() {
	workDir, _ := filepath.Abs("./")
	var file = filepath.Join(workDir, "conf", "log.json")
	_, err := os.Stat(file)
	if err == nil {
		buf, err := ioutil.ReadFile(file)
		if err != nil {
			panic(err)
		}
		if err := json.Unmarshal(buf, &conf); err != nil {
			panic(err)
		}
		if len(conf.LevelName) != 0 {
			SetLevel(LevelParse(conf.LevelName))
		}

	}
}