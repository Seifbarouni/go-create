package createService

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	colorize "github.com/Seifbarouni/go-create/internal/colorizeText"
	h "github.com/Seifbarouni/go-create/internal/helpers"
)

func CreateService(model string) {
	model = strings.ReplaceAll(model, ".", "")
	
	// validate the model name with regex
	re := regexp.MustCompile(`[^-(),:;!?&*^\s]`)
	
	if !re.MatchString(model) {
		colorize.PrintWithColor(fmt.Sprintf("Invalid model name %s", model), colorize.Red)
		os.Exit(1)
	}
	

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
	module:=h.GetModuleName()
	// go to the services folder
	err=os.Chdir("services")
	if err!=nil{
		panic(err)
	}
	h.CreateFile(fileName, "//GENERATED BY go-create\npackage "+folderName+"\n\nimport (\n\tm \""+module+"/models\"\n\td \""+module+"/database\"\n)\n\n func Get"+model+"s()(*[]m."+model+",error){\n\t var "+model+"s [] m."+model+"\n\tif result:=d.DB.Find(&"+model+"s); result.Error!=nil{\n\t\treturn nil,result.Error\n\t}\n\treturn &"+model+"s,nil\n}\n\nfunc Get"+model+"(id int)(*m."+model+",error){\n\tvar "+model+" m."+model+"\n\tif result:=d.DB.First(&"+model+",id); result.Error!=nil{\n\t\treturn nil,result.Error\n\t}\n\treturn &"+model+",nil\n}\n\nfunc Create"+model+"("+model+" *m."+model+")error{\n\tif result:=d.DB.Create("+model+"); result.Error!=nil{\n\t\treturn result.Error\n\t}\n\treturn nil\n}\n\nfunc Update"+model+"("+model+" *m."+model+")error{\n\tif result:=d.DB.Save("+model+"); result.Error!=nil{\n\t\treturn result.Error\n\t}\n\treturn nil\n}\n\nfunc Delete"+model+"(id int)error{\n\tif result:=d.DB.Delete(&m."+model+"{ID:id}); result.Error!=nil{\n\t\treturn result.Error\n\t}\n\treturn nil\n}")

	colorize.PrintWithColor("\n"+fileName+" file created\n", colorize.Green)
}
