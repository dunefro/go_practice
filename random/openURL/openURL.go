package main

import (
	"fmt"
	"io/ioutil"

	"github.com/pkg/browser"
)

func main() {
	// already present in your code
	browser.Stderr = ioutil.Discard
	browser.Stdout = ioutil.Discard

	// endpoint /cli-auth
	url := "https://coder.example.com/cli-auth"

	// opening browser in another goroutine
	go browser.OpenURL(url)

	// get your token, below steps will be part of the parent goroutine
	var token string
	fmt.Print("Enter your token here: ")
	fmt.Scanf("%s", &token)
	fmt.Println("Your token is:", token)
}
