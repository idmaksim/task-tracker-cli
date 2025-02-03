package commands

import (
	"fmt"
	"strconv"

	"github.com/idmaksim/task-tracker-cli/pkg/json"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update [id] [description]",
	Short: "Update task description",
	Run:   Update,
}

func Update(cmd *cobra.Command, args []string) {
	if len(args) != 2 {
		fmt.Println("You must specify task ID and new description")
		return
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("Invalid task ID")
		return
	}

	description := args[1]
	err = json.UpdateTaskDescription(id, description)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Task updated successfully")
}
