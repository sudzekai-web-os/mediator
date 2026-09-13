package abstractions

type ICommandHandler[TCommand ICommand] interface {
	Handle(TCommand) error
}
