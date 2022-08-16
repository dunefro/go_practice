package main

import (
	"fmt"
	"os"
)

func main() {
	var rmdirs []func()
	for _, d := range tempDirs() {
		os.MkdirAll(d, 0755)
		rmdirs = append(rmdirs, func() {
			fmt.Println(d)
			os.RemoveAll(d)
		})
	}
	for _, f := range rmdirs {
		f()
	}
}
func tempDirs() []string {
	return []string{"/tmp/a", "/tmp/b", "/tmp/c", "/tmp/d"}
}
