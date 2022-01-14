package createWeb

import (
	"fmt"
	"os"
	"time"

	colorize "github.com/Seifbarouni/go-create/internal/colorizeText"
	h "github.com/Seifbarouni/go-create/internal/helpers"
)

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

func createBackendWebApp(folderName string) {
	h.CreateFile(folderName, "main.go", "package main\n\nimport (\n\t\"fmt\"\n)\n\nfunc main() {\n\tfmt.Println(\"Hello World!\")\n}")
	h.CreateFile(folderName, ".env", "")
	h.CreateFile(folderName, "Dockefile", "")
	h.CreateFile(folderName,".dockerignore", "")
	h.CreateFile(folderName,".gitignore", "")

	addFolderAndReadme("models","# `/models`\n This is where you will create your models")
	addFolderAndReadme("controllers","# `/controllers`\n This is where you will create your controllers")
	addFolderAndReadme("services","# `/services`\n This is where you will create your services")
	addFolderAndReadme("routes","# `/routes`\n This is where you will create your routes")
	addFolderAndReadme("database","# `/database`\n This is where you will create your database")
	addFolderAndReadme("config","# `/config`\n This is where you will create your config")

	
}

func createFullstackWebApp(folderName string) {
	
}

func CreateWebApp(folderName string) {
	// get type of web app from user : fullstack, backend
	colorize.PrintWithColor("What type of web app do you want to create?\n1. fullstack\n2. backend", colorize.Gray)
	var appType string
	fmt.Scanln(&appType)
	if appType == "1" || appType == "fullstack" {
		// create a fullstack web app
		createFullstackWebApp(folderName)
	}else if appType == "2" || appType == "backend" {
		// start timer
		start := time.Now()
		if folderName != "." {
			h.CreateFolder(folderName)
		}
		createBackendWebApp(folderName)
	// execute git init
	colorize.PrintWithColor("Initializing git repository...", colorize.Blue)
	h.ExecuteCommand("git","init")
	// end timer
	elapsed := time.Since(start)
	colorize.PrintWithColor(fmt.Sprintf("Web app created in %s", elapsed), colorize.Green)
	}else{
		colorize.PrintWithColor("Invalid input", colorize.Red)
		os.Exit(1)
	}
}
