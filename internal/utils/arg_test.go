package utils

import "testing"

func TestParseArgs(t *testing.T) {
	iPath, oPath, isMove := ParseArgs("", []string{"-i", "input"})
	if iPath != "input" {
		t.Errorf("Parsed input folder path is not `input`")
	}
	if oPath != "out" {
		t.Errorf("Parsed output folder path is not `out`")
	}
	if isMove != false {
		t.Errorf("Mode is not `copy`")
	}
	iPath, oPath, isMove = ParseArgs("", []string{"-i", "input", "-o", "output", "-m"})
	if oPath != "output" {
		t.Errorf("Parsed output folder path is not `output`")
	}
	if isMove != true {
		t.Errorf("Mode is not `move`")
	}
}
