package parser

import "io"

// A buffered, reversable input reader.
type Vessel interface {
	io.Reader

	// Discards a number of bytes from the start of the buffer.
	Consume(int)

	// Resets the reader position to the start of the buffer, as though the
	// input after that point had not been read.
	Reset()
}

// Wraps a string as a vessel.
func StringVessel(input string) Vessel {
	return &sVessel {
		input,
		0,
	}
}

type sVessel struct {
	input string
	offset int
}

func (v *sVessel) Consume(i int) {
	v.input = v.input[i:]
	v.offset = 0
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func (v sVessel) Read(buf []byte) (int, error) {
	start := v.offset
	end := min(v.offset + len(buf), len(v.input))
	copy(buf, v.input[start:end])
	if end - start <= 0 {
		return 0, io.EOF
	} else {
		return end - start, nil
	}
}

func (v sVessel) Reset() {
	v.offset = 0
}
