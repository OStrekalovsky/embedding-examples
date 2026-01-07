package interface_embedding_into_errors

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
*
Встраивание интерфейса генератора ответа для кастомной ошибки позволяет описать единый верхнеуровневый механизм их обработки,
который при этом порождает разные результаты.

Чтобы ошибка стала обрабатываемой достаточно встроить структуру с готовой реализацией или сам интерфейс.
*/
func Test(t *testing.T) {
	choice := 1
	// верхнеуровнево работаем только с интерфейсом error
	err := checkChoice(choice)
	if err != nil {
		code, msg := handleError(err)
		assert.Equal(t, 400, code)
		assert.Equal(t, `{"bad address":"choose another PVZ"}`, msg)
		return
	}
	assert.Fail(t, "impossible")
}

func checkChoice(choice int) error {
	switch choice {
	case 1:
		return &BadPVZAddressError{
			&BadAddressResponseGenerator{msg: "choose another PVZ"},
		}
	case 2:
		return &BadCourierAddressError{
			&BadAddressResponseGenerator{msg: "choose another courier address"},
		}
	default:
		return errors.New("unexpected error")
	}
}

func handleError(err error) (int, string) {
	var handler errorResponseGenerator

	switch {
	case errors.As(err, &handler):
		// если ошибка реализовала интерфейс генератора ответа - используем его для ответа.
		return 400, handler.GetMessage()
	default:
		// если ошибка не реализовала интерфейс генератора ответа - используем текст ошибки.
		return 500, err.Error()
	}
}
