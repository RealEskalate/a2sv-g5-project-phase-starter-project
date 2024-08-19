package iquery 


type Ihandler[Query any, Result any] interface {
	Handle(query Query) ( Result, error)
}