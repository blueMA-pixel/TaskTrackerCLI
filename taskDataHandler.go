package main

import (
	"encoding/json"
	"io"
	"os"
)

type ITaskDataHandler interface {
	readTasks(tasks *Tasks) error
	writeTasks(tasks *Tasks) error
}

type TaskDataHandlerJSON struct {
	fileName string
}

func (handler TaskDataHandlerJSON) readTasks(tasks *Tasks) error {
	jsonFile, err := os.OpenFile(handler.fileName, os.O_RDWR|os.O_CREATE, 0644)

	if err != nil {
		return err
	}

	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)

	if err != nil {
		return err
	}

	return json.Unmarshal(byteValue, tasks)
}

func (handler TaskDataHandlerJSON) writeTasks(tasks *Tasks) error {
	jsonFile, err := os.OpenFile(handler.fileName, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)

	if err != nil {
		return err
	}

	defer jsonFile.Close()

	byteValue, err := json.Marshal(*tasks)

	if err != nil {
		return err
	}

	_, writeErr := jsonFile.Write(byteValue)

	return writeErr
}
