package main

import (
	"flag"
	"io"
	"io/fs"
	"os"
	"runtime"
	"strconv"
	"strings"

	"github.com/mrumyantsev/xor/pkg/lib/e"
	"github.com/mrumyantsev/xor/pkg/xor"
)

const (
	errorExitCode int    = 1
	appMark       string = "xor: "
	usageInfo     string = "Usage:\n  - xor <filepath> <enckey>\n" +
		"  - cat ./def_file | xor <enckey> > ./out_file\n" +
		"  - xor <enckey> < ./def_file > ./out_file\n\n" +
		"The encryption key (<enckey>) may contain spaces.\n\n" +
		"The decryption is provided by repeating of encryption\n" +
		"with the same key, that was used for encryption."
	space string = " "
	eol   string = "\n"
)

var (
	flagArgs      []string
	stdinFileInfo fs.FileInfo
	stdinFileMode fs.FileMode
	nCores        int = runtime.NumCPU()
	err           error
)

func init() {
	flag.Parse()

	flagArgs = flag.Args()

	stdinFileInfo, err = os.Stdin.Stat()
	if err != nil {
		printError("could not get stdin file info", err)

		os.Exit(errorExitCode)
	}

	stdinFileMode = stdinFileInfo.Mode()
}

func main() {
	// Encrypt the data from stdin to stdout.

	if (stdinFileMode & os.ModeCharDevice) == 0 {
		if len(flagArgs) < 1 {
			printUsage("encryption key parameter does not presents")

			os.Exit(errorExitCode)
		}

		var (
			encryptKey []byte = []byte(strings.Join(flagArgs, space))
			stdinData  []byte
		)

		stdinData, err = io.ReadAll(os.Stdin)
		if err != nil {
			printError("could not read data from stdin", err)

			os.Exit(errorExitCode)
		}

		xor.Encrypt(stdinData, encryptKey, nCores)

		_, err = os.Stdout.Write(stdinData)
		if err != nil {
			printError("could not write encrypted data to stdout", err)

			os.Exit(errorExitCode)
		}

		return
	}

	// Encrypt the file, given in file path parameter.

	if len(flagArgs) < 2 {
		printUsage("file path or encryption key is missing " +
			"in parameters")

		os.Exit(errorExitCode)
	}

	var (
		filePath   string = flagArgs[0]
		encryptKey []byte = []byte(strings.Join(flagArgs[1:], space))
		encBytes   int
	)

	encBytes, err = xor.EncryptFile(filePath, encryptKey, nCores)
	if err != nil {
		printError("could not encrypt file", err)

		os.Exit(errorExitCode)
	}

	printDone(encBytes)
}

func printDone(encBytes int) {
	os.Stdout.WriteString(strconv.Itoa(encBytes) + " bytes encrypted" + eol)
}

func printError(desc string, err error) {
	os.Stderr.WriteString(appMark + e.Wrap(desc, err).Error() + eol)
}

func printUsage(errMsg string) {
	os.Stderr.WriteString(appMark + errMsg + eol)
	os.Stderr.WriteString(usageInfo + eol)
}
