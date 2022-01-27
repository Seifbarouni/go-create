package cobra

import (
	console "github.com/Seifbarouni/go-create/internal/createConsole"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(consoleCmd)
}

var consoleCmd = &cobra.Command{
	Use:                   "cli [folder-name or .]",
	Short:                 "Create a cli application",
	Args:                  cobra.ExactArgs(1),
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		console.CreateConsoleApp(args[0])
	},
}
