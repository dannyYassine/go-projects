/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"todo-app-cli/app"

	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		id, _ := cmd.Flags().GetString("id")
		name, _ := cmd.Flags().GetString("name")
		description, _ := cmd.Flags().GetString("description")
		status, _ := cmd.Flags().GetString("status")

		if id == "" {
			fmt.Println("id is required")
			return
		}

		dto := app.NewUpdateTodoDto(id, name, description, status)

		useCase := app.Get[app.UpdateTodoUseCase]()

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
	updateCmd.Flags().String("id", "", "id of the todo")
	updateCmd.Flags().String("name", "", "name of todo")
	updateCmd.Flags().String("description", "", "small description (optional)")
	updateCmd.Flags().String("status", "", "status (optional)")
	rootCmd.AddCommand(updateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// updateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// updateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
