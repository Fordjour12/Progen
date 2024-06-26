package cmd

import (
	"github.com/Fordjour12/progen/cmd/ui/multichoice"
	"github.com/Fordjour12/progen/cmd/ui/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(createCmd)
}

type ProjectList struct {
	projectOptions []string
}

type CreateOptions struct {
	ProjectType *multichoice.Selection
	ProjectName *textinput.Output
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new project",
	Long: `Create a new project with a specific tools like
 React, Sveltekit, Nextjs or GoLang with predefined convection`,
	Run: func(cmd *cobra.Command, args []string) {
		options := CreateOptions{
			ProjectName: &textinput.Output{},
		}

		pjTypeList := ProjectList{
			projectOptions: []string{
				"React",
				"React Native",
				"Sveltekit",
				"Nextjs",
				"GoLang",
			},
		}

		tprogarm := tea.NewProgram(textinput.InitialModel("What is your Project Name: ?", options.ProjectName))
		if _, err := tprogarm.Run(); err != nil {
			cobra.CheckErr(err)
		}

		tprogarm = tea.NewProgram(multichoice.InitialChoiceModel("What is your Project Type: ", pjTypeList.projectOptions, options.ProjectType))
		if _, err := tprogarm.Run(); err != nil {
			cobra.CheckErr(err)
		}
	},
}
