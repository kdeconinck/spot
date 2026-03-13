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

// Verifies that [loc.Position.IsValid] reports whether a [loc.Position] contains a valid human-readable source
// location.
func Test_Position_IsValid(t *testing.T) {
	t.Parallel()

	for tcName, tc := range map[string]struct {
		input     loc.Position
		wantValid bool
	}{
		"When both the line and column are greater than zero, the `Position` is valid.": {
			input: loc.Position{
				Line:   1,
				Column: 1,
			},
			wantValid: true,
		},
		"When the line is zero, the `Position` is invalid.": {
			input: loc.Position{
				Line:   0,
				Column: 1,
			},
			wantValid: false,
		},
		"When the column is zero, the `Position` is invalid.": {
			input: loc.Position{
				Line:   1,
				Column: 0,
			},
			wantValid: false,
		},
		"When the line is negative, the `Position` is invalid.": {
			input: loc.Position{
				Line:   -1,
				Column: 1,
			},
			wantValid: false,
		},
		"When the column is negative, the `Position` is invalid.": {
			input: loc.Position{
				Line:   1,
				Column: -1,
			},
			wantValid: false,
		},
		"When both the line and column are negative, the `Position` is invalid.": {
			input: loc.Position{
				Line:   -1,
				Column: -1,
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
