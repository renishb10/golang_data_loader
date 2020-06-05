package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"

	"github.com/google/uuid"
)

func generateData() <-chan uuid.UUID {
	c := make(chan uuid.UUID)
	const FILE_GUUID_PATH = "guid.txt"

	go func() {
		file, err := os.Open(FILE_GUUID_PATH)
		if err != nil {
			log.Fatal("Error reading GUUID file")
		}

		defer file.Close()

		reader := bufio.NewReader(file)
		for {
			line, err := reader.ReadString('\n')
			if err == io.EOF {
				break
			}

			line = strings.TrimSuffix(line, "\n")
			guid, err := uuid.Parse(line)

			if err != nil {
				continue
			}

			c <- guid
		}

		close(c)
	}()

	return c
}
