package main

import (
	"fmt"
	"task-tracker-cli/pkg/commands"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	commands.Execute()
}
