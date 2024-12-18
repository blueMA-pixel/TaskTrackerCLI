package main

type ICommand interface {
	initialize(commandLineArguments []string) error
	execute(tasks *Tasks) error
}
