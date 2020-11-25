package main

import (
	"bufio"
	"fmt"
	"os"
)

var shell ShellCmd

func init() {}

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("$ ")
		cmdString, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		shell.runCommand(cmdString)
	}
}
