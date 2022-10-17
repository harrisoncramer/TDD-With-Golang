package deferandselect

import (
	"fmt"
	"net/http"
	"time"
)

/*
The "select" keyword lets you wait on multiple channels. The
first one to send a value wins.
*/
func Racer(a, b string, timeout time.Duration) (winner string, e error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
		/* The time.After() function returns a chan. We can use
		   it here to avoid blocking forever if our pings don't return */
	case <-time.After(timeout * time.Millisecond):
		return "", fmt.Errorf("Timed out waiting for %s and %s", a, b)
	}
}

/*
Ping returns a channel immediately, but only closes
it after the get call finishes inside the goroutine
*/
func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}

/*
This implementation hits each of the endpoints in
sequence and then measures the response time of each
of them to determine a winner
*/
// func measureResponseTime(url string) time.Duration {
// 	start := time.Now()
// 	http.Get(url)
// 	return time.Since(start)
// }
//
// func Racer(a, b string) (winner string) {
//
// 	aDuration := measureResponseTime(a)
// 	bDuration := measureResponseTime(b)
//
// 	if aDuration < bDuration {
// 		return a
// 	}
// 	return b
// }
