package generate

import (
	"os"

	"github.com/bushubdegefu/mongo-crud/mtemplates"
)

func GenerateDBConn(data mtemplates.ProjectSetting) {
	tmpl := mtemplates.LoadTemplate("database")
	err := os.MkdirAll("database", os.ModePerm)
	if err != nil {
		panic(err)
	}
	mtemplates.WriteTemplateToFileSetting("database/database.go", tmpl, data)
}
