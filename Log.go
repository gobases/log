package log

import (
	"path/filepath"
	"os"
	"io/ioutil"
	"encoding/json"
)

type Config struct {
	Use    	       string `json:"use" yaml:"use"`
	Level          Level
	LevelName      string `json:"level" yaml:"level"`
	OutputPaths    []string `json:"outputPaths" yaml:"outputPaths"`
	ErrorOutputPaths []string `json:"errorOutputPaths" yaml:"errorOutputPaths"`
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
	initialize(&conf)
}

func initialize(conf *Config) {
	if logger != nil {
		if conf != nil {
			conf.Level = LevelParse(conf.LevelName)
		}
		logger.Initialize(conf)
	}
}

func SetLevel(lvl Level) {
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
		initialize(&conf)
	}
}