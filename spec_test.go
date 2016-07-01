package spec

import (
	"testing"
	"time"
)

func TestSpec(t *testing.T) {
	config := AllFeatures()
	config.URL = "tcp://localhost:1883"

	// mosquitto specific config
	config.Authentication = false
	config.MessageRetainWait = 500 * time.Millisecond

	Run(t, config)
}
