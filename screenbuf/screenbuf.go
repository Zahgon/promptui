package screenbuf

import (
	"bytes"
	"io"
)

const esc = "\033["

var (
	clearLine = []byte(esc + "2K\r")
	moveUp    = []byte(esc + "1A")
	moveDown  = []byte(esc + "1B")
)

type ScreenBuf struct {
	w      io.Writer
	buf    *bytes.Buffer
	reset  bool
	cursor int
	height int
}

func New(w io.Writer) *ScreenBuf { _ = "STUB: not implemented"; return nil }

func (s *ScreenBuf) Reset() { _ = "STUB: not implemented"; return }

func (s *ScreenBuf) Clear() error { _ = "STUB: not implemented"; return nil }

func (s *ScreenBuf) Write(b []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (s *ScreenBuf) Flush() error { _ = "STUB: not implemented"; return nil }

func (s *ScreenBuf) WriteString(str string) (int, error) { _ = "STUB: not implemented"; return 0, nil }
