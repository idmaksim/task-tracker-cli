package commands

import (
	"fmt"

	"github.com/idmaksim/task-tracker-cli/pkg/json"
	"github.com/idmaksim/task-tracker-cli/pkg/models"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new task",
	Run:   Add,
}

func Add(cmd *cobra.Command, args []string) {
	ensureTaskSpecified(args)
	ensureNotTooLong(args)

	task := models.NewTask(args[0])

	id, err := json.WriteData(task)
	if err != nil {
		panic(err)
	}

	fmt.Println("Task added successfully with ID:", id)
}

func ensureTaskSpecified(args []string) {
	if len(args) < 1 {
		panic("You must specify a task description")
	}
}

func ensureNotTooLong(args []string) {
	if len(args) > 1 {
		panic("Too many arguments. Task description should be in quotes")
	}
}
