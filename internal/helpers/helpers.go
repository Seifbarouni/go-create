package helpers

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"

	colorize "github.com/Seifbarouni/go-create/internal/colorizeText"
	gen "github.com/Seifbarouni/go-create/internal/generators"
)

// Helpers struct
type Helpers struct{}

// InitializeHelpers initializes the helpers
func InitializeHelpers() *Helpers {
	return &Helpers{}
}

// CreateFolder creates a folder
func (*Helpers) CreateFolder(folderName string) {
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

// CreateFile creates a file
func (*Helpers) CreateFile(fileName string, content string) {
	// create the file
	file, err := os.Create(fileName)
	if err != nil {
		colorize.PrintWithColor(fmt.Sprintf("Error creating %s file", fileName), colorize.Red)
		os.Exit(1)
	}
	defer file.Close()

	// write the content
	if content != "" {
		_, err = file.WriteString(content)
		if err != nil {
			colorize.PrintWithColor(fmt.Sprintf("Error writing to file %s", fileName), colorize.Red)
			os.Exit(1)
		}
	}
}

// ExecuteCommand executes a command
func (*Helpers) ExecuteCommand(command ...string) {
	cmd := exec.Command(command[0], command[1:]...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	err := cmd.Run()
	if err != nil {
		colorize.PrintWithColor(fmt.Sprintf("Error executing command %s", command), colorize.Red)
		os.Exit(1)
	}

}

// AddFolderAndReadme creates folder and adds a README.md file
func (h *Helpers) AddFolderAndReadme(folderName string, content string) {
	colorize.PrintWithColor(fmt.Sprintf("Creating %s folder...\n", folderName), colorize.Purple)
	// create the folder
	h.CreateFolder(folderName)

	// create the file
	h.CreateFile("README.md", content)

	// go back to the root folder
	err := os.Chdir("..")
	if err != nil {
		colorize.PrintWithColor("Error moving to folder", colorize.Red)
		os.Exit(1)
	}
	colorize.PrintWithColor(fmt.Sprintf("Folder %s created\n", folderName), colorize.White)
}

// GetSeperator returns the seperator for the current OS
func (*Helpers) GetSeperator() string {
	os := runtime.GOOS
	if os == "windows" {
		return "\\"
	}
	return "/"
}

// ValidateFolderName validates the folder name
func (*Helpers) ValidateFolderName(folderName string) bool {
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

// AddPublicAndPrivateRoutes adds public and private routes
func (h *Helpers) AddPublicAndPrivateRoutes() {
	// go to the routes folder
	err := os.Chdir("routes")
	if err != nil {
		colorize.PrintWithColor("Error moving to folder", colorize.Red)
		os.Exit(1)
	}
	// create publicRoutes.go file
	h.CreateFile("publicRoutes.go", gen.PublicRoute)
	// create privateRoutes.go file
	h.CreateFile("privateRoutes.go", gen.PrivateRoute)
	// go to the parent folder
	err = os.Chdir("..")
	if err != nil {
		colorize.PrintWithColor("Error moving to folder", colorize.Red)
		os.Exit(1)
	}
}

// GetModuleName returns the module name
func (*Helpers) GetModuleName() string {
	// go to the root folder
	err := os.Chdir("..")
	if err != nil {
		colorize.PrintWithColor("Error moving to folder", colorize.Red)
		os.Exit(1)
	}
	// read the first line of the go.mod file
	file, err := os.Open("go.mod")
	if err != nil {
		colorize.PrintWithColor("Cannot find go.mod file, generating default name...\n", colorize.Bold)
		return "your/module/name"
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Scan()

	firstLine := scanner.Text()

	return strings.Split(firstLine, " ")[1]

}
