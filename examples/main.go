package main

import (
	"github.com/gobasis/log"
	"time"
	"github.com/gobasis/log/zapimpl"
	"go.uber.org/zap"
	"math/rand"
	"strconv"
)

/*
Description:
show demos of gobasis/log
 * Author: architect.bian
 * Date: 2018/09/10 18:30
*/
func main() {
	benchMark()
	levelDemo()
}

var url = "www.chain.com"

func benchMark() {
	var count = 1000000
	log.UseLog(&log.StandardLogger{})
	fromStd := time.Now().Unix()
	for i := 0; i < count; i++ {
		log.Info("using standard log, failed to fetch URL" + strconv.Itoa(rand.Int()), "url", url, "attempt", i, "backoff", time.Duration(rand.Int63()))
	}
	toStd := time.Now().Unix()
	log.UseLog(&zapimpl.Logger{})
	fromZap := time.Now().Unix()
	for i := 0; i < count; i++ {
		log.Info("using zap log, failed to fetch URL" + strconv.Itoa(rand.Int()), "url", url, "attempt", i, "backoff", time.Duration(rand.Int63()))
	}
	toZap := time.Now().Unix()
	zlog, _ := zap.NewProduction()
	sugar := zlog.Sugar()
	fromRawZap := time.Now().Unix()
	for i := 0; i < count; i++ {
		sugar.Infow("using raw zap log, failed to fetch URL" + strconv.Itoa(rand.Int()),
			// Structured context as loosely typed key-value pairs.
			"url", url,
			"attempt", i,
			"backoff", time.Duration(rand.Int63()),
		)
	}
	toRawZap := time.Now().Unix()
	log.Info("elapse time", "fromStd", fromStd, "toStd", toStd, "diff", toStd-fromStd)
	log.Info("elapse time", "fromZap", fromZap, "toZap", toZap, "diff", toZap-fromZap)
	log.Info("elapse time", "fromRawZap", fromRawZap, "toRawZap", toRawZap, "diff", toRawZap-fromRawZap)
}

func levelDemo() {
	log.UseLog(&zapimpl.Logger{}) // use zap log
	log.SetLevel(log.InfoLevel)
	log.Debug("failed to fetch URL", "url", url, "attempt", 3, "backoff", time.Second)
	log.Info("failed to fetch URL", "url", url, "attempt", 3, "backoff", time.Second)
	log.Warn("failed to fetch URL", "url", url, "attempt", 3, "backoff", time.Second)
	log.Error("failed to fetch URL", "url", url, "attempt", 3, "backoff", time.Second)
	//log.Panic("failed to fetch URL", "url", url, "attempt", 3, "backoff", time.Second)
	//log.Fatal("failed to fetch URL", "url", url, "attempt", 3, "backoff", time.Second)
}
