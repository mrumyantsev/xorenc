package xorenc

import (
	"errors"
	"io"
	"os"
)

const (
	_ERROR_INSERT string = ". error: "
)

// Make per bit XOR encryption for every byte by a key.
// Returns number of encrypted bytes.
func EncryptData(data []byte, key []byte) int {
	var (
		dLen int = len(data) // data length
		kLen int = len(key)  // key length
		i    int = 0         // data index
		j    int = 0         // key index
	)

	for i < dLen {
		if j >= kLen {
			j = 0
		}

		data[i] = data[i] ^ key[j]

		i++
		j++
	}

	return dLen
}

// Read a data from file, convert every its bit with XOR by a key, and
// then replace initial file content with the result.
// Returns number of encrypted bytes and error (if appeared).
func EcryptFile(path string, key []byte) (int, error) {
	data, err := readFile(path)
	if err != nil {
		return 0, wrapError("could not read file", err)
	}

	nbytes := EncryptData(data, key)

	err = overwriteFile(path, data)
	if err != nil {
		return 0, wrapError("could not overwrite file", err)
	}

	return nbytes, nil
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
