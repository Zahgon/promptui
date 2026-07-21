package promptui

import "errors"

var ErrEOF = errors.New("^D")

var ErrInterrupt = errors.New("^C")

var ErrAbort = errors.New("")

type ValidateFunc func(string) error
