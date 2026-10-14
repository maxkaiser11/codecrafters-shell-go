package main

import (
	"io"
	"os"
	"strconv"
)

func init() { register("exit", exitCmd) }

func exitCmd(args []string, stdout io.Writer) {
	code := 0
	if len(args) > 0 {
		if n, err := strconv.Atoi(args[0]); err == nil {
			code = n
		}
	}
	os.Exit(code)
}
