package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
)

func init() { register("type", typeCmd) }

func typeCmd(args []string, stdout io.Writer) {
	for _, name := range args {
		if _, ok := builtins[name]; ok {
			fmt.Fprintf(stdout, "%s is a shell builtin\n", name)
		} else if path, err := exec.LookPath(name); err == nil {
			fmt.Fprintf(stdout, "%s is %s\n", name, path)
		} else {
			fmt.Fprintf(os.Stderr, "%s: not found\n", name)
		}
	}
}
