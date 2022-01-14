package createWeb

import (
	"fmt"
	"os"

	colorize "github.com/Seifbarouni/go-create/internal/colorizeText"
	h "github.com/Seifbarouni/go-create/internal/helpers"
)

func createBackendWebApp(folderName string) {
	file, err := os.Create("main.go")
	if err != nil {
		colorize.PrintWithColor("Error creating main.go file", colorize.Red)
		os.Exit(1)
	}
	file.Close()

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
	}
	if appType == "2" || appType == "backend" {
		// create a backend web app
		if folderName != "." {
			h.CreateFolder(folderName)
		}
		createBackendWebApp(folderName)
	}
	fmt.Println("Creating web app in ", folderName)
}
