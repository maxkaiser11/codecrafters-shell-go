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

		fields := strings.Fields(line)
		if len(fields) == 0 {
			continue
		}

		name, args := fields[0], fields[1:]

		switch name {
		case "exit":
			os.Exit(0)
		case "echo":
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
			if _, err := exec.LookPath(name); err != nil {
				fmt.Printf("%s: command not found\n", name)
				continue
			}
			if err := execCommand(name, args...); err != nil {
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
	if path == "~" {
		if err := os.Chdir("/"); err != nil {
			fmt.Printf("Cannot change current working directory: %s\n", err)
		}
	} else {
		if err := os.Chdir(path); err != nil {
			fmt.Printf("cd: %s: No such file or directory\n", path)
		}
	}
}
