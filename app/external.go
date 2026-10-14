package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
)

// runExternal looks the program up on PATH and runs it, sending its output to stdout.
func runExternal(name string, args []string, stdout io.Writer) {
	if _, err := exec.LookPath(name); err != nil {
		fmt.Fprintf(os.Stderr, "%s: command not found\n", name)
		return
	}

	cmd := exec.Command(name, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = stdout
	cmd.Stderr = os.Stderr
	_ = cmd.Run() // non-zero exit codes are the program's business, not the shell's
}
