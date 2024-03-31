package utils

import (
	"log"
	"os"
)

func ReadFile(filePath string) *os.File {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalln(err.Error())
	}
	return file
}
