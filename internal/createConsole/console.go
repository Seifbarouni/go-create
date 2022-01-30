package createConsole

import (
	"fmt"
	"os"
	"time"

	colorize "github.com/Seifbarouni/go-create/internal/colorizeText"
	gen "github.com/Seifbarouni/go-create/internal/generators"
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

	// get the moddule name from the user
	moduleName := ""
	for {
		// get module name from user
		colorize.PrintWithColor("module name : ", colorize.Gray)
		fmt.Scanln(&moduleName)
		// if the module name is empty recall the function
		if moduleName != "" {
			break
		}
	}
	h.ExecuteCommand("go", "mod", "init", moduleName)

	h.AddFolderAndReadme("bin", gen.BinREADME)
	h.AddFolderAndReadme("cmd", gen.CmdREADME)
	h.AddFolderAndReadme("internal", gen.InternalREADME)
	h.AddFolderAndReadme("pkg", gen.PkgREADME)
	// go to cmd folder with the os package
	err := os.Chdir("cmd")
	if err != nil {
		colorize.PrintWithColor("Error moving to folder", colorize.Red)
		os.Exit(1)
	}

	// create the main.go file
	h.CreateFile("main.go", gen.ConsoleMain)

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
