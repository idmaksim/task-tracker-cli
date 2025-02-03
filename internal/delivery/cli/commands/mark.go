package commands

import (
	"strconv"

	"github.com/idmaksim/task-tracker-cli/pkg/constants"
	"github.com/spf13/cobra"

)

func (c *Commands) newMarkInProgressCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "mark-in-progress [id]",
		Short: "Mark task as in progress",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := strconv.Atoi(args[0])
			if err != nil {
				return err
			}
			return c.handler.UpdateTaskStatus(id, constants.InProgress)
		},
	}
}

func (c *Commands) newMarkDoneCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "mark-done [id]",
		Short: "Mark task as done",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := strconv.Atoi(args[0])
			if err != nil {
				return err
			}
			return c.handler.UpdateTaskStatus(id, constants.Done)
		},
	}
}
