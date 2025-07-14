package generate

import "github.com/bushubdegefu/mongo-crud/mtemplates"

func GenerateDjangoAuth(data mtemplates.Data) {
	tmpl := mtemplates.LoadTemplate("django")

	mtemplates.WriteTemplateToFile("config.json", tmpl, data)
}

func GenerateSSOAuth(data mtemplates.Data) {
	tmpl := mtemplates.LoadTemplate("sso")

	mtemplates.WriteTemplateToFile("config.json", tmpl, data)
}
