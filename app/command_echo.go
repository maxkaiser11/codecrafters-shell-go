package main

import (
	"fmt"
	"io"
	"strings"
)

func init() { register("echo", echoCmd) }

func echoCmd(args []string, stdout io.Writer) {
	fmt.Fprintln(stdout, strings.Join(args, " "))
}
