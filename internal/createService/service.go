package createService

import (
	"os"
	"strings"

	colorize "github.com/Seifbarouni/go-create/internal/colorizeText"
	h "github.com/Seifbarouni/go-create/internal/helpers"
)

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
	// go to the services folder
	err = os.Chdir("services")
	if err != nil {
		panic(err)
	}
	h.CreateFile(fileName, "//GENERATED BY go-create\npackage "+folderName+"\n\nimport (\n\tm \""+module+"/models\"\n\td \""+module+"/database\"\n)\n\ntype "+model+"sService struct {}\n\ntype "+model+"sServiceInterface interface {\n\tGet"+model+"s()(*[]m."+model+",error)\n\tGet"+model+"(id int)(*m."+model+",error)\n\tCreate"+model+"("+model+" *m."+model+")error\n\tUpdate"+model+"("+model+" *m."+model+")error\n\tDelete"+model+"(id int)error\n}\nfunc Create"+model+"sService() *"+model+"sService {\n\t return &"+model+"sService{}\n}\nfunc (*"+model+"sService)Get"+model+"s()(*[]m."+model+",error){\n\t var "+model+"s [] m."+model+"\n\tif result:=d.DB.Find(&"+model+"s); result.Error!=nil{\n\t\treturn nil,result.Error\n\t}\n\treturn &"+model+"s,nil\n}\n\nfunc (*"+model+"sService)Get"+model+"(id int)(*m."+model+",error){\n\tvar "+model+" m."+model+"\n\tif result:=d.DB.First(&"+model+",id); result.Error!=nil{\n\t\treturn nil,result.Error\n\t}\n\treturn &"+model+",nil\n}\n\nfunc (*"+model+"sService)Create"+model+"("+model+" *m."+model+")error{\n\tif result:=d.DB.Create("+model+"); result.Error!=nil{\n\t\treturn result.Error\n\t}\n\treturn nil\n}\n\nfunc (*"+model+"sService)Update"+model+"("+model+" *m."+model+")error{\n\tif result:=d.DB.Save("+model+"); result.Error!=nil{\n\t\treturn result.Error\n\t}\n\treturn nil\n}\n\nfunc (*"+model+"sService)Delete"+model+"(id int)error{\n\tif result:=d.DB.Delete(&m."+model+"{ID:id}); result.Error!=nil{\n\t\treturn result.Error\n\t}\n\treturn nil\n}")

	colorize.PrintWithColor("\n"+fileName+" file created\n", colorize.Green)
}
