package xor

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
