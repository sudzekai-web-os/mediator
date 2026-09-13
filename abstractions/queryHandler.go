package abstractions

type IQueryHandler[TQuery IQuery[TResult], TResult any] interface {
	Handle(TQuery) (TResult, error)
}
