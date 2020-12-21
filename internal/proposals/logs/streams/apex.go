package main

import (
	"github.com/apex/log"
)

func main() {

	//log.Interface()

	//log.SetLevel(log.ErrorLevel)
	var dict map[string]string
	log.WithField("len", len(dict)).WithField("dict is nil", dict == nil).Info("maps in go")
	testApex()
}



func testApex() {
	log.Infof("just normal log with msg: %s", "Hello From Zap")
	log.Errorf("Error msg: %s", "some error occured")
}