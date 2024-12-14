package main

import (
	"fmt"
	"os"
)

func main() {

	var commandLineArguments []string = os.Args[1:]
	var app Application

	app.readTasks()

	switch commandLineArguments[0] {
	case "List":
		var l ListCommand
		l.CheckCommandLineArguments(commandLineArguments)
		l.executeCommand(app.tasks)
	case "Add":
		var add AddCommand
		add.CheckCommandLineArguments(commandLineArguments)
		add.executeCommand(&app.tasks)
	case "Delete":
		var delete DeleteCommand
		delete.CheckCommandLineArguments(commandLineArguments)
		delete.executeCommand(&app.tasks)
	case "Update":
		var update UpdateCommand
		update.CheckCommandLineArguments(commandLineArguments)
		update.executeCommand(&app.tasks)
	case "mark-as-done", "mark-in-progress":
		var MarkAsCommand MarkAsCommand
		MarkAsCommand.CheckCommandLineArguments(commandLineArguments)
		MarkAsCommand.executeCommand(&app.tasks)
	default:
		fmt.Println("Operation not found")
	}

	app.writeTasks()

}
