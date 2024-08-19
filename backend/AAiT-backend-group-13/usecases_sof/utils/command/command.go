package icommand

type Ihandler [Command any, Result any] interface{
	Handle(command Command) (Result , error)
}