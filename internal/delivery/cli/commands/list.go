package commands

import (
	"github.com/spf13/cobra"
)

func (c *Commands) newListCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "list [status]",
		Short: "List tasks",
		RunE: func(cmd *cobra.Command, args []string) error {
			status := ""
			if len(args) > 0 {
				status = args[0]
			}
			return c.handler.ListTasks(status)
		},
	}
}
