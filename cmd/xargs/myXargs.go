package main

import (
	"fmt"
	"github.com/fanfaronDo/tools/pkg/xargs"
	"os"
	"os/exec"
	"strings"
)

func main() {
	var command string
	arguments := make([]string, 0)

	xargs.ReadLine(&arguments)

	if len(os.Args) <= 1 {
		for _, arg := range arguments {
			fmt.Print(arg + " ")
		}
		fmt.Println()
		return
	}

	args := os.Args[1:]
	command += args[0]
	if len(args[1:]) > 0 {
		for _, arg := range args[1:] {
			command += " " + arg
		}
	}

	fmt.Println(command, strings.Join(arguments, " "))

	cmd := exec.Command(command, arguments...)
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
		return
	}
}
