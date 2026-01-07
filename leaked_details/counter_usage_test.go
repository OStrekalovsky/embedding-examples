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
	// встроенный приватный тип не доступен
	// c.counter

	// утекший метод, который вызвал клиент
	c.Mutex.Lock()
	// deadlock тут
	c.Inc()
	assert.EqualValues(t, 2, c.Value())
}
