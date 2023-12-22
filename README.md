# XOR Encrypter

This CLI application encrypts and decrypts files within file system. And it can be imported as Go library. It uses the same encrypt key to decrypt the file that was encrypted. **Warning**: Always remember the key you using to encrypt files. Losing the key can cause permanent loss of data.

# Usage

Build by calling **make** command in terminal.

```
make build
```

Execute the application with 2 additional parameters:

```
cd build
usage: xorenc <path/to/file> <any number of words as an encryption key>
```
