package generate

import "github.com/bushubdegefu/mongo-crud/mtemplates"

func GenerateEchoSetup(data mtemplates.Data) {
	tmpl := mtemplates.LoadTemplate("echoSetup")

	mtemplates.WriteTemplateToFile("setup.go", tmpl, data)
}

func GenerateEchoAppMiddleware(data mtemplates.Data) {
	tmpl := mtemplates.LoadTemplate("echoAppMiddleware")

	mtemplates.WriteTemplateToFile("middleware.go", tmpl, data)
}

func GenerateGlobalEchoAppMiddleware(data mtemplates.Data) {
	tmpl := mtemplates.LoadTemplate("globalEchoMiddleware")

	mtemplates.WriteTemplateToFile("manager/middleware.go", tmpl, data)
}

func GenerateAppEchoGlobal(data mtemplates.Data) {
	tmpl := mtemplates.LoadTemplate("globalEchoApp")

	data.SetBackTick()
	mtemplates.WriteTemplateToFile("manager/app.go", tmpl, data)
}
