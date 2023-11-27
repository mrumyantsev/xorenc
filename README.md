# xore - XOR Encrypter

This CLI application encrypts and decrypts files within file system. And it can be imported as Go library. It uses the same encrypt key to decrypt the file that was encrypted. **Warning**: Always remember the key you using to encrypt files. Losing the key can cause permanent loss of data.

# Demonstration

Call this command twice and every time see what is happening to encrypting demo file called *Daisy.txt*:

```
make run/test
```

Spoiler: It was encrypted with the key ***SuNnY DaY*** twice. First encrypting makes the text unreadable, and second restores it right back. Try to encrypt or decrypt it independently by building and using the *xore* application.

# Usage

Build by calling **make** command in terminal.

```
make build
```

Execute the application with 2 additional parameters:

```
xore <path/to/file> <any amount of words as encrypt key>
```
