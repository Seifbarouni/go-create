package createDB

import (
	"os"
	"strings"

	colorize "github.com/Seifbarouni/go-create/internal/colorizeText"
	gen "github.com/Seifbarouni/go-create/internal/generators"
	"github.com/Seifbarouni/go-create/internal/helpers"
)

var h *helpers.Helpers = helpers.InitializeHelpers()

func CreateDB(fileName string) {
	// check if the user has provided a valid file name
	if fileName == "" {
		colorize.PrintWithColor("You must provide a valid file name\n", colorize.Red)
		os.Exit(1)
	}
	// check if the file name is valid
	if !strings.HasSuffix(fileName, ".go") {
		fileName += ".go"
	}

	// check if the file already exists
	if _, err := os.Stat(fileName); err == nil {
		colorize.PrintWithColor("The file already exists\n", colorize.Red)
		os.Exit(1)
	}


	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	// get the last string of the path
	arr := strings.Split(wd, h.GetSeperator())
	folderName := arr[len(arr)-1]

	h.CreateFile(fileName, gen.GenerateDB(folderName))

	h.ExecuteCommand("go", "mod","tidy")


	colorize.PrintWithColor("\n"+fileName+" file created\n", colorize.Green)
}
