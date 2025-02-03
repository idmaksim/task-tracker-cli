package main

import (
	"fmt"
	"task-tracker-cli/pkg/commands"
	"task-tracker-cli/pkg/json"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	json.StartUp()
	commands.Execute()
}
