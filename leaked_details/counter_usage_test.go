package leaked_details_test

import (
	"embedding-examples/leaked_details"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCounters_Inc(t *testing.T) {
	c := &leaked_details.TreadSafeCounter{}
	c.Inc()
	assert.EqualValues(t, 1, c.Value())
	// утекший метод, который вызвал клиент
	c.Mutex.Lock()
	// deadlock тут
	c.Inc()
	assert.EqualValues(t, 2, c.Value())
}

func TestCounters_Dec(t *testing.T) {
	c := &leaked_details.TreadSafeCounter{}
	c.Inc()
	c.Inc()
	// встроенный приватный тип не доступен
	// c.counter
	// но доступен его публичный метод
	c.Dec()

	assert.EqualValues(t, 1, c.Value())
}
