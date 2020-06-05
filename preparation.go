package main

import (
	"time"

	"github.com/google/uuid"
)

type inputData struct {
	id        string
	timestamp int64
}

func prepareData(ic <-chan uuid.UUID) <-chan inputData {
	oc := make(chan inputData)

	go func() {
		for id := range ic {
			input := inputData{
				id:        id.String(),
				timestamp: time.Now().UnixNano(),
			}
			oc <- input
		}
		close(oc)
	}()

	return oc
}
