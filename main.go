package main

import (
	"os"
)

func main() {

	var commandLineArguments []string = os.Args[1:]
	var app Application
	app.initialize(commandLineArguments, "tasks.json")
	app.run()

}
