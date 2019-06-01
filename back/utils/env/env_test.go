package env_test

import (
	"github.com/MasterYourLife/back/utils/env"
	"os"
	"testing"
)

func TestGetEnv(t *testing.T) {
	var key, value = "TEST_GET_ENV", "plop"
	if env.GetEnv(key, "notset") != "notset" {
		t.Errorf("undefined var should fallback to defaultValue")
	}
	os.Setenv(key, value)
	if env.GetEnv(key, "") != value {
		t.Errorf("failed to GetEnv")
	}
}
