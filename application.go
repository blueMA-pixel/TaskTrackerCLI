package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Application struct {
	tasks []Task
}

func (handler *Application) readTasks() {
	jsonFile, err := os.OpenFile("tasks.json", os.O_RDWR|os.O_CREATE, 0644)

	if err != nil {
		panic(err)
	}

	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)

	if err != nil {
		fmt.Println("Error unmarshaling:", err)
		panic(err)
	}

	json.Unmarshal(byteValue, &handler.tasks)

}

func (handler *Application) writeTasks() {

	jsonFile, err := os.OpenFile("tasks.json", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)

	if err != nil {
		panic(err)
	}

	defer jsonFile.Close()

	byteValue, err := json.Marshal(handler.tasks)

	if err != nil {
		fmt.Println("Error marshaling:", err)
		panic(err)
	}

	jsonFile.Write(byteValue)
}
