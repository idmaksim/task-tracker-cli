package commands

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "task-cli",
	Short: "Task tracker cli application",
	Run:   Root,
}

func Root(cmd *cobra.Command, args []string) {
	cmd.Usage()
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.AddCommand(addCmd)
}
