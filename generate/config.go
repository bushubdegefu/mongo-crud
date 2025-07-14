package generate

import (
	"os"

	"github.com/bushubdegefu/mongo-crud/mtemplates"
)

func GenerateConfig(data mtemplates.Data) {
	tmpl := mtemplates.LoadTemplate("config")
	err := os.MkdirAll("configs", os.ModePerm)
	if err != nil {
		panic(err)
	}
	mtemplates.WriteTemplateToFile("configs/configs.go", tmpl, data)
}

func GenerateConfigEnv(data mtemplates.Data) {
	tmpl := mtemplates.LoadTemplate("env")

	mtemplates.WriteTemplateToFile("configs/.env", tmpl, data)
}

func GenerateConfigAppEnv(data mtemplates.Data) {
	tmpl := mtemplates.LoadTemplate("projectEnv")

	mtemplates.WriteTemplateToFile("configs/.dev.env", tmpl, data)
	mtemplates.WriteTemplateToFile("configs/.prod.env", tmpl, data)
}

func GenerateConfigTestEnv(data mtemplates.Data) {
	tmpl := mtemplates.LoadTemplate("testEnv")
	err := os.MkdirAll("tests", os.ModePerm)
	if err != nil {
		panic(err)
	}
	mtemplates.WriteTemplateToFile("tests/.test.env", tmpl, data)
}
