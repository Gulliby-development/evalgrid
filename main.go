package main

import (
	"io"
	"os"
)

func main() {
	args := os.Args[1:]
	input := ""
	if readsStdin(args) {
		data, err := io.ReadAll(os.Stdin)
		if err != nil {
			os.Stderr.WriteString("evalgrid: cannot read standard input\n")
			os.Exit(1)
		}
		input = string(data)
	}
	status, out, errText := run(args, input)
	os.Stdout.WriteString(out)
	os.Stderr.WriteString(errText)
	os.Exit(status)
}
