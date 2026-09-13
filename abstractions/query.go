package abstractions

type IQuery[TResult any] interface {
	IsQuery()
}
