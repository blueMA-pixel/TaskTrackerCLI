package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

const commandIndex int = 0

const listCommand string = "list"
const addCommand string = "add"
const deleteCommand string = "delete"
const updateCommand string = "update"
const markasDoneCommand string = "mark-as-done"
const markinProgressCommand string = "mark-in-progress"

type Application struct {
	tasks    Tasks
	command  ICommand
	fileName string
}

func (app *Application) initialize(commandLineArguments []string, fileName string) error {

	if len(fileName) == 0 {
		return nil
	}

	app.fileName = fileName

	switch commandLineArguments[commandIndex] {
	case listCommand:
		app.command = &ListCommand{}
	case addCommand:
		app.command = &AddCommand{}
	case deleteCommand:
		app.command = &DeleteCommand{}
	case updateCommand:
		app.command = &UpdateCommand{}
	case markasDoneCommand, markinProgressCommand:
		app.command = &MarkAsCommand{}
	}

	return app.command.initialize(commandLineArguments)
}

func (handler *Application) readTasks() error {
	jsonFile, err := os.OpenFile(handler.fileName, os.O_RDWR|os.O_CREATE, 0644)

	if err != nil {
		panic(err)
	}

	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)

	if err != nil {
		fmt.Println("Error unmarshaling:", err)
		return err
	}

	return json.Unmarshal(byteValue, &handler.tasks)
}

func (handler *Application) writeTasks() error {

	jsonFile, err := os.OpenFile(handler.fileName, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)

	if err != nil {
		panic(err)
	}

	defer jsonFile.Close()

	byteValue, err := json.Marshal(handler.tasks)

	if err != nil {
		fmt.Println("Error marshaling:", err)
		return err
	}

	_, werr := jsonFile.Write(byteValue)

	return werr
}

func (app *Application) run() error {
	app.readTasks()
	app.command.execute(&app.tasks)
	app.writeTasks()

	return nil
}
