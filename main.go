package main

import "os"

import "github.com/tokenshift/scry/cli"

func main() {
	cli := cli.CreateStd()
	cli.Exec(os.Args[1:]...)
}
