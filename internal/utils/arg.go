package utils

import (
	"bytes"
	"flag"
)

func ParseArgs(pname string, args []string) (string, string, bool) {
	var buf bytes.Buffer
	var iPath, oPath string
	var move bool
	flags := flag.NewFlagSet(pname, flag.ContinueOnError)
	flags.StringVar(&iPath, "i", "", "Input folder path")
	flags.StringVar(&oPath, "o", "out", "Output folder path")
	flags.BoolVar(&move, "m", false, "Whether to move files, false means to copy only")
	flags.SetOutput(&buf)

	err := flags.Parse(args)
	if err != nil {
		panic(buf.String())
	}

	if len(iPath) == 0 {
		panic("You MUST assign input folder path")
	}

	return iPath, oPath, move
}
