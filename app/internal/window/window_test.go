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

// Verify the public API of the window package.
//
// Tests in this package are written against the exported API only.
// This ensures that traversal behavior is tested through the same surface that external consumers would use.
package window_test

import (
	"testing"

	"github.com/kdeconinck/spot/internal/assert"
	"github.com/kdeconinck/spot/internal/window"
)

// Verifies that [window.New] returns a usable [window.Window] at the start of the input.
func Test_New(t *testing.T) {
	t.Parallel()

	// Arrange.
	items := []int{10, 20, 30}

	// Act.
	gotWindow := window.New(items)
	gotItem, gotOK := gotWindow.Current()

	// Assert.
	assert.NotNil(t, "When a `Window` is created, it is usable.", gotWindow, "Window")
	assert.Equal(t, "When a `Window` is created, it starts at index 0.", gotWindow.Index(), 0, "Index")
	assert.Equal(t, "When a `Window` is created, it reports the input length.", gotWindow.Len(), 3, "Length")
	assert.Equal(t, "When a `Window` is created, all items remain.", gotWindow.Remaining(), 3, "Remaining")
	assert.Equal(t, "When a `Window` is created, it is not at the end.", gotWindow.IsAtEnd(), false, "Is at end?")
	assert.Equal(t, "When a `Window` is created, `Current` returns the first item.", gotItem, 10, "Current item")
	assert.Equal(t, "When a `Window` is created, `Current` reports success.", gotOK, true, "Current OK?")
}

// Verifies that [window.Window.Index], [window.Window.Len], [window.Window.Remaining], and [window.Window.IsAtEnd]
// reflects the cursor state.
func Test_Window_PositionState(t *testing.T) {
	t.Parallel()

	for tcName, tc := range map[string]struct {
		startIndex    int
		wantIndex     int
		wantLen       int
		wantRemaining int
		wantIsAtEnd   bool
	}{
		"When the cursor is at the start, the `Window` is not at the end and all items remain.": {
			startIndex:    0,
			wantIndex:     0,
			wantLen:       3,
			wantRemaining: 3,
			wantIsAtEnd:   false,
		},
		"When the cursor is in the middle, the `Window` is not at the end and only trailing items remain.": {
			startIndex:    1,
			wantIndex:     1,
			wantLen:       3,
			wantRemaining: 2,
			wantIsAtEnd:   false,
		},
		"When the cursor is at len(items), the `Window` is at the end and nothing remains.": {
			startIndex:    3,
			wantIndex:     3,
			wantLen:       3,
			wantRemaining: 0,
			wantIsAtEnd:   true,
		},
		"When the cursor is clamped beyond `len(items)`, the `Window` remains at the end and nothing remains.": {
			startIndex:    99,
			wantIndex:     3,
			wantLen:       3,
			wantRemaining: 0,
			wantIsAtEnd:   true,
		},
	} {
		t.Run(tcName, func(t *testing.T) {
			t.Parallel()

			// Arrange.
			items := []int{10, 20, 30}
			window := window.New(items)
			window.SetIndex(tc.startIndex)

			// Act.
			gotIndex := window.Index()
			gotLen := window.Len()
			gotRemaining := window.Remaining()
			gotIsAtEnd := window.IsAtEnd()

			// Assert.
			assert.Equal(t, tcName, gotIndex, tc.wantIndex, "Index")
			assert.Equal(t, tcName, gotLen, tc.wantLen, "Length")
			assert.Equal(t, tcName, gotRemaining, tc.wantRemaining, "Remaining elements #")
			assert.Equal(t, tcName, gotIsAtEnd, tc.wantIsAtEnd, "Is at end?")
		})
	}
}

