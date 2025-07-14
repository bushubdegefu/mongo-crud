package generate

import (
	"os"

	"github.com/bushubdegefu/mongo-crud/mtemplates"
)

func GenerateMainAndManager(data mtemplates.Data) {
	tmplMain := mtemplates.LoadTemplate("main")
	tmplManager := mtemplates.LoadTemplate("manager")
	err := os.MkdirAll("manager", os.ModePerm)
	if err != nil {
		panic(err)
	}
	mtemplates.WriteTemplateToFile("main.go", tmplMain, data)
	mtemplates.WriteTemplateToFile("manager/manager.go", tmplManager, data)
}
