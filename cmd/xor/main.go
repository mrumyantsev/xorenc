package main

import (
	"flag"
	"io"
	"io/fs"
	"os"
	"runtime"
	"strconv"
	"strings"

	"github.com/mrumyantsev/xor/pkg/xor"
)

const (
	errorExitCode int    = 1
	errorInsert   string = ". error: "
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
		reportError("could not get stdin file info", err)
	}

	stdinFileMode = stdinFileInfo.Mode()
}

func main() {
	// Encrypt the data from stdin to stdout.

	if (stdinFileMode & os.ModeCharDevice) == 0 {
		if len(flagArgs) < 1 {
			reportWrongUsage("encryption key parameter does not presents")
		}

		var (
			encryptKey []byte = []byte(strings.Join(flagArgs, space))
			stdinData  []byte
		)

		stdinData, err = io.ReadAll(os.Stdin)
		if err != nil {
			reportError("could not read data from stdin", err)
		}

		xor.Encrypt(stdinData, encryptKey, nCores)

		_, err = os.Stdout.Write(stdinData)
		if err != nil {
			reportError("could not write encrypted data to stdout", err)
		}

		return
	}

	// Encrypt the file, given in file path parameter.

	if len(flagArgs) < 2 {
		reportWrongUsage("file path or encryption key is missing " +
			"in parameters")
	}

	var (
		filePath   string = flagArgs[0]
		encryptKey []byte = []byte(strings.Join(flagArgs[1:], space))
		nBytes     int
	)

	nBytes, err = xor.EncryptFile(filePath, encryptKey, nCores)
	if err != nil {
		reportError("could not encrypt file", err)
	}

	os.Stdout.WriteString(strconv.Itoa(nBytes) + " bytes encrypted" + eol)
}

func reportError(desc string, err error) {
	os.Stderr.WriteString(appMark + desc + errorInsert + err.Error() + eol)
	os.Exit(errorExitCode)
}

func reportWrongUsage(desc string) {
	os.Stderr.WriteString(appMark + desc + eol)
	os.Stderr.WriteString(usageInfo + eol)
	os.Exit(errorExitCode)
}
