package main

import (
	"testing"
	"time"
)

func TestTimeFormat(t *testing.T) {
	//Example of log produced by apex
	timestamp := "2020-12-22T12:52:39.906885+01:00"

	_, err := time.Parse(time.RFC3339, timestamp)

	if err != nil {
		t.Error(err)
	}
}
