package util

import "fmt"

func NotImplementedError(name string) error {
	return fmt.Errorf("%s is not yet implemented.", name)
}
