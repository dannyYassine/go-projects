/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"todo-app-cli/app"
	"todo-app-cli/cmd"
)

func main() {
	app.Application.Register().Boot()

	cmd.Execute()
}
