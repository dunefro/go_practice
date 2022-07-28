package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetchURL(url, ch)
		// fmt.Println(<-ch) // uncomment this to and comment below loop to understand the different in time
	}
	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
	fmt.Printf("Total time: %.2fs\n", time.Since(start).Seconds())
}

func fetchURL(url string, ch chan<- string) {
	resp, err := http.Get(url)
	time.Sleep(3 * time.Second)
	if err != nil {
		panic(err)
	}
	ch <- fmt.Sprintf("URL %s --> %d", url, resp.StatusCode)
}
