package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "Progen",
	Short: "Progen is a simple tool to generate project structure",
	Long:  "Progen scaffolds new project using a specific framework like React, Sveltekit, Nextjs or GoLang with predefined convection and tools instead of creating directories from scratch every time. Progen automates this process for you.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Progen is a simple tool to scaffold projects use `progen --help` for more information.")
	},
}

var createProjectCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new project",
	Long:  "Create a new project using a specific framework like React, Sveltekit, Nextjs or GoLang with predefined convection and tools instead of creating directories from scratch every time. Progen automates this process for you.",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		projectName := args[0]
		projectType := args[1]
		createProject(projectName, projectType)
	},
}

func createProject(projectName, projectType string) {
	templateDir := fmt.Sprintf("templates/%s", projectType)
	destinationDir := projectName
	fmt.Printf("Creating project %s of from %s\n", destinationDir, templateDir)

	if err := os.Mkdir(destinationDir, 0755); err != nil {
		fmt.Printf("Error creating project %v", err)
		return
	}

	err := filepath.Walk(templateDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		relPath, err := filepath.Rel(templateDir, path)
		if err != nil {
			return err
		}

		destPath := filepath.Join(destinationDir, relPath)
		if info.IsDir() {
			// create directory if it does not exist
			if err := os.MkdirAll(destPath, info.Mode()); err != nil {
				return err
			}
		} else {
			// copy file content
			content, err := os.ReadFile(path)
			if err != nil {
				return err
			}

			if err := os.WriteFile(destPath, content, info.Mode()); err != nil {
				return err
			}
		}
		return nil
	})

	if err != nil {
		fmt.Printf("Error creating template files %v", err)
		return
	}

	fmt.Printf("Project %s created successfully form %s", projectName, templateDir)
}

func init() {
	rootCmd.AddCommand(createProjectCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
