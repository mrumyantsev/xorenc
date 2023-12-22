package xorenc

import (
	"errors"
	"io"
	"os"
)

const (
	_ERROR_INSERT string = ". error: "
)

// Make per bit XOR convertation with a key.
func EncryptData(data []byte, key []byte) {
	var (
		dLen = len(data) // data length
		kLen = len(key)  // key length
		i    = 0         // data index
		j    = 0         // key index
	)

	for i < dLen {
		if j >= kLen {
			j = 0
		}

		data[i] = data[i] ^ key[j]

		i++
		j++
	}
}

// Read a data from file, convert every its bit with XOR by a key, and
// then replace initial file content with the result.
func EcryptFile(path string, key []byte) error {
	data, err := readFile(path)
	if err != nil {
		return wrapError("could not read file", err)
	}

	EncryptData(data, key)

	err = overwriteFile(path, data)
	if err != nil {
		return wrapError("could not overwrite file", err)
	}

	return nil
}

// Read data from file.
func readFile(path string) ([]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	data, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// Overwrite whole file, if exists, or create new file with the data.
func overwriteFile(path string, data []byte) error {
	f, err := os.OpenFile(path, os.O_TRUNC|os.O_WRONLY, 0)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(data)
	if err != nil {
		return err
	}

	return nil
}

// Wrap error with a description.
func wrapError(desc string, err error) error {
	return errors.New(desc + _ERROR_INSERT + err.Error())
}
