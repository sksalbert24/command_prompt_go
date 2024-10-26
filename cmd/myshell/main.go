package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Fprint

func main() {
	// Uncomment this block to pass the first stage
	for {
		fmt.Fprint(os.Stdout, "$ ")

		// Wait for user input
		command, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		programs := strings.SplitN(command, " ", 2)

		switch programs[0] {
		case "exit":
			arguments := strings.Split(programs[1], " ")
			if len(arguments) > 2 {
				errors.New("Wrong Arguments")
			}
			argument, error := strconv.Atoi(arguments[0])
			if error != nil {
				errors.New("Wrong Argument")
			}
			os.Exit(argument)
		case "echo":
			fmt.Fprintf(os.Stdout, "%s", programs[1])
		default:
			fmt.Fprintf(os.Stdout, "%s: command not found\n", strings.TrimRight(command, "\n"))
		}
	}
}