// Verifies that [window.Window.Current] and [window.Window.Peek] report the current item when available.
func Test_Window_CurrentAndPeek(t *testing.T) {
	t.Parallel()

	for tcName, tc := range map[string]struct {
		startIndex int
		wantItem   string
		wantOK     bool
	}{
		"When the cursor is in bounds, `Current` and `Peek` return the current item.": {
			startIndex: 1,
			wantItem:   "b",
			wantOK:     true,
		},
		"When the cursor is at the end, `Current` and `Peek` report no item.": {
			startIndex: 3,
			wantItem:   "",
			wantOK:     false,
		},
	} {
		t.Run(tcName, func(t *testing.T) {
			t.Parallel()

			// Arrange.
			items := []string{"a", "b", "c"}
			window := window.New(items)
			window.SetIndex(tc.startIndex)

			// Act.
			gotCurrentItem, gotCurrentOK := window.Current()
			gotPeekItem, gotPeekOK := window.Peek()

			// Assert.
			assert.Equal(t, tcName, gotCurrentItem, tc.wantItem, "Current item")
			assert.Equal(t, tcName, gotCurrentOK, tc.wantOK, "Current OK?")
			assert.Equal(t, tcName, gotPeekItem, tc.wantItem, "Peek item")
			assert.Equal(t, tcName, gotPeekOK, tc.wantOK, "Peek OK?")
		})
	}
}

// Verifies that [window.Window.PeekN] supports bounded look-ahead and look-behind.
func Test_Window_PeekN(t *testing.T) {
	t.Parallel()

	for tcName, tc := range map[string]struct {
		startIndex int
		offset     int
		wantItem   int
		wantOK     bool
	}{
		"When the offset is zero, `PeekN` returns the current item.": {
			startIndex: 1,
			offset:     0,
			wantItem:   20,
			wantOK:     true,
		},
		"When the offset is positive and in bounds, `PeekN` returns a future item.": {
			startIndex: 1,
			offset:     1,
			wantItem:   30,
			wantOK:     true,
		},
		"When the offset is negative and in bounds, `PeekN` returns a previous item.": {
			startIndex: 1,
			offset:     -1,
			wantItem:   10,
			wantOK:     true,
		},
		"When the target position is before the start, `PeekN` reports no item.": {
			startIndex: 0,
			offset:     -1,
			wantItem:   0,
			wantOK:     false,
		},
		"When the target position is after the end, `PeekN` reports no item.": {
			startIndex: 2,
			offset:     1,
			wantItem:   0,
			wantOK:     false,
		},
		"When the cursor is already at the end, `PeekN` still supports valid look-behind access.": {
			startIndex: 3,
			offset:     -1,
			wantItem:   30,
			wantOK:     true,
		},
	} {
		t.Run(tcName, func(t *testing.T) {
			t.Parallel()

			// Arrange.
			items := []int{10, 20, 30}
			gotWindow := window.New(items)
			gotWindow.SetIndex(tc.startIndex)

			// Act.
			gotItem, gotOK := gotWindow.PeekN(tc.offset)

			// Assert.
			assert.Equal(t, tcName, gotItem, tc.wantItem, "Item")
			assert.Equal(t, tcName, gotOK, tc.wantOK, "OK?")
		})
	}
}

// Verifies that [window.Window.Previous] returns the previous item when available.
func Test_Window_Previous(t *testing.T) {
	t.Parallel()

	for tcName, tc := range map[string]struct {
		startIndex int
		wantItem   string
		wantOK     bool
	}{
		"When the cursor is at the start, `Previous` reports no item.": {
			startIndex: 0,
			wantItem:   "",
			wantOK:     false,
		},
		"When the cursor is in the middle, `Previous` returns the preceding item.": {
			startIndex: 2,
			wantItem:   "b",
			wantOK:     true,
		},
		"When the cursor is at the end, `Previous` returns the last item.": {
			startIndex: 3,
			wantItem:   "c",
			wantOK:     true,
		},
	} {
		t.Run(tcName, func(t *testing.T) {
			t.Parallel()

			// Arrange.
			items := []string{"a", "b", "c"}
			gotWindow := window.New(items)
			gotWindow.SetIndex(tc.startIndex)

			// Act.
			gotItem, gotOK := gotWindow.Previous()

			// Assert.
			assert.Equal(t, tcName, gotItem, tc.wantItem, "Item")
			assert.Equal(t, tcName, gotOK, tc.wantOK, "OK?")
		})
	}
}

