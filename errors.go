package hierarchy

import "fmt"

const (
	// DataFileOpenError message.
	DataFileOpenError = "Could not open the hierarchy data file"
	// DataFileCloseError message.
	DataFileCloseError = "Could not close the hierarchy data file"
	// DataFileDecodeError message.
	DataFileDecodeError = "Could not JSON decode the hierarchy data file"
	// EmployeeNotFoundError message.
	EmployeeNotFoundError = "Could not find employee"
)

// Error is an error implementation that
// includes type, description and message.
type Error struct {
	Type        string
	Description error
	Message     string
}

// Error implementation.
func (e Error) Error() string {
	if e.Description != nil {
		return fmt.Sprintf("[ERROR] %v:\nError description: %v",
			e.Type, e.Description)
	}

	return fmt.Sprintf("[ERROR] %v:\t%v",
		e.Type, e.Message)
}

// NewError is the constructor for Error.
func NewError(errType string, err error) error {
	return Error{
		errType,
		err,
		""}
}

// NewErrorMsg is the constructor for Error with a message.
func NewErrorMsg(errType string, msg string) error {
	return Error{
		errType,
		nil,
		msg}
}
