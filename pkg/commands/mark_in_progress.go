package commands

import (
	"fmt"
	"strconv"

	"github.com/idmaksim/task-tracker-cli/pkg/constants"
	"github.com/idmaksim/task-tracker-cli/pkg/json"
	"github.com/spf13/cobra"

)

var markInProgressCmd = &cobra.Command{
	Use:   "mark-in-progress [id]",
	Short: "Mark task as in progress",
	Run:   MarkInProgress,
}

func MarkInProgress(cmd *cobra.Command, args []string) {
	if len(args) != 1 {
		fmt.Println("You must specify a task ID")
		return
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("Invalid task ID")
		return
	}

	err = json.UpdateTaskStatus(id, constants.InProgress)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Task marked as in progress")
}
