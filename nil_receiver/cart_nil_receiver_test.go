package nil_receiver

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBasic(t *testing.T) {
	tests := []struct {
		cfg  Cfg
		want string
	}{
		{
			cfg:  Cfg{c: nil},
			want: "Empty cart",
		},
		{
			cfg:  Cfg{c: &Cart{[]int{1, 2, 3}}},
			want: "Cart items:[1 2 3]",
		},
	}
	for _, test := range tests {
		t.Run(test.want, func(t *testing.T) {
			got := renderCart(test.cfg)
			assert.Equal(t, test.want, got)
		})
	}
}
