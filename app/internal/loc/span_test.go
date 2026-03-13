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

// Verify the public API of the loc package.
//
// Tests in this package are written against the exported API only.
// This ensures that position and span validation behavior is tested through the same surface that external consumers
// would use.
package loc_test

import (
	"testing"

	"github.com/kdeconinck/spot/internal/assert"
	"github.com/kdeconinck/spot/internal/loc"
)

// Verifies that [loc.Span.IsValid] reports whether a [loc.Span] contains valid start and end positions in a
// non-decreasing order.
func Test_Span_IsValid(t *testing.T) {
	t.Parallel()

	for tcName, tc := range map[string]struct {
		input     loc.Span
		wantValid bool
	}{
		"When both positions are valid and the end is after the start on the same line, the `Span` is valid.": {
			input: loc.Span{
				Start: loc.Position{
					Line:   3,
					Column: 2,
				},
				End: loc.Position{
					Line:   3,
					Column: 8,
				},
			},
			wantValid: true,
		},
		"When both positions are valid and the end is equal to the start, the `Span` is valid.": {
			input: loc.Span{
				Start: loc.Position{
					Line:   3,
					Column: 2,
				},
				End: loc.Position{
					Line:   3,
					Column: 2,
				},
			},
			wantValid: true,
		},
		"When both positions are valid and the end is on a later line, the `Span` is valid.": {
			input: loc.Span{
				Start: loc.Position{
					Line:   3,
					Column: 9,
				},
				End: loc.Position{
					Line:   4,
					Column: 1,
				},
			},
			wantValid: true,
		},
		"When the start position is invalid, the `Span` is invalid.": {
			input: loc.Span{
				Start: loc.Position{
					Line:   0,
					Column: 1,
				},
				End: loc.Position{
					Line:   1,
					Column: 1,
				},
			},
			wantValid: false,
		},
		"When the end position is invalid, the `Span` is invalid.": {
			input: loc.Span{
				Start: loc.Position{
					Line:   1,
					Column: 1,
				},
				End: loc.Position{
					Line:   0,
					Column: 1,
				},
			},
			wantValid: false,
		},
		"When both positions are invalid, the `Span` is invalid.": {
			input: loc.Span{
				Start: loc.Position{
					Line:   0,
					Column: 0,
				},
				End: loc.Position{
					Line:   0,
					Column: 0,
				},
			},
			wantValid: false,
		},
		"When the end line is before the start line, the `Span` is invalid.": {
			input: loc.Span{
				Start: loc.Position{
					Line:   4,
					Column: 1,
				},
				End: loc.Position{
					Line:   3,
					Column: 99,
				},
			},
			wantValid: false,
		},
		"When the end column is before the start column on the same line, the `Span` is invalid.": {
			input: loc.Span{
				Start: loc.Position{
					Line:   4,
					Column: 10,
				},
				End: loc.Position{
					Line:   4,
					Column: 9,
				},
			},
			wantValid: false,
		},
	} {
		t.Run(tcName, func(t *testing.T) {
			t.Parallel()

			// Act.
			gotValid := tc.input.IsValid()

			// Assert.
			assert.Equal(t, tcName, gotValid, tc.wantValid, "Is valid?")
		})
	}
}
