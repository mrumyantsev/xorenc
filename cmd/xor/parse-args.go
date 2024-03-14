package main

import (
	"os"
	"strings"
)

func parseHelpArgs() {
	if len(os.Args) < 2 {
		return
	}

	if (os.Args[1] == "-h") || (os.Args[1] == "--help") {
		isHelp = true
	}
}

func parseVerArgs() {
	if len(os.Args) < 2 {
		return
	}

	if (os.Args[1] == "-v") || (os.Args[1] == "--version") {
		isVersion = true
	}
}

func parseStdinArgs() {
	if len(os.Args) < 2 {
		fatal("missing encryption key operand", nil)
	}

	encryptKey = []byte(strings.Join(os.Args[1:], " "))
}

func parseFileArgs() {
	if len(os.Args) < 3 {
		fatal("missing file path or encryption key operands", nil)
	}

	filePath = os.Args[1]
	encryptKey = []byte(strings.Join(os.Args[2:], " "))
}
