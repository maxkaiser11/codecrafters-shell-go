package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("$ ")
		line, err := reader.ReadString('\n')
		if err != nil {
			os.Exit(0)
		}

		tokens := parseArgs(strings.TrimSpace(line))
		if len(tokens) == 0 {
			continue
		}

		run(tokens)
	}
}

// run applies any redirection, then dispatches to a builtin or an external program.
func run(tokens []string) {
	args, outFile, err := parseRedirects(tokens)
	if err != nil {
		fmt.Fprintln(os.Stderr, "shell:", err)
		return
	}
	if len(args) == 0 {
		return
	}

	stdout, closeOut, err := openOutput(outFile)
	if err != nil {
		fmt.Fprintln(os.Stderr, "shell:", err)
		return
	}
	defer closeOut()

	name, rest := args[0], args[1:]
	if fn, ok := builtins[name]; ok {
		fn(rest, stdout)
		return
	}
	runExternal(name, rest, stdout)
}
