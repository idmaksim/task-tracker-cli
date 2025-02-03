package commands

import (
	"github.com/idmaksim/task-tracker-cli/internal/delivery/cli/handlers"
	"github.com/spf13/cobra"
)

type Commands struct {
	handler *handlers.TaskHandler
	rootCmd *cobra.Command
}

func NewCommands(handler *handlers.TaskHandler) *Commands {
	cmd := &Commands{
		handler: handler,
		rootCmd: &cobra.Command{
			Use:   "task-cli",
			Short: "Task tracker cli application",
		},
	}
	cmd.setup()
	return cmd
}

func (c *Commands) Execute() error {
	return c.rootCmd.Execute()
}

func (c *Commands) setup() {
	c.rootCmd.AddCommand(
		c.newAddCmd(),
		c.newListCmd(),
		c.newUpdateCmd(),
		c.newDeleteCmd(),
		c.newMarkInProgressCmd(),
		c.newMarkDoneCmd(),
	)
}
