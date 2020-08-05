package fs

import (
	"os"
	"strings"
	"testing"
)

func TestGetFilePathsRecursively(t *testing.T) {
	numFiles := 4
	files := GetFilePathsRecursively("../../test/data")
	if len(files) != numFiles {
		t.Errorf("Should find only %d files", numFiles)
	}
}

func TestGetModTime(t *testing.T) {
	mtime := GetModTime("../../test/data/101APPLE/slack-logo.png")
	date := strings.Split(mtime, " ")[0]
	if date != "2020-04-17" {
		t.Errorf("2020-04-17 should equal to %s", date)
	}
}

func TestCreateFolderIfNotExist(t *testing.T) {
	var ret bool
	path := "../../test/testdir"
	os.RemoveAll(path)
	ret = CreateFolderIfNotExist(path)
	if ret != true {
		t.Error("Cannot create subfolder")
	}
	ret = CreateFolderIfNotExist(path)
	if ret != false {
		t.Error("Should be OK when folder already exists")
	}
	os.RemoveAll(path)
}

func TestCopyFile(t *testing.T) {
	fromPath := "../../test/testfile"
	toPath := "../../test/testfile2"
	os.RemoveAll(fromPath)
	os.RemoveAll(toPath)
	CreateFile(fromPath)
	CopyFile(fromPath, toPath)
	if !(CheckIfPathExists(fromPath) && CheckIfPathExists(toPath)) {
		t.Error("Fail to copy file")
	}
	os.RemoveAll(fromPath)
	os.RemoveAll(toPath)
}

func TestMoveFile(t *testing.T) {
	fromPath := "../../test/testfile"
	toPath := "../../test/testfile2"
	os.RemoveAll(fromPath)
	os.RemoveAll(toPath)
	CreateFile(fromPath)
	MoveFile(fromPath, toPath)
	if !(!CheckIfPathExists(fromPath) && CheckIfPathExists(toPath)) {
		t.Error("Fail to move file")
	}
	os.RemoveAll(fromPath)
	os.RemoveAll(toPath)
}
