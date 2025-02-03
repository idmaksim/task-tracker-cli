package commands

import (
	"fmt"

	"github.com/idmaksim/task-tracker-cli/pkg/json"
	"github.com/idmaksim/task-tracker-cli/pkg/models"
	"github.com/idmaksim/task-tracker-cli/pkg/types"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list of tasks",
	Run:   List,
}

func List(cmd *cobra.Command, args []string) {
	var status string
	if len(args) > 0 {
		status = args[0]
	}

	tasks, err := json.ReadAllData(types.FindAllOptions{
		Status: status,
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	if len(tasks) == 0 {
		fmt.Println("No tasks found")
		return
	}

	PrettyPrint(tasks)
}

func PrettyPrint(tasks []models.Task) {
	fmt.Println("ID\t\tStatus\t\tDescription")
	for _, task := range tasks {
		fmt.Printf("%d\t\t%s\t\t%s\n", task.ID, task.Status, task.Description)
	}
}
