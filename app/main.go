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
}

func main() {
	// TODO: Uncomment the code below to pass the first stage

	for {
		fmt.Print("$ ")
		command, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		command = strings.TrimSpace(command)

		if command == "exit" {
			os.Exit(0)
		} else if strings.HasPrefix(command, "echo ") {
			fmt.Println(command[5:])
		} else if strings.HasPrefix(command, "type ") {
			if checkType(command[5:]) {
				fmt.Println(command[5:] + " is a shell builtin")
			} else {
				checkPATH(command[5:])
			}
		} else {
			fmt.Printf("%s: command not found\n", command)
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
