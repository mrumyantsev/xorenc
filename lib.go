package xor

import "errors"

const (
	errorInsert string = ". error: "
)

// Wrap error with a description.
// Returns new error with extended description.
func decorateError(desc string, err error) error {
	return errors.New(desc + errorInsert + err.Error())
}
