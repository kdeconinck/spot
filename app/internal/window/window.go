// =====================================================================================================================
// = LICENSE:       Copyright (c) 2026 Kevin De Coninck
// =
// =                Permission is hereby granted, free of charge, to any person
// =                obtaining a copy of this software and associated documentation
// =                files (the "Software"), to deal in the Software without
// =                restriction, including without limitation the rights to use,
// =                copy, modify, merge, publish, distribute, sublicense, and/or sell
// =                copies of the Software, and to permit persons to whom the
// =                Software is furnished to do so, subject to the following
// =                conditions:
// =
// =                The above copyright notice and this permission notice shall be
// =                included in all copies or substantial portions of the Software.
// =
// =                THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// =                EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
// =                OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// =                NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT
// =                HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY,
// =                WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
// =                FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
// =                OTHER DEALINGS IN THE SOFTWARE.
// =====================================================================================================================

// Package window provides a lightweight cursor for traversing a slice safely.
//
// It is intended for parser-style workflows that need to inspect the current item, look ahead, and advance through a
// stream with bounds checking.
package window

// Window represents a cursor over a slice of items of type T.
type Window[T any] struct {
	// The underlying slice of items being traversed.
	items []T

	// The current absolute cursor position in the slice, starting at zero.
	index int
}

// New returns a new [Window] for items.
func New[T any](items []T) *Window[T] {
	return &Window[T]{
		items: items,
	}
}

// IsAtEnd reports whether the cursor is at or beyond the end of the stream.
func (w *Window[T]) IsAtEnd() bool {
	return w.index >= len(w.items)
}

// Index returns the current cursor position.
func (w *Window[T]) Index() int {
	return w.index
}

// Len returns the number of items in the stream.
func (w *Window[T]) Len() int {
	return len(w.items)
}

// Remaining returns the number of items from the current cursor position to the end of the stream.
func (w *Window[T]) Remaining() int {
	if w.IsAtEnd() {
		return 0
	}

	return len(w.items) - w.index
}

// Current returns the item at the current cursor position.
func (w *Window[T]) Current() (T, bool) {
	if w.IsAtEnd() {
		var zero T

		return zero, false
	}

	return w.items[w.index], true
}

// Peek returns the current item without consuming it.
func (w *Window[T]) Peek() (T, bool) {
	return w.Current()
}

// PeekN returns the item at offset from the current cursor position.
func (w *Window[T]) PeekN(offset int) (T, bool) {
	target := w.index + offset

	if target < 0 || target >= len(w.items) {
		var zero T

		return zero, false
	}

	return w.items[target], true
}

// Previous returns the item immediately before the current cursor position.
func (w *Window[T]) Previous() (T, bool) {
	return w.PeekN(-1)
}

// Advance returns the current item and moves the cursor forward by one position.
func (w *Window[T]) Advance() (T, bool) {
	item, ok := w.Current()

	if !ok {
		var zero T

		return zero, false
	}

	w.index++

	return item, true
}

// AdvanceN moves the cursor forward by n positions.
func (w *Window[T]) AdvanceN(n int) {
	if n <= 0 {
		return
	}

	w.index += n

	if w.index > len(w.items) {
		w.index = len(w.items)
	}
}

// Reset moves the cursor to the start of the stream.
func (w *Window[T]) Reset() {
	w.index = 0
}

// SetIndex moves the cursor to index, clamped to the bounds of the stream.
func (w *Window[T]) SetIndex(index int) {
	switch {
	case index < 0:
		w.index = 0
	case index > len(w.items):
		w.index = len(w.items)
	default:
		w.index = index
	}
}