// Verifies that [window.Window.Advance] returns the current item and advances the cursor.
func Test_Window_Advance(t *testing.T) {
	t.Parallel()

	for tcName, tc := range map[string]struct {
		startIndex  int
		wantItem    int
		wantOK      bool
		wantIndex   int
		wantIsAtEnd bool
	}{
		"When the cursor is in bounds, `Advance` returns the current item and moves forward by one.": {
			startIndex:  1,
			wantItem:    20,
			wantOK:      true,
			wantIndex:   2,
			wantIsAtEnd: false,
		},
		"When the cursor is at the last item, `Advance` moves the cursor to the end.": {
			startIndex:  2,
			wantItem:    30,
			wantOK:      true,
			wantIndex:   3,
			wantIsAtEnd: true,
		},
		"When the cursor is already at the end, `Advance` reports no item and does not move.": {
			startIndex:  3,
			wantItem:    0,
			wantOK:      false,
			wantIndex:   3,
			wantIsAtEnd: true,
		},
	} {
		t.Run(tcName, func(t *testing.T) {
			t.Parallel()

			// Arrange.
			items := []int{10, 20, 30}
			gotWindow := window.New(items)
			gotWindow.SetIndex(tc.startIndex)

			// Act.
			gotItem, gotOK := gotWindow.Advance()

			// Assert.
			assert.Equal(t, tcName, gotItem, tc.wantItem, "Item")
			assert.Equal(t, tcName, gotOK, tc.wantOK, "OK?")
			assert.Equal(t, tcName, gotWindow.Index(), tc.wantIndex, "Index after")
			assert.Equal(t, tcName, gotWindow.IsAtEnd(), tc.wantIsAtEnd, "Is at end after?")
		})
	}
}

// Verifies that [window.Window.AdvanceN] ignores non-positive values and clamps at the end.
func Test_Window_AdvanceN(t *testing.T) {
	t.Parallel()

	for tcName, tc := range map[string]struct {
		startIndex       int
		advanceBy        int
		wantIndexAfter   int
		wantIsAtEndAfter bool
	}{
		"When n is negative, `AdvanceN` leaves the cursor unchanged.": {
			startIndex:       1,
			advanceBy:        -1,
			wantIndexAfter:   1,
			wantIsAtEndAfter: false,
		},
		"When n is zero, `AdvanceN` leaves the cursor unchanged.": {
			startIndex:       1,
			advanceBy:        0,
			wantIndexAfter:   1,
			wantIsAtEndAfter: false,
		},
		"When n is positive and stays in bounds, `AdvanceN` moves the cursor forward by n.": {
			startIndex:       1,
			advanceBy:        1,
			wantIndexAfter:   2,
			wantIsAtEndAfter: false,
		},
		"When n would move beyond the end, `AdvanceN` clamps the cursor to `len(items)`.": {
			startIndex:       1,
			advanceBy:        10,
			wantIndexAfter:   3,
			wantIsAtEndAfter: true,
		},
	} {
		t.Run(tcName, func(t *testing.T) {
			t.Parallel()

			// Arrange.
			items := []int{10, 20, 30}
			gotWindow := window.New(items)
			gotWindow.SetIndex(tc.startIndex)

			// Act.
			gotWindow.AdvanceN(tc.advanceBy)

			// Assert.
			assert.Equal(t, tcName, gotWindow.Index(), tc.wantIndexAfter, "Index after")
			assert.Equal(t, tcName, gotWindow.IsAtEnd(), tc.wantIsAtEndAfter, "Is at end after?")
		})
	}
}

