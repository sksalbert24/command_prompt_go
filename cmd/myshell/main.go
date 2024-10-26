package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type BuildInEnum int

const (
	Exit BuildInEnum = iota
	Echo
	Type
)

var builtin_functions = []string{
	"exit",
	"echo",
	"type",
}

func contains(array []string, value string) bool {
	for _, v := range array {
		if v == value {
			return true
		}
	}
	return false
}

func contain_command_in_dir(path string, command string) (bool, error) {
	files, err := os.ReadDir(path)
	if err != nil {
		return false, err
	}
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		if file.Name() == command {
			fmt.Fprintf(os.Stdout, "%s is %s\n", command, path+"/"+file.Name())
			return true, nil
		}
	}
	return false, nil
}

func search_path(command string) bool {
	path_list := strings.Split(os.Getenv("PATH"), ":")
	for _, value := range path_list {
		file := filepath.Join(value, command)
		if _, err := os.Stat(file); err == nil {
			fmt.Println(command + " is " + file)
			return true
		}
	}
	return false
}

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
		case builtin_functions[Exit]:
			arguments := strings.Split(programs[1], " ")
			if len(arguments) > 2 {
				errors.New("Wrong Arguments")
			}
			argument, error := strconv.Atoi(arguments[0])
			if error != nil {
				errors.New("Wrong Argument")
			}
			os.Exit(argument)
		case builtin_functions[Echo]:
			fmt.Fprintf(os.Stdout, "%s", programs[1])
		case builtin_functions[Type]:
			arguments := strings.Split(programs[1], " ")
			if len(arguments) > 2 {
				errors.New("Wrong Arguments")
			}
			argument := strings.TrimRight(arguments[0], "\n")
			if contains(builtin_functions, argument) {
				fmt.Fprintf(os.Stdout, "%s is a shell builtin\n", argument)
			} else {
				if search_path(argument) {

				} else {
					fmt.Fprintf(os.Stdout, "%s: not found\n", argument)
				}
			}
		default:
			fmt.Fprintf(os.Stdout, "%s: command not found\n", strings.TrimRight(command, "\n"))
		}
	}
}
