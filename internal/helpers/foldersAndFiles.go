package helpers

import (
	"fmt"
	"os"
	"os/exec"

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

func CreateFile(folderName string, fileName string,content string) {
	// create the file
	file, err := os.Create(fileName)
	if err != nil {
		colorize.PrintWithColor(fmt.Sprintf("Error creating %s file",fileName), colorize.Red)
		os.Exit(1)
	}
	defer file.Close()

	// write the content
	if content != "" {
		_, err = file.WriteString(content)
		if err != nil {
			colorize.PrintWithColor(fmt.Sprintf("Error writing to file %s",fileName), colorize.Red)
			os.Exit(1)
		}
	}
}

func ExecuteCommand(command ...string){
	cmd := exec.Command(command[0], command[1:]...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	err := cmd.Run()
	if err != nil {
		colorize.PrintWithColor(fmt.Sprintf("Error executing command %s",command), colorize.Red)
		fmt.Println(err)
		os.Exit(1)
	}
    
}

// create folder and add README.md file
func AddFolderAndReadme(folderName string, content string) {
	colorize.PrintWithColor(fmt.Sprintf("Creating %s folder...", folderName), colorize.Blue)
	// create the folder
	CreateFolder(folderName)
	
	// create the file
	CreateFile(folderName, "README.md",content)

	// go back to the root folder
	err := os.Chdir("..")
	if err != nil {
		colorize.PrintWithColor("Error moving to folder", colorize.Red)
		os.Exit(1)
	}
	colorize.PrintWithColor(fmt.Sprintf("Folder %s created", folderName), colorize.Green)
}