package promptui

import (
	"io"
	"text/template"
)

type Prompt struct {
	Label interface{}

	Default string

	AllowEdit bool

	Validate ValidateFunc

	Mask rune

	HideEntered bool

	Templates *PromptTemplates

	IsConfirm bool

	IsVimMode bool

	Pointer Pointer

	Stdin  io.ReadCloser
	Stdout io.WriteCloser
}

type PromptTemplates struct {
	Prompt string

	Confirm string

	Valid string

	Invalid string

	Success string

	ValidationError string

	FuncMap template.FuncMap

	prompt     *template.Template
	valid      *template.Template
	invalid    *template.Template
	validation *template.Template
	success    *template.Template
}

func (p *Prompt) Run() (string, error) { _ = "STUB: not implemented"; return "", nil }

func (p *Prompt) prepareTemplates() error { _ = "STUB: not implemented"; return nil }
