package partial_interface

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
*
Если необходимо проверить что-то, что зависит от интерфейса с несколькими методами, и при этом известно, что в тесте использовать
будем только один метод из интерфейса, то достаточно создать структуру мока с встроенным интерфейсом, переопределив нужный метод.
*/
func TestGet(t *testing.T) {
	s := Service{
		store: onlyGetMock{},
	}

	assert.Equal(t, 42, s.GetSome(0))
	assert.Panics(t, func() { s.AddSome(0) }, "panic expected")
}
