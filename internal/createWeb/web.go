package createWeb

import (
	"fmt"
	"os"
	"time"

	colorize "github.com/Seifbarouni/go-create/internal/colorizeText"
	gen "github.com/Seifbarouni/go-create/internal/generators"
	"github.com/Seifbarouni/go-create/internal/helpers"
)

var h *helpers.Helpers = helpers.InitializeHelpers()

func initializeModuleWithFiber(moduleName string) {
	// go mod init with the module name
	h.ExecuteCommand("go", "mod", "init", moduleName)
	// go get -u github.com/gofiber/fiber/v2
	h.ExecuteCommand("go", "get", "-u", gen.FiberPackage)
}

func createBackendWebApp(folderName string) {
	moduleName := ""
	// get module name from user
	colorize.PrintWithColor("module name : ", colorize.Gray)
	fmt.Scanln(&moduleName)
	// if the module name is empty recall the function
	if moduleName == "" {
		createBackendWebApp(folderName)
	}
	initializeModuleWithFiber(moduleName)
	// initialize go-fiber in the main.go file
	h.CreateFile("main.go", gen.GenerateMainWeb(moduleName))

	h.CreateFile(".env",gen.Env )
	h.CreateFile("Dockefile", gen.Dockerfile)
	h.CreateFile(".dockerignore", gen.Dockerignore)
	h.CreateFile(".gitignore", gen.Gitignore)

	h.AddFolderAndReadme("models", gen.ModelREADME)
	h.AddFolderAndReadme("controllers",gen.ControllerREADME)
	h.AddFolderAndReadme("services", gen.ServiceREADME)
	h.AddFolderAndReadme("routes", gen.RouteREADME)
	h.AddFolderAndReadme("database", gen.DatabaseREADME)
	h.AddFolderAndReadme("config", gen.ConfigREADME)

	h.AddPublicAndPrivateRoutes()

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

func CreateWebApp(folderName string, appType string) {
	if appType == "1" || appType == "fullstack" {
		start := time.Now()
		if !h.ValidateFolderName(folderName) {
			colorize.PrintWithColor("Invalid folder name", colorize.Red)
			os.Exit(1)
		}
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
		if !h.ValidateFolderName(folderName) {
			colorize.PrintWithColor("Invalid folder name", colorize.Red)
			os.Exit(1)
		}
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
