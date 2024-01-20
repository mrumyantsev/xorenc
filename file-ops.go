package xor

import (
	"io"
	"os"
)

// Read a data from file, convert every its bit with XOR by a key, and
// then replace initial file content with the result.
// Returns number of encrypted bytes and error (if appeared).
func EncryptFile(path string, key []byte) (int, error) {
	data, err := readFile(path)
	if err != nil {
		return 0, decorateError("could not read file", err)
	}

	nbytes := EncryptData(data, key)

	err = overwriteFile(path, data)
	if err != nil {
		return 0, decorateError("could not overwrite file", err)
	}

	return nbytes, nil
}

// Read data from file.
// Returns file data and error (if appeared).
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
// Returns error (if appeared).
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
