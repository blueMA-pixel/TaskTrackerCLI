package main

import (
	"fmt"
	"os"
)

func main() {

	var commandLineArguments []string = os.Args[1:]
	var app Application
	err := app.initialize(commandLineArguments, "tasks.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	app.run()
}
