package commands

import (
	"github.com/spf13/cobra"
)

func (c *Commands) newAddCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "add [description]",
		Short: "Add a new task",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return c.handler.CreateTask(args[0])
		},
	}
}
