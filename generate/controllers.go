package generate

import (
	"fmt"
	"os"
	"strings"

	"github.com/bushubdegefu/mongo-crud/mtemplates"
)

func GenerateControllers(data mtemplates.Data) {
	tmpl := mtemplates.LoadTemplate("controllers")

	_ = os.MkdirAll("controllers", os.ModePerm)

	for _, model := range data.Models {
		filePath := fmt.Sprintf("controllers/%s_controllers.go", strings.ToLower(model.Name))
		mtemplates.WriteTemplateToFileModel(filePath, tmpl, model)
	}

	// mtemplates.WriteTemplateToFile("models/init.go", migrationTmpl, data)

}
