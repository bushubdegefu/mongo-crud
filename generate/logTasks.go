package generate

import (
	"os"

	"github.com/bushubdegefu/mongo-crud/mtemplates"
)

func GenerateTasks(data mtemplates.Data) {
	tmpl := mtemplates.LoadTemplate("tasks")
	err := os.MkdirAll("bluetasks", os.ModePerm)
	if err != nil {
		panic(err)
	}
	mtemplates.WriteTemplateToFile("bluetasks/tasks.go", tmpl, data)
}

func GenerateLogs(data mtemplates.Data) {
	tmpl := mtemplates.LoadTemplate("logs")
	err := os.MkdirAll("logs", os.ModePerm)
	if err != nil {
		panic(err)
	}
	mtemplates.WriteTemplateToFile("logs/logfile.go", tmpl, data)
}
