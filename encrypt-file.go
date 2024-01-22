package xor

import (
	"io"
	"os"
)

const (
	errExecReadingSeq     = "could not execute reading sequence"
	errExecOverwritingSeq = "could not execute overwriting sequence"
	errOpenFile           = "could not open file"
	errReadFile           = "could not read from file"
	errWriteFile          = "could not write to file"
)

// EncryptFile performs per-bit XOR encryption of file, specified in
// path, by the key.
//
// The decryption is provided by repeating of encryption with the same
// key, that was used for encryption.
//
// The data processing operation in this function is processed with the
// workers, which divides the data to the equal pieces and encrypts
// each. The number of workers is capped to the length of the data, and
// it makes no processing, if given 0 or less.
//
// Returns the number of the encrypted bytes and the error with its
// description.
func EncryptFile(path string, key []byte, nWorkers int) (nBytes int, err error) {
	data, err := readFile(path)
	if err != nil {
		return 0, wrapError(errExecReadingSeq, err)
	}

	Encrypt(data, key, nWorkers)

	err = overwriteFile(path, data)
	if err != nil {
		return 0, wrapError(errExecOverwritingSeq, err)
	}

	return len(data), nil
}

// readFile reads the data from a file by its path.
// Returns the file data and the error with its description.
func readFile(path string) ([]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, wrapError(errOpenFile, err)
	}
	defer f.Close()

	data, err := io.ReadAll(f)
	if err != nil {
		return nil, wrapError(errReadFile, err)
	}

	return data, nil
}

// overwriteFile overwrites the whole file by its path with the data. If
// the file does not exists yet, it creates a new one.
// Returns an error with its description.
func overwriteFile(path string, data []byte) error {
	f, err := os.OpenFile(path, os.O_TRUNC|os.O_WRONLY, 0)
	if err != nil {
		return wrapError(errOpenFile, err)
	}
	defer f.Close()

	_, err = f.Write(data)
	if err != nil {
		return wrapError(errWriteFile, err)
	}

	return nil
}
