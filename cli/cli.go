package cli

import "fmt"
import "os"

import "github.com/tokenshift/args"
import "github.com/tokenshift/scry/repl"

type CLI struct {
	in, out, err *os.File
}

// Creates a CLI that will read from and write to the specified files.
func Create(in, out, err *os.File) CLI {
	return CLI { in, out, err }
}

// Creates a CLI wrapping STDIN, STDOUT and STDERR.
func CreateStd() CLI {
	return Create(os.Stdin, os.Stdout, os.Stderr)
}

// Executes a Scry command.
func (cli CLI) Exec(params ...string) {
	result, err := args.Load(params).
		ExpectParamNamed("command").
		ChopAndValidate()

	if err != nil {
		fmt.Println(err)
		return
	}

	command := result.ParamNamed("command")
	if result.ParamNamed("command") == "repl" {
		repl := repl.Create(cli.in, cli.out, cli.err)
		repl.Run()
	} else {
		fmt.Printf("Unrecognized command: %s\n", command)
	}
}
