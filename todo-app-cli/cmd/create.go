package cmd

import (
	"fmt"
	"todo-app-cli/app"

	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		description, _ := cmd.Flags().GetString("description")

		if name == "" {
			fmt.Println("name is required")
			return
		}

		dto := app.NewCreateTodoDto(name, description)

		useCase := app.Get[app.CreateTodoUseCase]()

		todo, err := useCase.Execute(dto)

		if err != nil {
			fmt.Println(err)
			return
		}

		renderer := app.Get[app.ConsoleRenderer]()
		renderer.PrintTodo(todo)
	},
}

func init() {
	createCmd.Flags().String("name", "", "name of todo")
	createCmd.Flags().String("description", "", "small description (optional)")
	rootCmd.AddCommand(createCmd)
}
