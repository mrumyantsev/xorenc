package dataenc

// A DataEnc structure is responsible for encryption of data in terms of
// XOR data conversion.
type DataEnc struct {
}

// New creates a new DataEnc instance.
// Returns a pointer to DataEnc struct in heap.
func New() *DataEnc {
	return &DataEnc{}
}

// Encrypt performs per-bit XOR encryption of data by the key.
// The decryption is provided by repeating of encryption with the same
// key, that was in use in encryption.
// Returns the number of the encrypted bytes.
func (e *DataEnc) Encrypt(data []byte, key []byte) (nBytes int) {
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
