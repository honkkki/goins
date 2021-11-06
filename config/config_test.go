package config

import "testing"

func TestGetSavePath(t *testing.T) {
	path := GetSavePath()
	t.Log("save path:", path)
	t.Log("root dir:", rootDir)
}