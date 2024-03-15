package main

import (
	"fmt"
	"io"
	"os"

	"github.com/mrumyantsev/xor/pkg/lib/e"
	"github.com/mrumyantsev/xor/pkg/xor"
)

const (
	appName  = "xor"
	verText  = "1.0.0"
	helpText = `Usage to files:
  - xor ORIG_FILE ENCRYPT_KEY

Usage to stdin->stdout:
  - xor ENCRYPT_KEY < ORIG_FILE > OUTPUT_FILE
  - cat ORIG_FILE | xor ENCRYPT_KEY > OUTPUT_FILE

Options:
  -h, --help     display this help and exit
  -v, --version  output version information and exit

The encryption key (ENCRYPT_KEY) may contain spaces.

To decrypt data, use the same encryption key that was
used to encrypt it.`

	errorExitCode = 1
)

var (
	isHelp    = false
	isVersion = false

	filePath   string
	encryptKey []byte
)

func main() {
	parseHelpArgs()

	if isHelp {
		help()

		return
	}

	parseVerArgs()

	if isVersion {
		version()

		return
	}

	if isStdinData() {
		parseStdinArgs()
		encryptStdinData()

		return
	}

	parseFileArgs()
	encryptFile()
}

func isStdinData() bool {
	stdinFileInfo, err := os.Stdin.Stat()
	if err != nil {
		fatal("could not get stdin file info", err)
	}

	return (stdinFileInfo.Mode() & os.ModeCharDevice) == 0
}

func encryptStdinData() {
	stdinData, err := io.ReadAll(os.Stdin)
	if err != nil {
		fatal("could not read data from stdin", err)
	}

	xor.Encrypt(stdinData, encryptKey)

	if _, err = os.Stdout.Write(stdinData); err != nil {
		fatal("could not write encrypted data to stdout", err)
	}
}

func encryptFile() {
	encBytes, err := xor.EncryptFile(filePath, encryptKey)
	if err != nil {
		fatal("could not encrypt file", err)
	}

	fmt.Printf("%d bytes encrypted\n", encBytes)
}

func fatal(desc string, err error) {
	fmt.Fprintf(os.Stderr, "%s: %v\n", appName, e.WrapOrMsg(desc, err))
	fmt.Fprintf(os.Stderr, "Try '%s --help' for more information.\n", appName)

	os.Exit(errorExitCode)
}

func help() {
	fmt.Println(helpText)
}

func version() {
	fmt.Printf("%s version %s\n", appName, verText)
}
