package createWeb

import (
	"fmt"
	"os"
	"time"

	colorize "github.com/Seifbarouni/go-create/internal/colorizeText"
	h "github.com/Seifbarouni/go-create/internal/helpers"
)



func createBackendWebApp(folderName string) {
	h.CreateFile(folderName, "main.go", "package main\n\nimport (\n\t\"fmt\"\n)\n\nfunc main() {\n\tfmt.Println(\"Hello World!\")\n}")
	h.CreateFile(folderName, ".env", "")
	h.CreateFile(folderName, "Dockefile", "")
	h.CreateFile(folderName,".dockerignore", "")
	h.CreateFile(folderName,".gitignore", "")

	h.AddFolderAndReadme("models","# `/models`\n This is where you will create your models")
	h.AddFolderAndReadme("controllers","# `/controllers`\n This is where you will create your controllers")
	h.AddFolderAndReadme("services","# `/services`\n This is where you will create your services")
	h.AddFolderAndReadme("routes","# `/routes`\n This is where you will create your routes")
	h.AddFolderAndReadme("database","# `/database`\n This is where you will create your database")
	h.AddFolderAndReadme("config","# `/config`\n This is where you will create your config")

	
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
