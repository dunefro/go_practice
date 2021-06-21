package sel

import (
	"fmt"
	"net/http"
	"time"
)

func Racer(url1 string, url2 string) (winner string) {

	start1 := time.Now()
	http.Get(url1)
	time1 := time.Since(start1)

	start2 := time.Now()
	http.Get(url2)
	time2 := time.Since(start2)

	fmt.Printf("%q: %d\n%q: %d\n", url1, time1, url2, time2)

	if time1 > time2 {
		return url2
	}

	return url1
}
