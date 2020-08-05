package utils

import (
	"strings"
	"testing"
)

func TestUtils_GetFilePathsRecursively(t *testing.T) {
	numFiles := 4
	files := GetFilePathsRecursively("../../test/data")
	if len(files) != numFiles {
		t.Errorf("Should find only %d files", numFiles)
	}
}

func TestUtils_GetModTime(t *testing.T) {
	mtime := GetModTime("../../test/data/101APPLE/slack-logo.png")
	date := strings.Split(mtime, " ")[0]
	if date != "2020-04-17" {
		t.Errorf("2020-04-17 should equal to %s", date)
	}
}
