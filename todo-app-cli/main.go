/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"todo-app-cli/app"
	"todo-app-cli/cmd"
)

func main() {

	app.Container.Bind(app.NewTodoRepositoryInterface)
	app.Container.Bind(app.NewCreateTodoUseCase)
	app.Container.Bind(app.NewUpdateTodoUseCase)
	app.Container.Bind(app.NewListTodosUseCase)
	app.Container.Bind(app.NewDeleteTodoUseCase)
	app.Container.Bind(app.NewConsoleRenderer)

	cmd.Execute()
}
