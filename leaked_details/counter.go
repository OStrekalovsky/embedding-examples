package leaked_details

import "sync"

type TreadSafeCounter struct {
	counter
	sync.Mutex
}

func (c *TreadSafeCounter) Inc() {
	c.Lock()
	c.counter++
	c.Unlock()
}

func (c *TreadSafeCounter) Value() int64 {
	c.Lock()
	defer c.Unlock()

	return int64(c.counter)
}

type counter int64
