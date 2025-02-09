package commands

import (
	"errors"

	"github.com/spf13/cobra"
)

func (c *Commands) newAddCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "add [description]",
		Short: "Add a new task",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return errors.New("description is required")
			}
			return c.handler.CreateTask(args[0])
		},
	}
}
