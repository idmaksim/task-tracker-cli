package commands

import (
	"strconv"

	"github.com/spf13/cobra"
)

func (c *Commands) newUpdateCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "update [id] [description]",
		Short: "Update task description",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := strconv.Atoi(args[0])
			if err != nil {
				return err
			}
			return c.handler.UpdateTaskDescription(id, args[1])
		},
	}
}
