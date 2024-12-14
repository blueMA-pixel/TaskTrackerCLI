package main

import (
	"time"
)

type Status int

const (
	ToDo Status = iota
	InProgress
	Done
)

type Task struct {
	Id           int       `json:"ID"`
	Description  string    `json:"Description"`
	Status       Status    `json:"Status"`
	CreationTime time.Time `json:"CreationTime"`
	UpdateTime   time.Time `json:"UpdateTime"`
}
