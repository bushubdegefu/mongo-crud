package generate

import (
	"github.com/bushubdegefu/mongo-crud/mtemplates"
)

func GitFrame(data mtemplates.Data) {
	tmpl := mtemplates.LoadTemplate("gitignore")

	mtemplates.WriteTemplateToFile(".gitignore", tmpl, data)
}

func DockerFrame(data mtemplates.Data) {
	tmpl := mtemplates.LoadTemplate("docker")

	mtemplates.WriteTemplateToFile("Dockerfile", tmpl, data)
}

func HaproxyFrame(data mtemplates.Data) {
	tmpl := mtemplates.LoadTemplate("haproxy.tmpl")
	mtemplates.WriteTemplateToFile("haproxy.cfg", tmpl, data)
}

func AppServiceFrame(data mtemplates.Data) {
	tmpl := mtemplates.LoadTemplate("service.tmpl")
	mtemplates.WriteTemplateToFile("app.service", tmpl, data)
}
