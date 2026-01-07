package inside_out

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

func (c *Cat) LegsDescription() string {
	fmt.Printf("descr %#v\n", c)
	return fmt.Sprintf("Legs:%d", c.Legs())
}

type OctoCat struct {
	Cat
}

func (o *OctoCat) Legs() int {
	fmt.Printf("octocat legs %#v\n", o)
	return 8
}

func Test(t *testing.T) {
	cat := Cat{}
	assert.Equal(t, 4, cat.Legs())
	assert.Equal(t, "Legs:4", cat.LegsDescription())
	octoCat := &OctoCat{}
	assert.Equal(t, 8, octoCat.Legs())
	assert.Equal(t, "Legs:8", octoCat.LegsDescription())
}
