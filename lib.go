package xor

import "errors"

const (
	errorInsert string = ". error: "
)

// decorateError decorates a given error with a given description,
// inserting ". error: " text at the end of the description.
// Returns the new error, which text is decorated by the specified
// description. If the given error is nil, then the returning error will
// only contain the description text.
func decorateError(desc string, err error) error {
	if err != nil {
		desc += errorInsert + err.Error()
	}

	return errors.New(desc)
}
