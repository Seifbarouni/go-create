package createConsole

import (
	"fmt"
	"os"
	"time"

	colorize "github.com/Seifbarouni/go-create/internal/colorizeText"
	h "github.com/Seifbarouni/go-create/internal/helpers"
)


func validateFolderName(folderName string) bool {
	// check if the folder name is valid and does not contain spaces and special characters
	if len(folderName) < 1 || len(folderName) > 50 {
		return false
	}
	
	for _, char := range folderName {
		if char == ' ' || char == '\\' || char == '/' || char == ':' || char == '*' || char == '?' || char == '"' || char == '<' || char == '>' || char == '|' {
			return false
		}
	}
	
	return true
}





// create folder and add README.md file
func addFolderAndReadme(folderName string, content string) {
	colorize.PrintWithColor(fmt.Sprintf("Creating %s folder...", folderName), colorize.Blue)
	// create the folder
	h.CreateFolder(folderName)
	
	// create the file
	h.CreateFile(folderName, "README.md",content)

	// go back to the root folder
	err := os.Chdir("..")
	if err != nil {
		colorize.PrintWithColor("Error moving to folder", colorize.Red)
		os.Exit(1)
	}
	colorize.PrintWithColor(fmt.Sprintf("Folder %s created", folderName), colorize.Green)
}

func CreateConsoleApp(folderName string) {
	// start timer
	start := time.Now()
	// check if the folder name is valid and does not contain spaces and special characters
	if !validateFolderName(folderName) {
		colorize.PrintWithColor("Invalid folder name", colorize.Red)
		os.Exit(1)
	}

	if folderName != "." {
		h.CreateFolder(folderName)
	}

	addFolderAndReadme("bin","# `/bin`\n This folder contains the binary files of the app")
	addFolderAndReadme("cmd","# `/cmd`\n This folder contains the main.go file of the app")
	addFolderAndReadme("internal","# `/internal`\n Private application and library code. This is the code you don't want others importing in their applications or libraries.")
	addFolderAndReadme("pkg","# `/pkg`\n Library code that's ok to use by external applications.")

	// end timer
	elapsed := time.Since(start)
	colorize.PrintWithColor(fmt.Sprintf("\nApp created in %s\n", elapsed), colorize.White)
}