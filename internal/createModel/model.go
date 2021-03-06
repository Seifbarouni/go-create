package createmodel

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	colorize "github.com/Seifbarouni/go-create/internal/colorizeText"
	gen "github.com/Seifbarouni/go-create/internal/generators"
	"github.com/Seifbarouni/go-create/internal/helpers"
)

var h *helpers.Helpers = helpers.InitializeHelpers()

func generateModelName(fileName string) string {
	// split the file name by " "
	arr := strings.Split(fileName, " ")
	// for each word in the array capitalize the first letter
	for i := 0; i < len(arr); i++ {
		arr[i] = strings.Title(arr[i])
	}
	// join the array to get the model name
	return strings.Join(arr, "")
}

func getModelAttributes() string {
	colorize.PrintWithColor("Enter the attributes of the model (ID, CreatedAt and UpdatedAt already added)\nSyntax : attribute-name:attributes-type ...\nExample : Num:int Price:float Name:string\n", colorize.Yellow)
	colorize.PrintWithColor("Examples of attribute types : int8, int16, int32,int64, float, float32, float64, string, time, bool ...\n", colorize.Purple)
	attributes := make(map[string]string)
	// get input from the user in a form of a string
	userInput := ""
	var structStr string
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("-> ")
	userInput, _ = reader.ReadString('\n')
	// split the string by " "
	arr := strings.Split(userInput, " ")
	// for each element in the array split it by ":"
	for i := 0; i < len(arr); i++ {
		items := strings.Split(arr[i], ":")
		// check if the user has provided a valid attribute name
		if len(items) != 2 {
			colorize.PrintWithColor("Invalid attribute\n", colorize.Red)
			os.Exit(1)
		}
		attr := strings.Title(items[0])
		tp := strings.ToLower(items[1])
		if tp == "time" {
			tp = "time.Time"
		}
		attributes[attr] = tp
	}

	// create the struct
	for k, v := range attributes {
		structStr += "\t" + k + "\t" + v + "\n"
	}
	return structStr
}

// CreateModel creates a new model
func CreateModel(fileName string) {
	// check if the user has provided a valid file name
	if fileName == "" {
		colorize.PrintWithColor("You must provide a valid file name\n", colorize.Red)
		os.Exit(1)
	}
	// capitalize the first letter of the file name
	modelName := generateModelName(fileName)
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
	// generate struct with the model name
	structStr := getModelAttributes()
	imports := "import \"time\"\n\n"

	h.CreateFile(fileName, gen.GenerateModel(folderName, imports, modelName, structStr))

	colorize.PrintWithColor("\n"+fileName+" file created\n", colorize.Green)

}
