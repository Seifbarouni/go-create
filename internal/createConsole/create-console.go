package createConsole

import (
	"fmt"
	"os"
	"sync"
	"time"

	colorize "github.com/Seifbarouni/go-create/internal/colorizeText"
)

var wg sync.WaitGroup

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

func createFolder(folderName string) {
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
func createReadme(folderName string, content string) {
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

// create folder and add README.md file
func addFolderAndReadme(folderName string, content string) {
	defer wg.Done()
	colorize.PrintWithColor(fmt.Sprintf("Creating %s folder...", folderName), colorize.Blue)
	// create the folder
	createFolder(folderName)
	
	// create the file
	createReadme(folderName, content)

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
	wg.Add(4)
	// check if the folder name is valid and does not contain spaces and special characters
	if !validateFolderName(folderName) {
		colorize.PrintWithColor("Invalid folder name", colorize.Red)
		os.Exit(1)
	}

	if folderName != "." {
		createFolder(folderName)
	}

	addFolderAndReadme("bin","# `/bin`\n This folder contains the binary files of the app")
	addFolderAndReadme("cmd","# `/cmd`\n This folder contains the main.go file of the app")
	addFolderAndReadme("internal","# `/internal`\n Private application and library code. This is the code you don't want others importing in their applications or libraries.")
	addFolderAndReadme("pkg","# `/pkg`\n Library code that's ok to use by external applications.")

	wg.Wait()
	// end timer
	elapsed := time.Since(start)
	colorize.PrintWithColor(fmt.Sprintf("\nApp created in %vs\n", elapsed.Seconds()), colorize.White)
}