package mediator

import (
	"fmt"
	"reflect"

	"github.com/sudzekai-web-os/mediator/abstractions"
)

var handlers = make(map[reflect.Type]any)

func RegisterHandler[
	TQuery abstractions.IQuery[TResult],
	TResult any,
](hnd abstractions.IQueryHandler[TQuery, TResult]) {
	queryType := reflect.TypeFor[TQuery]()
	handlers[queryType] = hnd
}

func Query[
	TQuery abstractions.IQuery[TResult],
	TResult any,
](query TQuery) (TResult, error) {
	queryType := reflect.TypeFor[TQuery]()

	unitypeHnd := handlers[queryType]

	if unitypeHnd == nil {
		panic(fmt.Sprintf("обработчик для запроса типа %v не зарегистрирован", queryType))
	}

	hnd := unitypeHnd.(abstractions.IQueryHandler[TQuery, TResult])

	return hnd.Handle(query)
}
