package nil_receiver

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWrapped(t *testing.T) {
	tests := []struct {
		cfg  CfgWithWrappedCart
		want string
	}{
		{
			cfg:  CfgWithWrappedCart{c: nil},
			want: "Empty cart",
		},
		{
			cfg:  CfgWithWrappedCart{c: &RenderedCart{Cart: nil}},
			want: "Empty cart",
		},

		{
			cfg:  CfgWithWrappedCart{c: &RenderedCart{Cart: &Cart{[]int{1, 2, 3}}, tabID: 1}},
			want: "Cart items:[1 2 3]",
		},
	}
	for _, test := range tests {
		t.Run(test.want, func(t *testing.T) {
			got := renderWrappedCart(test.cfg)
			assert.Equal(t, test.want, got)
		})
	}
}
