package nil_receiver

import "fmt"

type RenderedCart struct {
	*Cart
	tabID int
}

type CfgWithWrappedCart struct {
	c *RenderedCart
}

// Раскомментировать метод, чтобы исправить ошибку
//func (c *RenderedCart) Empty() bool {
//	if c == nil {
//		return true
//	}
//	return c.Cart.Empty()
//}

func renderWrappedCart(cfg CfgWithWrappedCart) string {
	if cfg.c.Empty() {
		return "Empty cart"
	}

	return fmt.Sprintf("Cart items:%v", cfg.c.items)
}
