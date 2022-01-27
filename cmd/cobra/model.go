package cobra

import (
	m "github.com/Seifbarouni/go-create/internal/createModel"

	"github.com/spf13/cobra"
)

func init() {
  rootCmd.AddCommand(modelCmd)
}

var modelCmd = &cobra.Command{
  Use:   "model [model-name]",
  Short: "Create a model",
  Args: cobra.ExactArgs(1),
  DisableFlagsInUseLine: true,
  Run: func(cmd *cobra.Command, args []string) {
	m.CreateModel(args[0])
  },
}