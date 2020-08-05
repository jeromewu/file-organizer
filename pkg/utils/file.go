package utils

import (
	"os"
	"path/filepath"
	"syscall"
	"time"
)

func timespecToTime(ts syscall.Timespec) time.Time {
	return time.Unix(int64(ts.Sec), int64(ts.Nsec))
}

func GetFilePathsRecursively(rootPath string) []string {
	files := make([]string, 0)
	e := filepath.Walk(rootPath, func(path string, f os.FileInfo, err error) error {
		if !f.IsDir() {
			files = append(files, path)
		}
		return err
	})

	if e != nil {
		panic(e)
	}

	return files
}

func GetModTime(path string) string {
	finfo, _ := os.Stat(path)
	stat_t := finfo.Sys().(*syscall.Stat_t)
	return timespecToTime(stat_t.Mtim).String()
}
