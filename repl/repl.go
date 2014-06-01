package repl

import "bufio"
import "fmt"
import "os"

type Repl struct {
	in *bufio.Reader
	out, err *os.File
	input string
}

// Creates a REPL that will read from and write to the specified files.
func Create(in, out, err *os.File) Repl {
	return Repl {
		bufio.NewReader(in),
		out,
		err,
		"",
	}
}

// Creates a REPL wrapping STDIN, STDOUT and STDERR.
func CreateStd() Repl {
	return Create(os.Stdin, os.Stdout, os.Stderr)
}

// Starts running the REPL.
func (r Repl) Run() {
	for r.prompt() {
		err := r.process()

		if err != nil {
			fmt.Fprintf(r.err, "Error: %v", err)
		}
	}
}

// Prompt for further input.
func (r *Repl) prompt() bool {
	r.out.WriteString(" > ")
	input, err := r.in.ReadString('\n')
	if err != nil {
		return false
	} else {
		r.input = input
		return true
	}
}

// Process the previous input.
func (r *Repl) process() error {
	fmt.Fprintf(r.err, "=> %s\n", r.input)
	return nil
}
