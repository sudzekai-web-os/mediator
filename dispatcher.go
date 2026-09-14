package mediator

import (
	"fmt"
	"reflect"
)

var handlers = make(map[reflect.Type]any)

func RegisterHandler[
	TQuery any,
	TResult any,
](hnd IHandler[TQuery, TResult]) {
	queryType := reflect.TypeFor[TQuery]()

	handlers[queryType] = hnd
}

func Dispatch[
	TQuery any,
	TResult any,
](query TQuery, result *TResult) error {
	queryType := reflect.TypeFor[TQuery]()

	unitypeHnd := handlers[queryType]

	if unitypeHnd == nil {
		return fmt.Errorf(
			"обработчик для запроса типа %v не зарегистрирован",
			queryType,
		)
	}

	hnd := unitypeHnd.(IHandler[TQuery, TResult])

	r := hnd.Handle(query)

	*result = r

	return nil
}
