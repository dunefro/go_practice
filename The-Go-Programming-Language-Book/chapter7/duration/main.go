package main

import (
	"flag"
	"fmt"
	"time"
)

// cmd flag that accepts time
// default time is 1 second
var period = flag.Duration("period", 1*time.Second, "Sleeping period")

func main() {
	flag.Parse()
	fmt.Printf("Sleeping for %v\n", *period)
	time.Sleep(*period)
}

// Usage: go run main.go -period 2m30s
// Will sleep for 2 minutes 30 seconds
