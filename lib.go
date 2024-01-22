package xor

import "errors"

const errorInsert string = ". error: "

// wrapError wraps a given error with the description, inserting
// ". error: " text in between. If the error is nil, is cause panic.
// Returns the new error, which text is decorated by the specified
// description.
func wrapError(desc string, err error) error {
	return errors.New(desc + errorInsert + err.Error())
}
