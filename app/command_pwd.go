package main

import (
	"fmt"
	"os"
	"io"
)

func init() { register("pwd", pwdCmd) }

func pwdCmd(args []string, stdout io.Writer) {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Fprintln(os.Stderr, "pwd: could not get current working directory")
		return
	}
	fmt.Fprintln(stdout, dir)
}
