package fs

import (
	"io"
	"os"
	"path/filepath"
	"syscall"
	"time"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

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

func CheckIfPathExists(path string) bool {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

/*
 * Return true when there is a folder created
 */
func CreateFolderIfNotExist(path string) bool {
	if !CheckIfPathExists(path) {
		dirErr := os.MkdirAll(path, 0755)
		if dirErr != nil {
			panic("Cannot create output subfolder")
		}
		return true
	}

	return false
}

func CreateFile(path string) {
	f, err := os.Create(path)
	check(err)
	defer f.Close()
}

func CopyFile(fromPath, toPath string) {
	from, err := os.Open(fromPath)
	check(err)
	defer from.Close()

	to, err := os.OpenFile(toPath, os.O_RDWR|os.O_CREATE, 0666)
	check(err)
	defer to.Close()

	_, err = io.Copy(to, from)
	check(err)
}

func MoveFile(fromPath, toPath string) {
	err := os.Rename(fromPath, toPath)
	check(err)
}
