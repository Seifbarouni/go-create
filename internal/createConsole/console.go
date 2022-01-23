package createConsole

import (
	"fmt"
	"os"
	"time"

	colorize "github.com/Seifbarouni/go-create/internal/colorizeText"
	"github.com/Seifbarouni/go-create/internal/helpers"
)

var h *helpers.Helpers = helpers.InitializeHelpers()

func CreateConsoleApp(folderName string) {
	// start timer
	start := time.Now()
	// check if the folder name is valid and does not contain spaces and special characters
	if !h.ValidateFolderName(folderName) {
		colorize.PrintWithColor("Invalid folder name", colorize.Red)
		os.Exit(1)
	}

	if folderName != "." {
		h.CreateFolder(folderName)
	}

	h.AddFolderAndReadme("bin", "# `/bin`\n This folder contains the binary files of the app")
	h.AddFolderAndReadme("cmd", "# `/cmd`\n This folder contains the main.go file of the app")
	h.AddFolderAndReadme("internal", "# `/internal`\n Private application and library code. This is the code you don't want others importing in their applications or libraries.")
	h.AddFolderAndReadme("pkg", "# `/pkg`\n Library code that's ok to use by external applications.")
	// go to cmd folder with the os package
	err := os.Chdir("cmd")
	if err != nil {
		colorize.PrintWithColor("Error moving to folder", colorize.Red)
		os.Exit(1)
	}

	// create the main.go file
	h.CreateFile("main.go", "package main\n\nimport \"fmt\"\n\nfunc main() {\n\tfmt.Println(\"Hello World!\")\n}")

	// go back to the root folder
	err = os.Chdir("..")
	if err != nil {
		colorize.PrintWithColor("Error moving to folder", colorize.Red)
		os.Exit(1)
	}

	h.ExecuteCommand("git", "init")

	// end timer
	elapsed := time.Since(start)
	colorize.PrintWithColor(fmt.Sprintf("\nApp created in %s\n", elapsed), colorize.Green)
}
