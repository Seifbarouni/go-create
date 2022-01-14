package helpers

import (
	"fmt"
	"os"

	colorize "github.com/Seifbarouni/go-create/internal/colorizeText"
)

func CreateFolder(folderName string) {
	// check if the folder exists
	if _, err := os.Stat(folderName); err == nil {
		colorize.PrintWithColor(fmt.Sprintf("Folder %s already exists\n", folderName), colorize.Yellow)
		os.Exit(1)
	}
	// create the folder
	err := os.Mkdir(folderName, 0755)
	if err != nil {
		colorize.PrintWithColor(err.Error(), colorize.Red)
		os.Exit(1)
	}
	// move to the folder
	err = os.Chdir(folderName)
	if err != nil {
		colorize.PrintWithColor("Error moving to folder", colorize.Red)
		os.Exit(1)
	}
}

// create README.md file and write the content
func CreateReadme(folderName string, content string) {
	// create the file
	file, err := os.Create("README.md")
	if err != nil {
		colorize.PrintWithColor("Error creating README.md file", colorize.Red)
		os.Exit(1)
	}
	defer file.Close()

	// write the content
	_, err = file.WriteString(content)
	if err != nil {
		colorize.PrintWithColor("Error writing to README.md file", colorize.Red)
		os.Exit(1)
	}
}