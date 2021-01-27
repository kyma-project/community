package main

import (
	"github.com/apex/log"
	"github.com/apex/log/handlers/json"
	"os"
)

func main() {
	log.SetHandler(json.New(os.Stderr))
	testApexLogger(log.WithField("request_id", "random_string"))
}

func testApexLogger(logEntry *log.Entry) {
	logEntry.WithField("context", "a").Infof("just a normal log entry with the msg: %s", "Hello From Apex")
	logEntry.Errorf("Error msg: %s", "an error occurred")
}
