package cobra

import (
	database "github.com/Seifbarouni/go-create/internal/createDB"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(dbCmd)
}

var dbCmd = &cobra.Command{
	Use:                   "db [file-name]",
	Short:                 "Create a database",
	Long:                  `Create a database file with the default postgresql configuration.`,
	Args:                  cobra.ExactArgs(1),
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		database.CreateDB(args[0])
	},
}
