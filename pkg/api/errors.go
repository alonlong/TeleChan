package api

import "fmt"

func (e Error) Error() string {
	return fmt.Sprintf("%d: %s", e.Code, e.Message)
}

// Here we define erros that can be exposed in server API replies
var (
	// ErrorInternal means server error, if returned this is a signal
	// that something went wrong with TeleChan itself.
	ErrorInternal = &Error{
		Code:    100,
		Message: "internal server error",
	}
)
