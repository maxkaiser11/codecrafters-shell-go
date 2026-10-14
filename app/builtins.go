package main

import "io"

// builtinFunc is the signature every builtin command implements.
// Normal output goes to stdout (which may be a file); errors go to os.Stderr.
type builtinFunc func(args []string, stdout io.Writer)

// builtins is the registry of shell builtins, filled by each builtin_*.go file's init().
var builtins = map[string]builtinFunc{}

func register(name string, fn builtinFunc) {
	builtins[name] = fn
}
