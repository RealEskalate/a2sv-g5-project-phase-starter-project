/*
Package command provides a generic interface for handling commands.

It includes the `IHandler` interface for processing commands of any type and returning results of any type.
*/
package icmd

// IHandler defines a generic interface for handling commands.
//
// Type Parameters:
// - Command: The type of the command to be handled.
// - Result: The type of the result returned after handling the command.
type IHandler[Command any, Result any] interface {

	// Handle processes the provided command and returns the result or an error.
	Handle(command Command) (Result, error)
}
