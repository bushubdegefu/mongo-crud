package generate

import (
	"fmt"
	"os"
	"strings"

	"github.com/bushubdegefu/mongo-crud/mtemplates"
)

func GenerateModels(data mtemplates.Data) {
	tmpl := mtemplates.LoadTemplate("model")
	// migrationTmpl := mtemplates.LoadTemplate("migration")
	helperTmpl := mtemplates.LoadTemplate("helperModel")

	_ = os.MkdirAll("models", os.ModePerm)

	for _, model := range data.Models {
		filePath := fmt.Sprintf("models/%s.go", strings.ToLower(model.Name))
		mtemplates.WriteTemplateToFileModel(filePath, tmpl, model)
	}

	// mtemplates.WriteTemplateToFile("models/init.go", migrationTmpl, data)
	mtemplates.WriteTemplateToFile("models/helper.go", helperTmpl, data)
}
