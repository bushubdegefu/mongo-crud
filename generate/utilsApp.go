package generate

import (
	"os"

	"github.com/bushubdegefu/mongo-crud/mtemplates"
)

func GenerateUtilsApp(data mtemplates.ProjectSetting) {
	tmpl := mtemplates.LoadTemplate("utilsApp")
	err := os.MkdirAll("utils", os.ModePerm)
	if err != nil {
		panic(err)
	}
	mtemplates.WriteTemplateToFileSetting("utils/jwt_utils.go", tmpl, data)
}
