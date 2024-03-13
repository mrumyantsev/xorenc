package xor

import (
	"github.com/mrumyantsev/xor/pkg/fileops"
	"github.com/mrumyantsev/xor/pkg/lib/e"
)

const (
	errCompleteFileReading     = "could not complete file reading"
	errCompleteFileOverwriting = "could not complete file overwriting"
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
	data, err := fileops.ReadFile(path)
	if err != nil {
		return 0, e.Wrap(errCompleteFileReading, err)
	}

	Encrypt(data, key, nWorkers)

	err = fileops.OverwriteFile(path, data)
	if err != nil {
		return 0, e.Wrap(errCompleteFileOverwriting, err)
	}

	return len(data), nil
}
