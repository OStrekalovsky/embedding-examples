package inside_out_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Cat struct{}

func (c *Cat) Legs() int {
	fmt.Printf("cat legs %#v\n", c)
	return 4
}

func (c *Cat) LegsDescription(counter LegsCounter) string {
	fmt.Printf("descr %#v\n", c)
	return fmt.Sprintf("Legs:%d", counter.Legs())
}

type LegsCounter interface {
	// Дополнительный параметр нужен, т.к. внутри вызывается переопределяемый метод, и нам нужно сохранить оригинальный ресивер.
	LegsDescription(LegsCounter) string
	// Дополнительный параметр не нужен т.к. внутри этого метода не вызвается ничего, что может быть переопределено (для чего нужно сохранить оригинальный ресивер).
	Legs() int
}

type OctoCat struct {
	*Cat
}

func (o *OctoCat) Legs() int {
	fmt.Printf("octocat legs %#v\n", o)
	return 8
}

func TestOctoCatFixed(t *testing.T) {
	cat := &Cat{}
	assert.Equal(t, 4, cat.Legs())
	assert.Equal(t, "Legs:4", cat.LegsDescription(cat))

	octoCat := &OctoCat{}
	assert.Equal(t, 8, octoCat.Legs())
	assert.Equal(t, "Legs:8", octoCat.LegsDescription(octoCat))
}
