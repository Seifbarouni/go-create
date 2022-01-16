package createWeb

import (
	"fmt"
	"os"
	"time"

	colorize "github.com/Seifbarouni/go-create/internal/colorizeText"
	h "github.com/Seifbarouni/go-create/internal/helpers"
)

func createBackendWebApp(folderName string) {
	h.CreateFile("main.go", "package main\n\nimport (\n\t\"fmt\"\n)\n\nfunc main() {\n\tfmt.Println(\"Hello World!\")\n}")
	h.CreateFile(".env", "")
	h.CreateFile("Dockefile", "")
	h.CreateFile(".dockerignore", "")
	h.CreateFile(".gitignore", "")

	h.AddFolderAndReadme("models", "# `/models`\n This is where you will create your models")
	h.AddFolderAndReadme("controllers", "# `/controllers`\n This is where you will create your controllers")
	h.AddFolderAndReadme("services", "# `/services`\n This is where you will create your services")
	h.AddFolderAndReadme("routes", "# `/routes`\n This is where you will create your routes")
	h.AddFolderAndReadme("database", "# `/database`\n This is where you will create your database")
	h.AddFolderAndReadme("config", "# `/config`\n This is where you will create your config")

}

func createFrontendApp() {
	// get type of frontend app from user : react, nextjs, vue
	colorize.PrintWithColor("What type of frontend app do you want to create?\n1. react\n2. nextjs\n3. vue\n4. svelte\n", colorize.Gray)
	var frontendType string
	fmt.Print("-->")
	fmt.Scanln(&frontendType)

	if frontendType == "1" || frontendType == "react" {
		h.ExecuteCommand("npx", "create-react-app", "frontend", "--template", "typescript")
	} else if frontendType == "2" || frontendType == "nextjs" {
		h.ExecuteCommand("npx", "create-next-app@latest", "--typescript", "frontend")
	} else if frontendType == "3" || frontendType == "vue" {
		h.ExecuteCommand("npx", "@vue/cli", "create", "--default", "frontend")
	} else if frontendType == "4" || frontendType == "svelte" {
		h.ExecuteCommand("npx", "degit", "sveltejs/template", "frontend")
	} else {
		colorize.PrintWithColor("Invalid input", colorize.Red)
		os.Exit(1)
	}
}

func createFullstackWebApp() {
	// create the backend folder
	h.CreateFolder("backend")
	// generate the backend files
	createBackendWebApp("backend")
	err := os.Chdir("..")
	if err != nil {
		colorize.PrintWithColor("Error moving to folder", colorize.Red)
		os.Exit(1)
	}
	createFrontendApp()
}

func CreateWebApp(folderName string) {
	// get type of web app from user : fullstack, backend
	colorize.PrintWithColor("What type of web app do you want to create?\n1. fullstack\n2. backend\n", colorize.Gray)
	var appType string
	fmt.Print("-->")
	fmt.Scanln(&appType)
	if appType == "1" || appType == "fullstack" {
		start := time.Now()
		if folderName != "." {
			h.CreateFolder(folderName)
		}
		// create a fullstack web app
		createFullstackWebApp()
		// end timer
		elapsed := time.Since(start)
		colorize.PrintWithColor(fmt.Sprintf("Web app created in %s", elapsed), colorize.Green)
	} else if appType == "2" || appType == "backend" {
		// start timer
		start := time.Now()
		if folderName != "." {
			h.CreateFolder(folderName)
		}
		createBackendWebApp(folderName)
		// execute git init
		colorize.PrintWithColor("Initializing git repository...", colorize.Blue)
		h.ExecuteCommand("git", "init")
		// end timer
		elapsed := time.Since(start)
		colorize.PrintWithColor(fmt.Sprintf("Web app created in %s", elapsed), colorize.Green)
	} else {
		colorize.PrintWithColor("Invalid input", colorize.Red)
		os.Exit(1)
	}
}
