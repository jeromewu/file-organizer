package main

import (
	"fmt"
	"github.com/jeromewu/file-organizer/internal/utils"
	"github.com/jeromewu/file-organizer/pkg/fs"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	handleFile := fs.CopyFile
	newPaths := make([]string, 0)
	op := "copy"
	iPath, oPath, isMove := utils.ParseArgs(os.Args[0], os.Args[1:])
	if isMove {
		op = "move"
		handleFile = fs.MoveFile
	}
	paths := fs.GetFilePathsRecursively(iPath)
	for idx, path := range paths {
		newPath := strings.Trim(oPath, "/") + "/" + fs.GetModTime(path)[:7] + "/" + filepath.Base(path)
		newPaths = append(newPaths, newPath)
		fmt.Printf("[%d] %s -> %s\n", idx+1, path, newPath)
	}
	fmt.Printf("\n%d files to %s.\n", len(paths), op)

	yes := utils.AskForConfirmation("Are you sure to execute?")

	if yes {
		for idx, path := range paths {
			newPath := newPaths[idx]
			fs.CreateFolderIfNotExist(filepath.Dir(newPath))
			fmt.Printf("%s %s to %s\n", op, path, newPath)
			handleFile(path, newPath)
		}
	} else {
		fmt.Println("Cancelled")
	}
}
