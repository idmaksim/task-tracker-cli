package main

import (
	"fmt"

	"github.com/idmaksim/task-tracker-cli/pkg/commands"
	"github.com/idmaksim/task-tracker-cli/pkg/json"
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
