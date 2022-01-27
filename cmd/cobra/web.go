package cobra

import (
	w "github.com/Seifbarouni/go-create/internal/createWeb"

	"github.com/spf13/cobra"
)

// still have bugs to fix

var appType string = ""

func init() {
	// add flag to get the app type
	rootCmd.PersistentFlags().StringVarP(&appType, "type", "t", "", "can be backend or fullstack")
	rootCmd.MarkFlagRequired("type")
	rootCmd.AddCommand(webCmd)
}

var webCmd = &cobra.Command{
	Use:   "web [folder-name]",
	Short: "Create a web app",
	Long:  `Create a backend web app or a full stack web app.`,
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		w.CreateWebApp(args[0], appType)
	},
}
