package commands

import (
	"fmt"
	"github.com/spf13/cobra"
	"task-tracker-cli/pkg/json"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list of tasks",
	Run:   List,
}

func List(cmd *cobra.Command, args []string) {
	tasks, err := json.ReadAllData()
	if err != nil {
		fmt.Println(err)
		return
	}

	if len(tasks) == 0 {
		fmt.Println("No tasks found")
		return
	}

	fmt.Println(tasks)
}
