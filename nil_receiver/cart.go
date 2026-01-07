package nil_receiver

import "fmt"

type Cart struct {
	items []int
}

func (c *Cart) Empty() bool {
	if c == nil {
		return true
	}
	return len(c.items) == 0
}

type Cfg struct {
	c *Cart
}

func renderCart(cfg Cfg) string {
	if cfg.c.Empty() {
		return "Empty cart"
	}

	return fmt.Sprintf("Cart items:%v", cfg.c.items)
}
