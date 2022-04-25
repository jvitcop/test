package errors

type InputError struct {
	msg string
}

func (e InputError) Error() string {
	return e.msg
}
func NewInputError(msg string) *InputError {
	return &InputError{msg}
}

type NotFoundError struct {
	msg string
}

func (e NotFoundError) Error() string {
	return e.msg
}
func NewNotFoundError(msg string) *NotFoundError {
	return &NotFoundError{msg}
}

type InvalidFormatError struct {
	msg string
}

func (e InvalidFormatError) Error() string {
	return e.msg
}
func NewInvalidFormatError(msg string) *InvalidFormatError {
	return &InvalidFormatError{msg}
}
