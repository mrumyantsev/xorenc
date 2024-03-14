package e

import "fmt"

func WrapOrMsg(msg string, err error) error {
	if err != nil {
		return Wrap(msg, err)
	}

	return fmt.Errorf("%s", msg)
}

func WrapOrNil(msg string, err error) error {
	if err != nil {
		return Wrap(msg, err)
	}

	return nil
}

func Wrap(msg string, err error) error {
	return fmt.Errorf("%s: %w", msg, err)
}
