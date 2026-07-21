package list

type Searcher func(input string, index int) bool

const NotFound = -1

type List struct {
	items    []*interface{}
	scope    []*interface{}
	cursor   int
	size     int
	start    int
	Searcher Searcher
}

func New(items interface{}, size int) (*List, error) { _ = "STUB: not implemented"; return nil, nil }

func (l *List) Prev() { _ = "STUB: not implemented"; return }

func (l *List) Search(term string) { _ = "STUB: not implemented"; return }

func (l *List) CancelSearch() { _ = "STUB: not implemented"; return }

func (l *List) search(term string) { _ = "STUB: not implemented"; return }

func (l *List) Start() int { _ = "STUB: not implemented"; return 0 }

func (l *List) SetStart(i int) { _ = "STUB: not implemented"; return }

func (l *List) SetCursor(i int) { _ = "STUB: not implemented"; return }

func (l *List) Next() { _ = "STUB: not implemented"; return }

func (l *List) PageUp() { _ = "STUB: not implemented"; return }

func (l *List) PageDown() { _ = "STUB: not implemented"; return }

func (l *List) CanPageDown() bool { _ = "STUB: not implemented"; return false }

func (l *List) CanPageUp() bool { _ = "STUB: not implemented"; return false }

func (l *List) Index() int { _ = "STUB: not implemented"; return 0 }

func (l *List) Items() ([]interface{}, int) { _ = "STUB: not implemented"; return nil, 0 }
