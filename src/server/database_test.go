package main

import "testing"

func TestDatabasePing(t *testing.T) {
	testConfigFilePath := "./configs/test.toml"
	config, err := NewAppConfig(testConfigFilePath)
	if err != nil {
		t.Fatalf("failed to create config: %s", err)
	}
	db, err := NewDB(config)
	if err != nil {
		t.Fatalf("failed to get db: %s", err)
	}
	err = db.conn.Ping()
	if err != nil {
		t.Fatalf("failed to ping: %s", err)
	}
}
