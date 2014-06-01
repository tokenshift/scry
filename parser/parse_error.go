package parser

import "fmt"

import "github.com/dustin/go-parse"

type ParseError interface {
	error

	Msg() string
	Pos() parsec.Position
}

type parseError struct {
	msg string
	pos parsec.Position
}

func newParseError(msg string, pos parsec.Position) ParseError {
	return parseError {
		msg: msg,
		pos: pos,
	}
}

func (e parseError) Error() string {
	return fmt.Sprintf("%s (at %v)", e.msg, e.pos)
}

func (e parseError) Msg() string {
	return e.msg
}

func (e parseError) Pos() parsec.Position {
	return e.pos
}
