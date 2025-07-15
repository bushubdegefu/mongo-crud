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

	// Models command for generating data models
	modelscli = &cobra.Command{
		Use:   "models",
		Short: "Generate data models based on GORM with annotations",
		Long:  `This command generates data models using GORM, based on the provided spec in the config.json file, along with GORM relationship annotations.`,
		Run:   runModelsCommand,
	}
)

// runModelsCommand handles the execution of the 'models' command
func runModelsCommand(cmd *cobra.Command, args []string) {

	modelsType, _ := cmd.Flags().GetString("type")
	appName, _ := cmd.Flags().GetString("app")

	if appName == "" {
		fmt.Println("Error: --app flag is required.")
		return
	}

	// Change to the app's directory and load the config data
	if err := handleAppDirectoryAndLoadConfig(appName); err != nil {
		fmt.Println(err)
		return
	}

	// Generate models and migrations
	if modelsType == "init" {
		generate.GenerateModels(mtemplates.RenderData)

	} else {
		generate.GenerateModels(mtemplates.RenderData)
		mtemplates.CommonCMD()
	}
}

// handleAppDirectoryAndLoadConfig changes the working directory to the app's directory and loads the config data
func handleAppDirectoryAndLoadConfig(appName string) error {
	mtemplates.InitProjectJSON()
	currentDir, _ := os.Getwd()
	newDir := filepath.Join(currentDir, appName)
	if err := os.Chdir(newDir); err != nil {
		fmt.Println("Errorr Changing directory")
		return fmt.Errorf("error changing directory: %v", err)
	}
	mtemplates.RenderData.AppName = appName
	if err := mtemplates.LoadData(config_file); err != nil {
		return fmt.Errorf("error loading data: %v", err)
	}
	return nil
}

func init() {
	// Register flags for CRUD and Models commands
	modelscli.Flags().StringVarP(&config_file, "config", "c", "config.json", "Specify the data file to load")
	modelscli.Flags().StringP("type", "t", "", "Rerender the migration function by setting type to \"init\"")
	modelscli.Flags().StringP("app", "a", "", "Set app name, e.g., \"blue-auth\"")
	modelscli.Flags().BoolP("auth", "p", false, "Tell if generating models for auth app true or false")

	// Register commands to the root (goFrame)

	goFrame.AddCommand(modelscli)
}
