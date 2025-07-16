package generate

import (
	"os"

	"github.com/bushubdegefu/mongo-crud/mtemplates"
)

func GenerateTracerEchoSetup(data mtemplates.Data) {
	tmpl := mtemplates.LoadTemplate("echoObserve")
	promTmpl := mtemplates.LoadTemplate("prometheus")
	tmplMetric := mtemplates.LoadTemplate("promyml")
	err := os.MkdirAll("observe", os.ModePerm)
	if err != nil {
		panic(err)
	}
	mtemplates.WriteTemplateToFile("observe/tracer.go", tmpl, data)
	mtemplates.WriteTemplateToFile("observe/prometheus.go", promTmpl, data)
	mtemplates.WriteTemplateToFile("prometheus.yml", tmplMetric, data)
}
