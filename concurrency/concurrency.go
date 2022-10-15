package concurrency

type result struct {
	string
	bool
}

type WebsiteChecker func(string) bool

/*
Here we are starting a goroutine for each check because we know
that the check can take some time. However, we cannot parallelize the writing
of the output to the map, because that would cause a race condition where two goroutines
are trying to write to the same memory address at the same time. Instead, we send the
result of the goroutine to a channel, which coordinates the writing to our map.
*/

/*
We can benchmark this code with go test -bench=.
We can look for race conditions with go test -race
*/

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {

	results := make(map[string]bool)
	channel := make(chan (result))
	for _, url := range urls {
		go func(u string) {
			channel <- result{u, wc(u)} /* Send result to channel... */
		}(url)
	}

	/* Pass in an argument or it'll use a "reference" to previous url variable,
	which is constantly changing as the loop changes*/

	for i := 0; i < len(urls); i++ {
		r := <-channel /* Read 100 results from channel into variable as they are ready */
		results[r.string] = r.bool
	}

	return results
}
