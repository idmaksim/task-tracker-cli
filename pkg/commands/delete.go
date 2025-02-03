package commands

import (
	"fmt"
	"strconv"

	"github.com/idmaksim/task-tracker-cli/pkg/json"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a task",
	Run:   Delete,
}

func Delete(cmd *cobra.Command, args []string) {
	if len(args) != 1 {
		fmt.Println("You must specify a task ID")
		return
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("Invalid task ID")
		return
	}

	err = json.DeleteData(id)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Task deleted successfully")
}
