package deferandselect

import (
	"net/http"
	"time"
)

func Racer(url_a string, url_b string) string {
	startA := time.Now()
	http.Get(url_a)
	aDuration := time.Since(startA)

	startB := time.Now()
	http.Get(url_b)
	bDuration := time.Since(startB)

	if aDuration > bDuration {
		return url_b
	}
	return url_a
}
