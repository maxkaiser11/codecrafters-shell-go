package main

import (
	"errors"
	"io"
	"os"
)

// parseRedirects pulls "> file" / "1> file" out of the tokens.
// It returns the remaining args and the target filename ("" if none).
func parseRedirects(tokens []string) (args []string, outFile string, err error) {
	for i := 0; i < len(tokens); i++ {
		switch tokens[i] {
		case ">", "1>":
			if i+1 >= len(tokens) {
				return nil, "", errors.New("syntax error near unexpected token `newline'")
			}
			outFile = tokens[i+1]
			i++ // skip the filename
		default:
			args = append(args, tokens[i])
		}
	}
	return args, outFile, nil
}

// openOutput returns where stdout should go, plus a function to clean up afterwards.
func openOutput(outFile string) (io.Writer, func(), error) {
	if outFile == "" {
		return os.Stdout, func() {}, nil
	}
	f, err := os.Create(outFile)
	if err != nil {
		return nil, nil, err
	}
	return f, func() { f.Close() }, nil
}
