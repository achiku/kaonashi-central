package main

import "testing"

func TestConfig(t *testing.T) {
	testConfigFilePath := "./configs/test.toml"
	config, err := NewAppConfig(testConfigFilePath)
	if err != nil {
		t.Fatalf("failed to create config: %s", err)
	}
	if config.Database.UserName != "pgtest" {
		t.Fatalf("expected pgtest, but got %s", config.Database.UserName)
	}
}
