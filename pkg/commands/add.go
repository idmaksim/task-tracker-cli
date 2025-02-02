package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new task",
	Run:   Add,
}

func Add(cmd *cobra.Command, args []string) {
	ensureTaskSpecified(args)
	task := args[0]

	fmt.Println(task)
}

func ensureTaskSpecified(args []string) {
	if len(args) < 1 {
		panic("You must specify a task")
	}
}
