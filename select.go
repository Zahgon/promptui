package promptui

import (
	"io"
	"text/template"

	"github.com/manifoldco/promptui/list"
	"github.com/manifoldco/promptui/screenbuf"
)

const SelectedAdd = -1

type Select struct {
	Label interface{}

	Items interface{}

	Size int

	CursorPos int

	IsVimMode bool

	HideHelp bool

	HideSelected bool

	Templates *SelectTemplates

	Keys *SelectKeys

	Searcher list.Searcher

	StartInSearchMode bool

	list *list.List

	Pointer Pointer

	Stdin  io.ReadCloser
	Stdout io.WriteCloser
}

type SelectKeys struct {
	Next Key

	Prev Key

	PageUp Key

	PageDown Key

	Search Key
}

type Key struct {
	Code rune

	Display string
}

type SelectTemplates struct {
	Label string

	Active string

	Inactive string

	Selected string

	Details string

	Help string

	FuncMap template.FuncMap

	label    *template.Template
	active   *template.Template
	inactive *template.Template
	selected *template.Template
	details  *template.Template
	help     *template.Template
}

var SearchPrompt = "Search: "

func (s *Select) Run() (int, string, error) { _ = "STUB: not implemented"; return 0, "", nil }

func (s *Select) RunCursorAt(cursorPos, scroll int) (int, string, error) {
	_ = "STUB: not implemented"
	return 0, "", nil
}

func (s *Select) innerRun(cursorPos, scroll int, top rune) (int, string, error) {
	_ = "STUB: not implemented"
	return 0, "", nil
}

func (s *Select) ScrollPosition() int { _ = "STUB: not implemented"; return 0 }

func (s *Select) prepareTemplates() error { _ = "STUB: not implemented"; return nil }

type SelectWithAdd struct {
	Label string

	Items []string

	AddLabel string

	Validate ValidateFunc

	IsVimMode bool

	Pointer Pointer

	HideHelp bool
}

func (sa *SelectWithAdd) Run() (int, string, error) { _ = "STUB: not implemented"; return 0, "", nil }

func (s *Select) setKeys() { _ = "STUB: not implemented"; return }

func (s *Select) renderDetails(item interface{}) [][]byte { _ = "STUB: not implemented"; return nil }

func (s *Select) renderHelp(b bool) []byte { _ = "STUB: not implemented"; return nil }

func render(tpl *template.Template, data interface{}) []byte { _ = "STUB: not implemented"; return nil }

func clearScreen(sb *screenbuf.ScreenBuf) { _ = "STUB: not implemented"; return }
