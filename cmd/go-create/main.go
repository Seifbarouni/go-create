package main

import (
	"fmt"
	"os"

	console "github.com/Seifbarouni/go-create/internal/createConsole"
	web "github.com/Seifbarouni/go-create/internal/createWeb"
)

func main() {
	//check if the user has provided 3 arguments
	if len(os.Args) != 3 {
		fmt.Println("Usage: go-create <app-type> <folder-name>\nExample: go-create web .\nAvailable app types: web, console")
		os.Exit(1)
	}
	// get the app type and path
	appType := os.Args[1]
	folderName := os.Args[2]

	// check if the app type is valid with a switch statement
	switch appType {
	case "web":
		web.CreateWebApp(folderName)
	case "console":
	 	console.CreateConsoleApp(folderName)	
	default:
		fmt.Println("Invalid app type\nAvailable app types: web, console")
		os.Exit(1)
	}
}