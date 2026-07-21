package promptui

import (
	"fmt"
	"text/template"
)

const esc = "\033["

type attribute int

const (
	reset attribute = iota

	FGBold
	FGFaint
	FGItalic
	FGUnderline
)

const (
	FGBlack attribute = iota + 30
	FGRed
	FGGreen
	FGYellow
	FGBlue
	FGMagenta
	FGCyan
	FGWhite
)

const (
	BGBlack attribute = iota + 40
	BGRed
	BGGreen
	BGYellow
	BGBlue
	BGMagenta
	BGCyan
	BGWhite
)

var ResetCode = fmt.Sprintf("%s%dm", esc, reset)

const (
	hideCursor = esc + "?25l"
	showCursor = esc + "?25h"
	clearLine  = esc + "2K"
)

var FuncMap = template.FuncMap{
	"black":     Styler(FGBlack),
	"red":       Styler(FGRed),
	"green":     Styler(FGGreen),
	"yellow":    Styler(FGYellow),
	"blue":      Styler(FGBlue),
	"magenta":   Styler(FGMagenta),
	"cyan":      Styler(FGCyan),
	"white":     Styler(FGWhite),
	"bgBlack":   Styler(BGBlack),
	"bgRed":     Styler(BGRed),
	"bgGreen":   Styler(BGGreen),
	"bgYellow":  Styler(BGYellow),
	"bgBlue":    Styler(BGBlue),
	"bgMagenta": Styler(BGMagenta),
	"bgCyan":    Styler(BGCyan),
	"bgWhite":   Styler(BGWhite),
	"bold":      Styler(FGBold),
	"faint":     Styler(FGFaint),
	"italic":    Styler(FGItalic),
	"underline": Styler(FGUnderline),
}

func upLine(n uint) string { _ = "STUB: not implemented"; return "" }

func movementCode(n uint, code rune) string { _ = "STUB: not implemented"; return "" }

func Styler(attrs ...attribute) func(interface{}) string { _ = "STUB: not implemented"; return nil }
