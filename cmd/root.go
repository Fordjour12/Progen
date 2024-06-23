package cmd

import (
	"fmt"
	"os"

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

func createProject(projectName string, projectType string) {
	fmt.Printf("Creating project %s of type %s\n", projectName, projectType)
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
