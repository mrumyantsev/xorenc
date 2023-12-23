# Xcrypter

This CLI application encrypts/decrypts files within file system by XOR per bit converting. Also it is a Go library, so it is able to import converting functions from other applications. It uses the same encrypt key to decrypt the file that was encrypted. **Warning**: Always remember the key you using to encrypt files. Losing the key can cause permanent loss of data.

# Usage

Build by calling **make** command in terminal.

```
make build
```

Execute the application with 2 additional parameters:

```
cd build
usage: xcrypter <path/to/file> <any number of words as an encryption key>
```
