package main

import (
	"flag"
	"io"
	"os"
	"runtime"
	"strconv"
	"strings"

	"github.com/mrumyantsev/xor/pkg/lib/e"
	"github.com/mrumyantsev/xor/pkg/xor"
)

const (
	appMark   = "xor: "
	usageInfo = "Usage:\n  - xor <filepath> <enckey>\n" +
		"  - cat ./def_file | xor <enckey> > ./out_file\n" +
		"  - xor <enckey> < ./def_file > ./out_file\n\n" +
		"The encryption key (<enckey>) may contain spaces.\n\n" +
		"The decryption is provided by repeating of encryption\n" +
		"with the same key, that was used for encryption."
	space = " "
	eol   = "\n"

	errorExitCode = 1
)

var (
	cliArgs  []string
	cpuCores = runtime.NumCPU()
)

func init() {
	flag.Parse()

	cliArgs = flag.Args()
}

func main() {
	stdinFileInfo, err := os.Stdin.Stat()
	if err != nil {
		printError("could not get stdin file info", err)

		os.Exit(errorExitCode)
	}

	if (stdinFileInfo.Mode() & os.ModeCharDevice) == 0 {
		encryptStdinData()

		return
	}

	encryptFile()
}

func encryptStdinData() {
	if len(cliArgs) < 1 {
		printUsage("encryption key parameter does not presents")

		os.Exit(errorExitCode)
	}

	encryptKey := []byte(strings.Join(cliArgs, space))

	stdinData, err := io.ReadAll(os.Stdin)
	if err != nil {
		printError("could not read data from stdin", err)

		os.Exit(errorExitCode)
	}

	xor.Encrypt(stdinData, encryptKey, cpuCores)

	if _, err = os.Stdout.Write(stdinData); err != nil {
		printError("could not write encrypted data to stdout", err)

		os.Exit(errorExitCode)
	}
}

func encryptFile() {
	if len(cliArgs) < 2 {
		printUsage("file path or encryption key is missing " +
			"in parameters")

		os.Exit(errorExitCode)
	}

	filePath := cliArgs[0]

	encryptKey := []byte(strings.Join(cliArgs[1:], space))

	encBytes, err := xor.EncryptFile(filePath, encryptKey, cpuCores)
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
