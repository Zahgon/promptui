//go:build !windows
// +build !windows

package promptui

import "github.com/chzyer/readline"

var (
	KeyBackspace rune = readline.CharBackspace
)
