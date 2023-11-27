package xore

import (
	"github.com/mrumyantsev/go-fsops"
)

// Read a data from file, convert every its bit with XOR by a key, and
// then replace initial file content with the result.
func EcryptFile(path string, key []byte) error {
	var (
		data []byte
		err  error
	)

	data, err = fsops.ReadFile(path)
	if err != nil {
		return err
	}

	EncryptBytes(data, key)

	err = fsops.OverwriteFile(path, data)
	if err != nil {
		return err
	}

	return nil
}

// Make per bit convertation with XOR by a key.
func EncryptBytes(data []byte, key []byte) {
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
