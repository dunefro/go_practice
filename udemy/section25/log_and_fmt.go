package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// Create a file for logging
	logFile, err := os.Create("log.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer logFile.Close()
	log.SetOutput(logFile)

	// Opening a file which doesn't exists to log error
	noFile, err := os.Open("no-file.txt")
	if err != nil {
		log.Println(err)
	}
	defer noFile.Close()
	fmt.Println("To Check the logs, open log.txt")
}
