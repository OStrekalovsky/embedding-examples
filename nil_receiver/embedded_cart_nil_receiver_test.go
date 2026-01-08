package nil_receiver

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWrapped(t *testing.T) {
	tests := []struct {
		name string
		cfg  CfgWithWrappedCart
		want string
	}{
		{
			name: "nil outer",
			cfg:  CfgWithWrappedCart{c: nil},
			want: "Empty cart",
		},
		{
			name: "nil inner",
			cfg:  CfgWithWrappedCart{c: &RenderedCart{Cart: nil}},
			want: "Empty cart",
		},

		{
			name: "regular",
			cfg:  CfgWithWrappedCart{c: &RenderedCart{Cart: &Cart{[]int{1, 2, 3}}, tabID: 1}},
			want: "Cart items:[1 2 3]",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := renderWrappedCart(test.cfg)
			assert.Equal(t, test.want, got)
		})
	}
}
