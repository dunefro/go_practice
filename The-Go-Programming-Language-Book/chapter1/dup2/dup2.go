package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

// This program prints the name of the file having duplicate lines
func main() {
	if len(os.Args) <= 1 {
		panic(errors.New("Please specify atleast one file"))
	}
	for _, fileName := range os.Args[1:] {
		file, err := os.Open(fileName)
		if err != nil {
			panic(err)
		}
		defer file.Close()
		if checkDuplicateLines(file) {
			fmt.Println(fileName)
		}
	}
}

func checkDuplicateLines(f *os.File) bool {
	lines := make(map[string]int)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines[scanner.Text()]++
	}
	for _, v := range lines {
		if v > 1 {
			return true
		}
	}
	return false
}
