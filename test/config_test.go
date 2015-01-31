package config

import (
	_ "github.com/joho/godotenv/autoload"

	"os"
	"testing"
)

func TestDotEnv(t *testing.T) {
	var v string
	v = os.Getenv("TEST")
	if v != "FOO" {
		t.Error("Expected FOO, got ", v)
	}
}
