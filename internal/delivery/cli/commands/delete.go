package commands

import (
	"strconv"

	"github.com/spf13/cobra"
)

func (c *Commands) newDeleteCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "delete [id]",
		Short: "Delete a task",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := strconv.Atoi(args[0])
			if err != nil {
				return err
			}
			return c.handler.DeleteTask(id)
		},
	}
}
