package sync

import "sync"

type Counter struct {
	mu  sync.Mutex
	val int
}

/* We can use mutexes to lock shared memory */
func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.val++
}

func (c *Counter) Value() int {
	return c.val
}
