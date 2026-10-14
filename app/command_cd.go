package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func init() { register("cd", cdCmd) }

func cdCmd(args []string, stdout io.Writer) {
	arg := "~" // bare `cd` goes home, like bash
	if len(args) > 0 {
		arg = args[0]
	}

	target := arg
	if arg == "~" || strings.HasPrefix(arg, "~/") {
		home, err := os.UserHomeDir()
		if err != nil {
			fmt.Fprintln(os.Stderr, "cd: could not get home directory")
			return
		}
		target = home + arg[1:]
	}

	if err := os.Chdir(target); err != nil {
		fmt.Fprintf(os.Stderr, "cd: %s: No such file or directory\n", arg)
	}
}
