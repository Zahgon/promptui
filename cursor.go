package promptui

type Pointer func(to []rune) []rune

func defaultCursor(ignored []rune) []rune { _ = "STUB: not implemented"; return nil }

func blockCursor(input []rune) []rune { _ = "STUB: not implemented"; return nil }

func pipeCursor(input []rune) []rune { _ = "STUB: not implemented"; return nil }

var (
	DefaultCursor Pointer = defaultCursor

	BlockCursor Pointer = blockCursor

	PipeCursor Pointer = pipeCursor
)

type Cursor struct {
	Cursor Pointer

	input []rune

	Position int
	erase    bool
}

func NewCursor(startinginput string, pointer Pointer, eraseDefault bool) Cursor {
	_ = "STUB: not implemented"
	return *new(Cursor)
}

func (c *Cursor) String() string { _ = "STUB: not implemented"; return "" }

func (c *Cursor) End() { _ = "STUB: not implemented"; return }

func (c *Cursor) Start() { _ = "STUB: not implemented"; return }

func (c *Cursor) correctPosition() { _ = "STUB: not implemented"; return }

func format(a []rune, c *Cursor) string { _ = "STUB: not implemented"; return "" }

func (c *Cursor) Format() string { _ = "STUB: not implemented"; return "" }

func (c *Cursor) FormatMask(mask rune) string { _ = "STUB: not implemented"; return "" }

func (c *Cursor) Update(newinput string) { _ = "STUB: not implemented"; return }

func (c *Cursor) Get() string { _ = "STUB: not implemented"; return "" }

func (c *Cursor) GetMask(mask rune) string { _ = "STUB: not implemented"; return "" }

func (c *Cursor) Replace(input string) { _ = "STUB: not implemented"; return }

func (c *Cursor) Place(position int) { _ = "STUB: not implemented"; return }

func (c *Cursor) Move(shift int) { _ = "STUB: not implemented"; return }

func (c *Cursor) Backspace() { _ = "STUB: not implemented"; return }

func (c *Cursor) Listen(line []rune, pos int, key rune) ([]rune, int, bool) {
	_ = "STUB: not implemented"
	return nil, 0, false
}
