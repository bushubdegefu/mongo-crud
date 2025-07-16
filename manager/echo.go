package manager

import (
	"fmt"
	"os/exec"

	"github.com/bushubdegefu/mongo-crud/generate"
	"github.com/bushubdegefu/mongo-crud/mtemplates"
	"github.com/spf13/cobra"
)

var (
	echocli = &cobra.Command{
		Use:   "echo",
		Short: "generate the basic structure file to start app using echo",
		Long:  `generate the basic structure file to start app using echo`,
		Run: func(cmd *cobra.Command, args []string) {

			// Initialize the project settings
			mtemplates.InitProjectJSON()
			mtemplates.RenderData.ProjectName = mtemplates.ProjectSettings.ProjectName
			mtemplates.RenderData.AppNames = mtemplates.ProjectSettings.AppNames
			mtemplates.RenderData.AuthAppType = mtemplates.ProjectSettings.AuthAppType

			appName, _ := cmd.Flags().GetString("app")
			globalName, _ := cmd.Flags().GetBool("global")

			if appName != "" {
				handleAppDirectory(appName)
				if err := mtemplates.LoadData(config_file); err != nil {
					fmt.Printf("Error loading data: %v\n", err)
					return
				}
				// generate.GenerateFiberAppMiddleware(mtemplates.RenderData)
				generate.GenerateEchoAppMiddleware(mtemplates.RenderData)
				generate.GenerateEchoSetup(mtemplates.RenderData)
				mtemplates.ProjectSettings.CurrentAppName = appName

			} else if globalName {
				generate.GenerateGlobalEchoAppMiddleware(mtemplates.RenderData)
				generate.GenerateAppEchoGlobal(mtemplates.RenderData)
				runSwagInitForApps()
			} else {
				fmt.Println("No app name specified")
			}
			mtemplates.CommonCMD()
		},
	}
)

func runSwagInitForApps() {
	mtemplates.InitProjectJSON()
	// swag init --generalInfo setup.go --output  blue-auth/docs --dir=blue-auth,common
	for _, appName := range mtemplates.ProjectSettings.AppNames {
		dirArg := fmt.Sprintf("%s,common", appName)
		outputDir := fmt.Sprintf("%s/docs", appName)

		// Prepare the swag init command
		cmd := exec.Command(
			"swag", "init",
			"--generalInfo", "setup.go",
			"--output", outputDir,
			"--dir", dirArg,
		)

		// Run the command and handle errors
		if err := cmd.Run(); err != nil {
			fmt.Printf("Error generating swagger for app '%s': %v\n", appName, err)
		} else {
			fmt.Printf("Swagger generated for app '%s'\n", appName)
		}
	}
}

func init() {
	echocli.Flags().StringP("app", "a", "", "Specify the app name, so that echo app will be generated")
	echocli.Flags().BoolP("global", "g", false, "basic echo app with for global, creates app.go( in manager package) and middleware.go on the main module takes true or false")
	echocli.Flags().StringVarP(&config_file, "config", "c", "config.json", "Specify the data file to load")
	goFrame.AddCommand(echocli)
}
