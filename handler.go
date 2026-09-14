package mediator

type IHandler[T any, TResult any] interface {
	Handle(T) TResult
}
