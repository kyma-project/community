package main

import (
	"github.com/apex/log"
	"github.com/apex/log/handlers/json"
	"os"
)

func main() {
	log.SetHandler(json.New(os.Stderr))
	log.ErrorLevel
	testApexLogger(log.WithField("request_id", "dajdhaskj"))
}

func testApexLogger(logEntry *log.Entry) {
	logEntry.Infof("just normal log with msg: %s", "Hello From Zap")
	logEntry.Errorf("Error msg: %s", "some error occured")
}
