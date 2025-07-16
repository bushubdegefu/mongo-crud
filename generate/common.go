package generate

import (
	"os"

	"github.com/bushubdegefu/mongo-crud/mtemplates"
)

func GenerateCommon(data mtemplates.Data) {
	tmpl := mtemplates.LoadTemplate("common")
	err := os.MkdirAll("common", os.ModePerm)
	if err != nil {
		panic(err)
	}
	data.SetBackTick()
	mtemplates.WriteTemplateToFile("common/common.go", tmpl, data)
}
