package createservice

import (
	"os"
	"strings"

	colorize "github.com/Seifbarouni/go-create/internal/colorizeText"
	gen "github.com/Seifbarouni/go-create/internal/generators"
	"github.com/Seifbarouni/go-create/internal/helpers"
)

var h *helpers.Helpers = helpers.InitializeHelpers()

// CreateService creates a service file
func CreateService(model string) {
	// delete every special character from the file name
	model = strings.ReplaceAll(model, "-", "")
	model = strings.ReplaceAll(model, ".", "")
	model = strings.ReplaceAll(model, "(", "")
	model = strings.ReplaceAll(model, ")", "")
	model = strings.ReplaceAll(model, ",", "")
	model = strings.ReplaceAll(model, ":", "")
	model = strings.ReplaceAll(model, ";", "")
	model = strings.ReplaceAll(model, "!", "")
	model = strings.ReplaceAll(model, "?", "")
	model = strings.ReplaceAll(model, "&", "")
	model = strings.ReplaceAll(model, "*", "")
	model = strings.ReplaceAll(model, "^", "")
	model = strings.Title(model)
	// add Service.go to the file name
	fileName := ""
	if strings.HasSuffix(model, "s") {
		fileName = strings.ToLower(model) + "Service.go"
	} else {
		fileName = strings.ToLower(model) + "sService.go"
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
	module := h.GetModuleName()
	err = os.Chdir(wd)
	if err != nil {
		panic(err)
	}
	h.CreateFile(fileName, gen.GenerateService(folderName, module, model))

	colorize.PrintWithColor("\n"+fileName+" file created\n", colorize.Green)
}
