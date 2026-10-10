package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

var builtins = map[string]bool{
	"echo": true,
	"exit": true,
	"type": true,
	"pwd":  true,
	"cd":   true,
}

func main() {
	// TODO: Uncomment the code below to pass the first stage

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("$ ")
		line, err := reader.ReadString('\n')
		if err != nil {
			os.Exit(0)
		}
		line = strings.TrimSpace(line)

		tokens := parseArgs(line)
		if len(tokens) == 0 {
			return
		}

		cmd, args := tokens[0], tokens[1:]

		switch cmd {
		case "exit":
			os.Exit(0)
		case "echo":
			if len(args) > 0 &&
				strings.HasPrefix(args[0], "'") &&
				strings.HasSuffix(args[len(args)-1], "'") &&
				!(len(args) == 1 && args[0] == "'") {
				// quoted
				fmt.Println(strings.Join(args, " "))
			}
			fmt.Println(strings.Join(args, " "))
		case "type":
			if len(args) == 0 {
				continue
			}
			if builtins[args[0]] {
				fmt.Println(args[0] + " is a shell builtin")
			} else {
				checkPATH(args[0])
			}
		case "pwd":
			printWorkingDirectory()
		case "cd":
			changeDirectory(args[0])
		default:
			if _, err := exec.LookPath(cmd); err != nil {
				fmt.Printf("%s: command not found\n", cmd)
				continue
			}
			if err := execCommand(cmd, args...); err != nil {
				_ = err
			}
		}
	}
}

func checkType(command string) bool {
	return builtins[command]
}

func checkPATH(name string) {
	path, err := exec.LookPath(name)
	if err != nil {
		fmt.Printf("%s: not found\n", name)
		return
	}
	fmt.Printf("%s is %s\n", name, path)

}

func checkIfExecutable(name string) bool {
	path, err := exec.LookPath(name)
	if err != nil {
		return false
	}
	fmt.Printf("%s\n", path)
	return true
}

func execCommand(command string, args ...string) error {
	cmd := exec.Command(command, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func printWorkingDirectory() {
	pwd, err := os.Getwd()
	if err != nil {
		log.Print("Could not get current working directory")
	}
	fmt.Println(pwd)
}

func changeDirectory(path string) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Could not get home directory")
	}
	if path == "~" {
		if err := os.Chdir(homeDir); err != nil {
			fmt.Printf("Cannot change current working directory: %s\n", err)
		}
	} else {
		if err := os.Chdir(path); err != nil {
			fmt.Printf("cd: %s: No such file or directory\n", path)
		}
	}
}

func parseArgs(line string) []string {
	var args []string
	var cur strings.Builder
	var quote rune
	hasToken := false

	for _, r := range line {
		switch {
		case quote == 0 && (r == '\'' || r == '"'):
			quote = r // open quote
			hasToken = true
		case r == quote:
			quote = 0 // close quote
		case r == ' ' && quote == 0:
			if hasToken {
				args = append(args, cur.String())
				cur.Reset()
				hasToken = false
			}
		default:
			cur.WriteRune(r)
			hasToken = true
		}
	}

	if hasToken {
		args = append(args, cur.String())
	}
	return args
}
