package sel

import (
	"fmt"
	"net/http"
	"time"
)

func Racer(url1 string, url2 string) (winner string) {

	time1 := measureDuration(url1)
	time2 := measureDuration(url2)

	fmt.Printf("%q: %d\n%q: %d\n", url1, time1, url2, time2)

	if time1 > time2 {
		return url2
	}

	return url1
}
func measureDuration(url string) time.Duration {

	start := time.Now()
	http.Get(url)
	duration := time.Since(start)

	return duration
}
