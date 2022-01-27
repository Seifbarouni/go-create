package cobra

import (
	w "github.com/Seifbarouni/go-create/internal/createWeb"

	"github.com/spf13/cobra"
)

func init() {
	webCmd.Flags().StringVarP(&appType, "type", "t", "backend", "can be backend or fullstack")
	rootCmd.MarkFlagRequired("type")
	rootCmd.AddCommand(webCmd)
}

var (
	webCmd = &cobra.Command{
		Use:   "web [folder-name]",
		Short: "Create a web app",
		Long:  `Create a backend web app or a full stack web app.`,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			w.CreateWebApp(args[0], appType)
		},
	}
	appType string = ""
)
