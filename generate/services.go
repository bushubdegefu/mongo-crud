package generate

import (
	"fmt"
	"os"
	"strings"

	"github.com/bushubdegefu/mongo-crud/mtemplates"
)

func GenerateServices(data mtemplates.Data) {
	tmpl := mtemplates.LoadTemplate("services")

	_ = os.MkdirAll("services", os.ModePerm)

	for _, model := range data.Models {
		filePath := fmt.Sprintf("services/%s_service.go", strings.ToLower(model.Name))
		mtemplates.WriteTemplateToFileModel(filePath, tmpl, model)
	}

	// mtemplates.WriteTemplateToFile("models/init.go", migrationTmpl, data)

}
