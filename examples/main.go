package main

import (
	"github.com/gobasis/log"
	"time"
	"github.com/gobasis/log/zapimpl"
)

/*
Description:
show demos of gobasis/log
 * Author: architect.bian
 * Date: 2018/09/10 18:30
*/
func main() {
	benchMark()
	//levelDemo()
}

var url = "www.chain.com"

func benchMark() {
	log.UseLog(&log.StandardLogger{})
	fromStd := time.Now().UnixNano()
	for i := 0; i < 1000000; i++ {
		log.Info("failed to fetch URL", "url", url, "attempt", 3, "backoff", time.Second)
	}
	toStd := time.Now().UnixNano()
	log.UseLog(&zapimpl.Logger{})
	fromZap := time.Now().UnixNano()
	for i := 0; i < 1000000; i++ {
		log.Info("failed to fetch URL", "url", url, "attempt", 3, "backoff", time.Second)
	}
	toZap := time.Now().UnixNano()
	log.Info("elapse time", "fromStd", fromStd, "toStd", toStd, "diff", toStd-fromStd)
	log.Info("elapse time", "fromZap", fromZap, "toZap", toZap, "diff", toZap-fromZap)
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
