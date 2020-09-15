package main

import (
	"fmt"
	"os"
	"strings"
)

// ShellMenu declares the name of each menu
type ShellMenu string

// ShellResource declares the name of shell resources
type ShellResource string

// ShellCmd is the struct of shell
type ShellCmd struct {
	pools []*Pool
}

const (
	showOptionsMenu   ShellMenu = "showOptionsMenu"
	createOptionsMenu ShellMenu = "createOptionsMenu"
	checkOptionsMenu  ShellMenu = "checkOptionsMenu"
	deleteOptionsMenu ShellMenu = "deleteOptionsMenu"
)

const (
	poolResource ShellResource = "poolResource"
)

// createMenu returns menu content according to the input string
func createMenu(menu ShellMenu) {
	var sb strings.Builder
	switch menu {
	case showOptionsMenu:
		sb.WriteString("No available show options...\n")
		break
	case createOptionsMenu:
		sb.WriteString("No available create options...\n")
		break
	case checkOptionsMenu:
		sb.WriteString("No available check options...\n")
		break
	case deleteOptionsMenu:
		sb.WriteString("No available delete options...\n")
		break
	}
	fmt.Print(sb.String())
}

// invokeResource acts differently according to received resource
func invokeResource(resource ShellResource, actions interface{}) {
	switch resource {
	case poolResource:
		break
	}
}

func (shell ShellCmd) runCommand(cmd string) error {
	cmd = strings.TrimSuffix(cmd, "\n")
	arrCommandStr := strings.Fields(cmd)
	if len(arrCommandStr) == 0 {
		return nil
	}
	switch arrCommandStr[0] {
	case "exit":
		os.Exit(0)
	case "show":
		// display all options for "show" command
		if len(arrCommandStr) == 1 {
			createMenu(showOptionsMenu)
			return nil
		}
		res := arrCommandStr[1]
		switch res {
		case "pools":
			fmt.Println("pools empty")
			break
		default:
			fmt.Println("Option \"" + res + "\" is not recognized")
		}
		break
	case "create":
		cmdLen := len(arrCommandStr)
		// display all options for "create" command
		if cmdLen == 1 {
			createMenu(createOptionsMenu)
			return nil
		}
		res := arrCommandStr[1]
		switch res {
		// pool
		case "pool":
			if cmdLen == 2 {
				fmt.Println("A \"name\" is required to create pool")
				return nil
			}
			if cmdLen == 3 && strings.Trim(arrCommandStr[2], " ") == "" {
				fmt.Println("Pool \"name\" can not be empty")
				return nil
			}
			pool, err := createPool(arrCommandStr[2])
			if err != nil {
				return err
			}
			// TODO: handle pool create
			invokeResource(poolResource, pool)
			break
		}
		break
	case "check":
		// display all options for "check" command
		if len(arrCommandStr) == 1 {
			createMenu(checkOptionsMenu)
			return nil
		}
		break
	case "delete":
		// display all options for "delete" command
		if len(arrCommandStr) == 1 {
			createMenu(deleteOptionsMenu)
			return nil
		}
		break
	}
	return nil
}
