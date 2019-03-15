package envconfig

import (
	"os"
	"testing"
)

func TestEnvConfig(t *testing.T) {
	homeFromEnv := os.Getenv("HOME")
	homeFromCfg, err := Get("HOME")
	if err != nil {
		t.Errorf("expected %s to be defined", "HOME")
	}
	if homeFromCfg != homeFromEnv {
		t.Errorf("expected HOME to be %s but got %s", homeFromEnv, homeFromCfg)
	}
}
