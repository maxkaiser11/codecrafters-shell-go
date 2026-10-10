package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Builtin struct {
	Echo string
	Exit string
}

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Print

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
				fmt.Println("invalid_command: not found")
			}
		} else {
			fmt.Printf("%s: command not found\n", command)
		}

	}

}

func checkType(command string) bool {
	if command == "echo" || command == "exit" {
		return true
	}
	return false
}
