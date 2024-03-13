package fileops

import (
	"io"
	"os"

	"github.com/mrumyantsev/xor/pkg/lib/e"
)

const (
	errOpenFile  = "could not open file"
	errReadFile  = "could not read from file"
	errWriteFile = "could not write to file"
)

// ReadFile reads the data from a file by its path.
// Returns the file data and the error with its description.
func ReadFile(path string) ([]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, e.Wrap(errOpenFile, err)
	}
	defer func() { _ = f.Close() }()

	data, err := io.ReadAll(f)
	if err != nil {
		return nil, e.Wrap(errReadFile, err)
	}

	return data, nil
}

// OverwriteFile overwrites the whole file by its path with the data.
// If the file does not exists by given path it will create a new file.
// Returns an error with its description.
func OverwriteFile(path string, data []byte) error {
	f, err := os.OpenFile(path, os.O_TRUNC|os.O_WRONLY, 0)
	if err != nil {
		return e.Wrap(errOpenFile, err)
	}
	defer func() { _ = f.Close() }()

	_, err = f.Write(data)
	if err != nil {
		return e.Wrap(errWriteFile, err)
	}

	return nil
}
