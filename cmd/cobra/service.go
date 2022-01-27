package cobra

import (
	s "github.com/Seifbarouni/go-create/internal/createService"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(serviceCmd)
}

var serviceCmd = &cobra.Command{
	Use:                   "service [concerned-model]",
	Short:                 "Create a service",
	Long:                  `Create a crud service file to a specific model.`,
	Args:                  cobra.ExactArgs(1),
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		s.CreateService(args[0])
	},
}
