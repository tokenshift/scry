package parser

import "fmt"

type Parser func(v Vessel) (Result, error)
type Result interface {}

// Matches an exact string.
func Lit(s string) Parser {
	return func(v Vessel) (Result, error) {
		buf := make([]byte, len(s))
		n, err := v.Read(buf)
		if err != nil {
			v.Reset()
			return nil, err
		}
		if n != len(s) {
			v.Reset()
			return nil, fmt.Errorf("Did not match %s", s)
		}
		if string(buf) != s {
			v.Reset()
			return nil, fmt.Errorf("Did not match %s", s)
		}
		v.Consume(len(s))
		return s, nil
	}
}

// Attempts to match any parser in order. The first success will be returned.
func Or(ps ...Parser) Parser {
	return func(v Vessel) (Result, error) {
		for _, p := range(ps) {
			res, err := p(v)
			if err != nil {
				continue
			}
			return res, nil
		}

		return nil, fmt.Errorf("No match found.")
	}
}

// Attempts to match all parsers in order. A slice containing all of the
// results will be returned.
func Seq(ps ...Parser) Parser {
	return func(v Vessel) (Result, error) {
		res := make([]interface{}, len(ps))
		for i, p := range(ps) {
			r, err := p(v)
			if err != nil {
				return nil, err
			}
			res[i] = r
		}
		return res, nil
	}
}
