package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

// Program prints only unique lines separated by newline
func main() {
	lines := make(map[string]int)
	if len(os.Args) <= 1 {
		panic(errors.New("Please specify a file for input"))
	}
	fileName := os.Args[1]
	file, err := os.Open(fileName)
	if err != nil {
		panic(err.Error())
	}
	defer file.Close()
	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		lines[fileScanner.Text()]++
	}
	// This will print lines appearing more then once
	for line, count := range lines {
		if count > 1 {
			fmt.Println(line)
		}
	}
}
