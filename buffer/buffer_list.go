package buffer

// ElementMapper provides an identity mapping by default.
//
// This can be replaced to provide a struct that maps elements to linker
// objects, if they are not the same. An ElementMapper is not typically
// required if: Linker is left as is, Element is left as is, or Linker and
// Element are the same type.
type bufferElementMapper struct{}

// linkerFor maps an Element to a Linker.
//
// This default implementation should be inlined.
//
//go:nosplit
func (bufferElementMapper) linkerFor(elem *Buffer) *Buffer { return elem }

// List is an intrusive list. Entries can be added to or removed from the list
// in O(1) time and with no additional memory allocations.
//
// The zero value for List is an empty list ready to use.
//
// To iterate over a list (where l is a List):
//      for e := l.Front(); e != nil; e = e.Next() {
// 		// do something with e.
//      }
//
// +stateify savable
type bufferList struct {
	head *Buffer
	tail *Buffer
}

// Reset resets list l to the empty state.
func (l *bufferList) Reset() {
	l.head = nil
	l.tail = nil
}

// Empty returns true iff the list is empty.
func (l *bufferList) Empty() bool {
	return l.head == nil
}

// Front returns the first element of list l or nil.
func (l *bufferList) Front() *Buffer {
	return l.head
}

// Back returns the last element of list l or nil.
func (l *bufferList) Back() *Buffer {
	return l.tail
}

// PushFront inserts the element e at the front of list l.
func (l *bufferList) PushFront(e *Buffer) {
	linker := bufferElementMapper{}.linkerFor(e)
	linker.SetNext(l.head)
	linker.SetPrev(nil)

	if l.head != nil {
		bufferElementMapper{}.linkerFor(l.head).SetPrev(e)
	} else {
		l.tail = e
	}

	l.head = e
}

// PushBack inserts the element e at the back of list l.
func (l *bufferList) PushBack(e *Buffer) {
	linker := bufferElementMapper{}.linkerFor(e)
	linker.SetNext(nil)
	linker.SetPrev(l.tail)

	if l.tail != nil {
		bufferElementMapper{}.linkerFor(l.tail).SetNext(e)
	} else {
		l.head = e
	}

	l.tail = e
}

// PushBackList inserts list m at the end of list l, emptying m.
func (l *bufferList) PushBackList(m *bufferList) {
	if l.head == nil {
		l.head = m.head
		l.tail = m.tail
	} else if m.head != nil {
		bufferElementMapper{}.linkerFor(l.tail).SetNext(m.head)
		bufferElementMapper{}.linkerFor(m.head).SetPrev(l.tail)

		l.tail = m.tail
	}

	m.head = nil
	m.tail = nil
}

// InsertAfter inserts e after b.
func (l *bufferList) InsertAfter(b, e *Buffer) {
	bLinker := bufferElementMapper{}.linkerFor(b)
	eLinker := bufferElementMapper{}.linkerFor(e)

	a := bLinker.Next()

	eLinker.SetNext(a)
	eLinker.SetPrev(b)
	bLinker.SetNext(e)

	if a != nil {
		bufferElementMapper{}.linkerFor(a).SetPrev(e)
	} else {
		l.tail = e
	}
}

// InsertBefore inserts e before a.
func (l *bufferList) InsertBefore(a, e *Buffer) {
	aLinker := bufferElementMapper{}.linkerFor(a)
	eLinker := bufferElementMapper{}.linkerFor(e)

	b := aLinker.Prev()
	eLinker.SetNext(a)
	eLinker.SetPrev(b)
	aLinker.SetPrev(e)

	if b != nil {
		bufferElementMapper{}.linkerFor(b).SetNext(e)
	} else {
		l.head = e
	}
}

// Remove removes e from l.
func (l *bufferList) Remove(e *Buffer) {
	prev := bufferElementMapper{}.linkerFor(e).Prev()
	next := bufferElementMapper{}.linkerFor(e).Next()

	if prev != nil {
		bufferElementMapper{}.linkerFor(prev).SetNext(next)
	} else {
		l.head = next
	}

	if next != nil {
		bufferElementMapper{}.linkerFor(next).SetPrev(prev)
	} else {
		l.tail = prev
	}
}

// Entry is a default implementation of Linker. Users can add anonymous fields
// of this type to their structs to make them automatically implement the
// methods needed by List.
//
// +stateify savable
type bufferEntry struct {
	next *Buffer
	prev *Buffer
}

// Next returns the entry that follows e in the list.
func (e *bufferEntry) Next() *Buffer {
	return e.next
}

// Prev returns the entry that precedes e in the list.
func (e *bufferEntry) Prev() *Buffer {
	return e.prev
}

// SetNext assigns 'entry' as the entry that follows e in the list.
func (e *bufferEntry) SetNext(elem *Buffer) {
	e.next = elem
}

// SetPrev assigns 'entry' as the entry that precedes e in the list.
func (e *bufferEntry) SetPrev(elem *Buffer) {
	e.prev = elem
}