// Verifies that [window.Window.Reset] moves the cursor back to the start of the stream.
func Test_Window_Reset(t *testing.T) {
	t.Parallel()

	// Arrange.
	items := []int{10, 20, 30}
	gotWindow := window.New(items)
	gotWindow.SetIndex(2)

	// Act.
	gotWindow.Reset()
	gotItem, gotOK := gotWindow.Current()

	// Assert.
	assert.Equal(t, "When the `Window` is reset, its index becomes 0.", gotWindow.Index(), 0, "Index")
	assert.Equal(t, "When the `Window` is reset, all items remain.", gotWindow.Remaining(), 3, "Remaining #")
	assert.Equal(t, "When the `Window` is reset, it is not at the end.", gotWindow.IsAtEnd(), false, "Is at end?")
	assert.Equal(t, "When the `Window` is reset, `Current` returns the first item.", gotItem, 10, "Current item")
	assert.Equal(t, "When the `Window` is reset, `Current` reports success.", gotOK, true, "Current OK?")
}

// Verifies that [window.Window.SetIndex] clamps indices to the bounds of the stream.
func Test_Window_SetIndex(t *testing.T) {
	t.Parallel()

	for tcName, tc := range map[string]struct {
		inputIndex    int
		wantIndex     int
		wantIsAtEnd   bool
		wantRemaining int
	}{
		"When the index is negative, `SetIndex` clamps it to 0.": {
			inputIndex:    -5,
			wantIndex:     0,
			wantIsAtEnd:   false,
			wantRemaining: 3,
		},
		"When the index is within bounds, `SetIndex` moves the cursor to that position.": {
			inputIndex:    2,
			wantIndex:     2,
			wantIsAtEnd:   false,
			wantRemaining: 1,
		},
		"When the index is greater than `len(items)`, `SetIndex` clamps it to the end.": {
			inputIndex:    42,
			wantIndex:     3,
			wantIsAtEnd:   true,
			wantRemaining: 0,
		},
	} {
		t.Run(tcName, func(t *testing.T) {
			t.Parallel()

			// Arrange.
			items := []int{10, 20, 30}
			gotWindow := window.New(items)

			// Act.
			gotWindow.SetIndex(tc.inputIndex)

			// Assert.
			assert.Equal(t, tcName, gotWindow.Index(), tc.wantIndex, "Index")
			assert.Equal(t, tcName, gotWindow.IsAtEnd(), tc.wantIsAtEnd, "Is at end?")
			assert.Equal(t, tcName, gotWindow.Remaining(), tc.wantRemaining, "Remaining #")
		})
	}
}

// Verifies that [window.New] preserves `<nil>` input and still returns a usable window.
func Test_New_NilSlice(t *testing.T) {
	t.Parallel()

	// Arrange.
	var items []int

	gotWindow := window.New(items)

	// Act.
	gotItem, gotOK := gotWindow.Current()

	// Assert.
	assert.NotNil(t, "When a `Window` is created from a `<nil>` slice, it is still usable.", gotWindow, "Window")
	assert.Equal(t, "When a `Window` is created from a `<nil>` slice, the input slice remains `<nil>`.", items == nil, true, "Input slice is nil")
	assert.Equal(t, "When a `Window` is created from a `<nil>` slice, its length is 0.", gotWindow.Len(), 0, "Length")
	assert.Equal(t, "When a `Window` is created from a `<nil>` slice, its index is 0.", gotWindow.Index(), 0, "Index")
	assert.Equal(t, "When a `Window` is created from a `<nil>` slice, no items remain.", gotWindow.Remaining(), 0, "Remaining #")
	assert.Equal(t, "When a `Window` is created from a `<nil>` slice, it is at the end.", gotWindow.IsAtEnd(), true, "Is at end?")
	assert.Equal(t, "When a `Window` is created from a `<nil>` slice, `Current` returns the zero value.", gotItem, 0, "Current item")
	assert.Equal(t, "When a `Window` is created from a `<nil>` slice, `Current` reports failure.", gotOK, false, "Current OK?")
}
