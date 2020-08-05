package main

import (
	"fmt"
	"github.com/jeromewu/file-organizer/pkg/utils/file"
)

func main() {
	fmt.Println("file organizer")
	file.GetFilePathsRecursively("../test/data")
}
