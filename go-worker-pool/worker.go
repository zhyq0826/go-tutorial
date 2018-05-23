package main

import (
	"github.com/satori/go.uuid"
)

// ID for job id
type ID struct {
	uuid.UUID
}

// Worker for job
type Worker struct {
	ID                ID
	ProcessingChannel chan<- bool
	FinishedChannel   chan<- ID
}

func NewWorker(ProcessingChannel chan<- bool, FinishedChannel chan<- ID) worker {
	return Worker{
		ID:                generateID(),
		ProcessingChannel: ProcessingChannel,
		FinishedChannel:   FinishedChannel,
	}
}

func generateID() ID {
	return ID{UUID: uuid.NewV4()}
}
