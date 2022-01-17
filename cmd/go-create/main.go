package main

import (
	"fmt"
	"os"

	console "github.com/Seifbarouni/go-create/internal/createConsole"
	db "github.com/Seifbarouni/go-create/internal/createDB"
	md "github.com/Seifbarouni/go-create/internal/createModel"
	sv "github.com/Seifbarouni/go-create/internal/createService"
	web "github.com/Seifbarouni/go-create/internal/createWeb"
)

func main() {
	//check if the user has provided 3 arguments
	if len(os.Args) != 3 {
		fmt.Println("Usage: go-create <app-type> <folder-name>\nExample: go-create web .\nAvailable app types: web, console\nOther functionnalities: db <file-name>, crud-service <concerned-model>, model <file-name>")
		os.Exit(1)
	}
	// get the app type and path
	appType := os.Args[1]
	folderOrFileName := os.Args[2]

	// check if the app type is valid with a switch statement
	switch appType {
	case "web":
		web.CreateWebApp(folderOrFileName)
	case "console":
		console.CreateConsoleApp(folderOrFileName)
	case "db":
		db.CreateDB(folderOrFileName)
	case "model":
		md.CreateModel(folderOrFileName)
	case "crud-service":
		sv.CreateService(folderOrFileName)
	default:
		fmt.Println("Invalid app type\nAvailable types: web, console, db, crud-service, model")
		os.Exit(1)
	}
}
