package dataenc

type DataEnc struct {
}

func New() *DataEnc {
	return &DataEnc{}
}

// Encrypt performs per-bit XOR encryption of data by the key. The same
// key must be used for the decryption.
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
