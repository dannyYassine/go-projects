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
		fmt.Println("create called")
		name, _ := cmd.Flags().GetString("name")

		if name == "" {
			fmt.Println("name is required")
			return
		}

		dto := app.NewCreateTodoDto(name)
		todo, err := app.NewCreateTodoUseCase().Execute(dto)

		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(todo)
	},
}

func init() {
	createCmd.Flags().String("name", "", "name of todo")
	rootCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
