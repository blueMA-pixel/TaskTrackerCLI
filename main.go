package main

import (
	"fmt"
	"os"
)

func main() {

	var commandLineArguments []string = os.Args[1:]
	var app Application
	var loader TaskDataHandlerJSON = TaskDataHandlerJSON{fileName: "tasks.json"}
	err := app.initialize(commandLineArguments, os.Stdout, loader)

	if err != nil {
		fmt.Fprintln(app.writer, err)
		return
	}

	app.run()
}
