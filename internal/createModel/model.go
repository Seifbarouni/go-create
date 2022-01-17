package createModel

import (
	"os"
	"strings"

	colorize "github.com/Seifbarouni/go-create/internal/colorizeText"
)

func CreateModel(fileName string) {
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
}