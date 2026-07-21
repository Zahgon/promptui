package promptui

import "github.com/chzyer/readline"

var (
	KeyEnter rune = readline.CharEnter

	KeyCtrlH rune = readline.CharCtrlH

	KeyPrev        rune = readline.CharPrev
	KeyPrevDisplay      = "↑"

	KeyNext        rune = readline.CharNext
	KeyNextDisplay      = "↓"

	KeyBackward        rune = readline.CharBackward
	KeyBackwardDisplay      = "←"

	KeyForward        rune = readline.CharForward
	KeyForwardDisplay      = "→"
)
