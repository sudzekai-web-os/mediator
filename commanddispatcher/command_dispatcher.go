package commanddispatcher

import (
	"fmt"
	"reflect"

	"github.com/sudzekai-web-os/mediator/abstractions"
)

var handlers = make(map[reflect.Type]any)

func RegisterHandler[TCommand abstractions.ICommand](
	hnd abstractions.ICommandHandler[TCommand],
) {
	commandType := reflect.TypeFor[TCommand]()
	handlers[commandType] = hnd
}

func Execute[TCommand abstractions.ICommand](
	command TCommand,
) error {
	commandType := reflect.TypeFor[TCommand]()

	unitypeHnd := handlers[commandType]

	if unitypeHnd == nil {
		panic(fmt.Sprintf("обработчик для команды типа %v не зарегистрирован", commandType))
	}

	hnd := unitypeHnd.(abstractions.ICommandHandler[TCommand])

	return hnd.Handle(command)
}
