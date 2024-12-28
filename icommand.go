package main

import "io"

type ICommand interface {
	initialize(commandLineArguments []string, writer io.Writer) error
	execute(tasks *Tasks) error
}
