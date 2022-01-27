package cobra

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "go-create",
	Short: "A simple CLI tool to create a new Go application",
	Long: `A simple CLI tool that helps me create my apps (console apps, back-end web apps, full-stack web apps).`, 
}


func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

