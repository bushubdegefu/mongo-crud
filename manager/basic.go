package manager

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/bushubdegefu/mongo-crud/generate"
	"github.com/bushubdegefu/mongo-crud/mtemplates"
	"github.com/spf13/cobra"
)

var (
	config_file     string
	initalizemodule = &cobra.Command{
		Use:   "init",
		Short: "Initialize the module with name",
		Long:  `Provide name to initialize the project using the "name" flag.`,
		Run: func(cmd *cobra.Command, args []string) {
			moduleName, _ := cmd.Flags().GetString("name")
			appName, _ := cmd.Flags().GetString("app")
			authAppName, _ := cmd.Flags().GetString("auth")
			authAppType, _ := cmd.Flags().GetString("type")

			if appName == "" && moduleName == "" {

				fmt.Println("Please provide app name with app name flag or module name with  name flag")

			} else {
				// Initialize the module
				if moduleName != "" {
					mtemplates.CommonProjectName(moduleName, authAppName, authAppType)
					mtemplates.CommonModInit(moduleName)
					// mtemplates.CommonCMDInit()
				}

				// If no module name, fetch the project name
				if moduleName == "" {
					mtemplates.InitProjectJSON()
					mtemplates.RenderData.ProjectName = mtemplates.ProjectSettings.ProjectName
				} else {
					mtemplates.RenderData.ProjectName = moduleName
				}

				fmt.Println(moduleName)
				// Get current working directory
				currentDir, _ := os.Getwd()
				generate.GenerateMainAndManager(mtemplates.RenderData)
				generate.GenerateConfig(mtemplates.RenderData)

				// Handle appName if provided
				if appName != "" {
					handleAppInitialization(appName, currentDir, authAppName)
				}
			}
			mtemplates.CommonCMD()
		},
	}
	configcli = &cobra.Command{
		Use:   "config",
		Short: "Template Configuration Variables need for the apps registerd to run",
		Long:  `Template Configuration Variables need for the apps registerd to run.`,
		Run: func(cmd *cobra.Command, args []string) {
			mtemplates.InitProjectJSON()
			mtemplates.RenderData.ProjectName = mtemplates.ProjectSettings.ProjectName
			mtemplates.RenderData.AppNames = mtemplates.ProjectSettings.AppNames
			generate.GenerateConfig(mtemplates.RenderData)
			generate.GenerateConfigEnv(mtemplates.RenderData)
			generate.GenerateConfigAppEnv(mtemplates.RenderData)
		},
	}

	basicCommand = &cobra.Command{
		Use:   "basic",
		Short: "Generate a basic folder structure for a project",
		Long:  `This command generates a basic folder structure for a project. The type flag determines the specific setup.`,
		Run: func(cmd *cobra.Command, args []string) {
			projectType, _ := cmd.Flags().GetString("type")
			appName, _ := cmd.Flags().GetString("app")
			frame, _ := cmd.Flags().GetString("frame")

			mtemplates.InitProjectJSON()

			// Handle appName if provided
			if appName != "" {
				handleAppDirectory(appName)
				if err := mtemplates.LoadData(config_file); err != nil {
					fmt.Printf("Error loading data: %v\n", err)
					return
				}
				mtemplates.RenderData.ProjectName = mtemplates.ProjectSettings.ProjectName
				mtemplates.RenderData.AppName = appName
			}

			// Generate structure based on project type
			handleProjectType(projectType, frame, cmd)
		},
	}
)

func handleAppInitialization(appName, currentDir, authAppName string) {

	mtemplates.RenderData.AppName = appName
	mtemplates.RenderData.ProjectName = mtemplates.ProjectSettings.ProjectName
	mtemplates.ProjectSettings.AppendAppName(appName, authAppName)
	// Create app directory and switch to it
	os.Mkdir(appName, os.ModePerm)
	newDir := filepath.Join(currentDir, appName)
	_ = os.Chdir(newDir)
	if mtemplates.ProjectSettings.AuthAppType == "sso" {
		// Generate the SSO schema app
		generate.GenerateSSOAuth(mtemplates.RenderData)
	} else {
		// Generate the Django auth Schema app
		generate.GenerateDjangoAuth(mtemplates.RenderData)
	}
}

func handleAppDirectory(appName string) {
	currentDir, _ := os.Getwd()
	newDir := filepath.Join(currentDir, appName)
	_ = os.Chdir(newDir)
}

func handleProjectType(projectType, frame string, cmd *cobra.Command) {
	switch projectType {
	case "service":
		basiccmd()
		mtemplates.CommonCMD()
	case "db":
		mtemplates.InitProjectJSON()
		generate.GenerateDBConn(mtemplates.ProjectSettings)
		generate.GenerateCommon(mtemplates.RenderData)
	case "config":
		mtemplates.InitProjectJSON()
		mtemplates.RenderData.AppNames = mtemplates.ProjectSettings.AppNames
		generate.GenerateConfig(mtemplates.RenderData)
		generate.GenerateConfigEnv(mtemplates.RenderData)
		generate.GenerateConfigAppEnv(mtemplates.RenderData)
	case "tracer":
		mtemplates.InitProjectJSON()
		mtemplates.RenderData.ProjectName = mtemplates.ProjectSettings.ProjectName
		generate.GenerateTracerEchoSetup(mtemplates.RenderData)
	case "logs":
		mtemplates.InitProjectJSON()
		mtemplates.RenderData.ProjectName = mtemplates.ProjectSettings.ProjectName
		generate.GenerateLogs(mtemplates.RenderData)
	case "tasks":
		appName, _ := cmd.Flags().GetString("app")
		if appName == "" {
			fmt.Println("tasks flag need additional flag app")
		} else {
			generate.GenerateTasks(mtemplates.RenderData)
			mtemplates.CommonCMD()
		}
	default:
		fmt.Println(frame)
		// fmt.Printf("Args: %#v\n", args)
		appName, _ := cmd.Flags().GetString("app")
		if appName == "" {
			fmt.Println("No app name provided")
		}
		fmt.Println("Unknown type specified. Valid types are: service( gives you docker, gitignore, linux service and haproxy.cfg).")
	}
}

func basiccmd() {

	generate.GitFrame(mtemplates.RenderData)
	generate.HaproxyFrame(mtemplates.RenderData)
	generate.DockerFrame(mtemplates.RenderData)
}

func init() {
	// Register flags for all commands
	initalizemodule.Flags().StringP("name", "n", "", "Specify the module name  (github.com/someuser/someproject)")
	initalizemodule.Flags().StringP("app", "a", "", "Specify the application name  like auth-app,hrm-app")
	initalizemodule.Flags().StringP("auth", "p", "", "Specify the authentication application name  defaults to django_auth")
	initalizemodule.Flags().StringP("type", "t", "", "specify if you are using standalone authentication like django admin or sso like solution")

	// Register flags for the 'basic' command
	basicCommand.Flags().StringP("type", "t", "", "Specify the type of folder structure to generate: rsa, db, producer,logs, consumer, tasks, pagination, otel,migration,config")
	basicCommand.Flags().StringP("frame", "f", "", "Specify the Spanner function you want for the tracer, echo/fiber, meant to be used with otel flag")
	basicCommand.Flags().StringP("name", "n", "", "Specify the project module name as in github.com/someuser/someproject for the json template generation")
	basicCommand.Flags().StringP("app", "a", "", "Specify the app name, all it will try to generate for all jsons")

	goFrame.AddCommand(basicCommand)
	goFrame.AddCommand(configcli)
	goFrame.AddCommand(initalizemodule)
}
